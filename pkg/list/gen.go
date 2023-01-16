package list

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/adzsx/xcrack/pkg/check"
)

func WgenSetup(chars []string, path string, min int, max int) {
	now := time.Now()
	fmt.Println("Generating wordlist...")

	// some variables for generating the wordlist
	file, _ := os.Create(path)

	// Create list with characters included in the password

	// length of passwords to be generated
	jobs := make(chan int, max-min)
	response := make(chan bool, max-min)

	for i := 0; i < max-min+1; i++ {
		go gen(chars, jobs, response, file)
	}

	for i := min; i <= max; i++ {
		jobs <- i
	}

	close(jobs)

	var finished []bool
	for i := range response {
		finished = append(finished, i)
		if len(finished) > max-min {
			fmt.Println("Done")
			fmt.Printf("\n[%v]\n", time.Since(now))
			os.Exit(0)
		}
	}
}

// Wordlist generation mode
func gen(chars []string, jobs <-chan int, response chan<- bool, file *os.File) {
	// main.Chars = characters for password
	// hashed = hashed password to crack
	// length = length of characters in main.Chars
	// jobs = jobs for lengths for multiple gorutines

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