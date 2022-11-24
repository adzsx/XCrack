package list

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/adzsx/xcrack/pkg/format"
	//"github.com/adzsx/xcrack/pkg/check"
)

var (
	chars []string
)

func WgenSetup(args [6]string, path string) {
	fmt.Println("Starting wordlist generation mode.")

	//some variables for generating the wordlist
	file, _ := os.Create(path)
	min, err := strconv.Atoi(args[4])
	max, err2 := strconv.Atoi(args[5])

	//check errors in string conversion
	check.Err(err)
	check.Err(err2)

	//Create list with characters included in the password

	chars = format.CharList(args)

	//length of passwords to be generated
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
			os.Exit(0)
		}
	}
}

// Wordlist generation mode
func gen(chars []string, jobs <-chan int, response chan<- bool, file *os.File) {
	//main.Chars = characters for password
	//hashed = hashed password to crack
	//length = length of characters in main.Chars
	//jobs = jobs for lengths for multiple gorutines

	for currentLength := range jobs {
		counter := make([]int, currentLength)
		password := make([]string, currentLength)
		counter[0] = -1
		total := len(counter) * (len(chars) - 1)
		for sum(counter) < total {

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

func contains(s [6]string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
