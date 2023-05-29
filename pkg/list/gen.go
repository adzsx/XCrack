// Wordlist generation

package list

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/adzsx/xcrack/pkg/check"
	"github.com/adzsx/xcrack/pkg/format"
)

func WgenSetup(query format.Query) (bool, time.Duration) {
	now := time.Now()

	if len(query.Chars) == 0 {
		fmt.Println("Specify the characters used to generate the wordlist")
		os.Exit(0)
	}

	// some variables for generating the wordlist
	file, _ := os.Create(query.Output)

	// Create list with characters included in the password

	// length of passwords to be generated
	jobs := make(chan int, query.Max-query.Min)
	response := make(chan bool, query.Max-query.Min)

	for i := 0; i < query.Max-query.Min+1; i++ {
		go gen(query.Chars, jobs, response, file)
	}

	for i := query.Min; i <= query.Max; i++ {
		jobs <- i
	}

	close(jobs)

	var finished []bool
	for i := range response {
		finished = append(finished, i)
		if len(finished) > query.Max-query.Min {

			return true, time.Since(now)
		}
	}
	return false, time.Since(now)
}

// Wordlist generation mode
func gen(chars []string, jobs <-chan int, response chan<- bool, file *os.File) {

	// currentLength: length to generate every password from.
	for currentLength := range jobs {
		counter := make([]int, currentLength)
		password := make([]string, currentLength)
		counter[0] = -1
		total := len(counter) * (len(chars) - 1)
		for check.Sum(counter) < total {

			counter[0] += 1

			for index, value := range counter {
				if value > len(chars)-1 {
					counter[index] = 0

					if len(counter) > index+1 {

						counter[index+1] += 1
						continue

					} else {
						break
					}
				}
			}

			for index, value := range counter {
				password[index] = chars[value]
			}
			pw := strings.Join(password[:], "")
			io.WriteString(file, pw+"\n")
		}

	}
	response <- true
}
