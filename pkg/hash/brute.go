package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adzsx/xcrack/pkg/check"
)

// setting up brute force mode
func bruteSetup(password string, type_ string, chars []string, min int, max int) {
	now := time.Now()
	fmt.Println("Starting brute force mode")

	// chars: all chars used in password
	// password: hashed password
	// type_: type of hash
	// jobs: length to generate password

	jobs := make(chan int, max-min)
	result := make(chan bool)

	for i := 0; i < (max - min + 1); i++ {
		go brute(chars, password, jobs, result)
	}

	for i := min; i <= max; i++ {
		jobs <- i
	}

	close(jobs)

	var finished []bool
	for i := range result {
		finished = append(finished, i)
		if len(finished) >= max-min {
			fmt.Println("Password not found")
			fmt.Printf("\n[%v]\n", time.Since(now))
			os.Exit(0)
		}
	}
}

// Brute forcer
func brute(chars []string, hashed string, jobs <-chan int, response chan<- bool) {
	// chars = characters for password
	// hashed = hashed password to crack
	// length = length of characters in chars
	// jobs = jobs for lengths for multiple gorutines

	for currentLength := range jobs {
		// if len(jobs) == 0 {
		//  fmt.Println("Password not found! Password probably longer than specified length")
		//  fmt.Printf("\n[%v]\n", time.Since(Now))
		//  os.Exit(1)
		// }
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
			pwh := hash(pw, type_)
			if pwh == hashed {
				fmt.Printf("Password: %v\n", pw)
				fmt.Printf("\n[%v]\n", time.Since(now))
				os.Exit(0)
			}

		}

	}
	response <- false
}

// hashing function
func hash(text string, type_ string) string {
	switch type_ {
	case "md5":
		hash := md5.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	case "sha1":
		hash := sha1.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	case "sha256":
		h := sha256.New()
		h.Write([]byte(text))
		hash := h.Sum(nil)
		return fmt.Sprintf("%x", hash)
	}
	return ""
}
