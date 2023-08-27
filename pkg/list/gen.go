// Wordlist generation

package list

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"github.com/adzsx/xcrack/pkg/utils"
)

var (
	sizes   = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB", "RB", "QB"}
	numbers = []string{"", "Thousand", "Million", "Billion", "Trillion", "Quadrillion", "Quintillion", "Sextillion", "Septillion", "Octillion", "Nonillion", "Decillion"}

	pws    float64
	tchars float64

	sizeUnit string
	numUnit  string
)

func WgenSetup(input utils.Input, showSize bool) (bool, time.Duration) {

	if showSize {
		var inp string
		stSize, pwsize := size(input)
		fmt.Printf("The wordlist will be %v big and contain %v Passwords.\nDo you want to continue? [y/n] ", stSize, pwsize)
		fmt.Scanln(&inp)
		if inp != "y" {
			os.Exit(0)
		}
	}

	now := time.Now()

	if len(input.Chars) == 0 {
		fmt.Println("Specify the characters used to generate the wordlist")
		os.Exit(0)
	}

	// some variables for generating the wordlist
	file, _ := os.Create(input.Output)

	// Create list with characters included in the password

	// length of passwords to be generated
	jobs := make(chan int, input.Max-input.Min)
	response := make(chan bool, input.Max-input.Min)

	for i := 0; i < input.Max-input.Min+1; i++ {
		go gen(input.Chars, jobs, response, file)
	}

	for i := input.Min; i <= input.Max; i++ {
		jobs <- i
	}

	close(jobs)

	var finished []bool
	for i := range response {
		finished = append(finished, i)
		if len(finished) > input.Max-input.Min {

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
		for utils.Sum(counter) < total {

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

func size(input utils.Input) (string, string) {
	// return: size, in unit, number of passwords

	chars := len(input.Chars)
	var lengths []int

	for i := input.Min; i <= input.Max; i++ {
		lengths = append(lengths, i)
	}

	for _, length := range lengths {
		pws += math.Pow(float64(chars), float64(length))
		tchars += float64(math.Pow(float64(chars), float64(length)) * float64(length))
	}

	tchars += pws

	for _, element := range sizes {
		sizeUnit = element
		if tchars < 1024 {
			break
		}
		tchars = tchars / 1024
	}

	for _, element := range numbers {
		numUnit = element
		if pws < 1000 {
			break
		}
		pws = pws / 1000
	}

	return fmt.Sprintf("%.2f %v", tchars, sizeUnit), fmt.Sprintf("%.2f %v", pws, numUnit)
}
