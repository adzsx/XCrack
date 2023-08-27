// Cracking a hash with brute force

package crack

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adzsx/xcrack/pkg/utils"
)

// setting up brute force mode
func BruteSetup(input utils.Input) (string, time.Duration) {
	now := time.Now()
	var status int

	if input.Password == "" {
		fmt.Println("Please specify the password")
		os.Exit(0)
	}

	// Jobs (cores) for each length on cpu
	jobs := make(chan int, input.Max-input.Min)
	result := make(chan string)

	// Cracking
	for i := 0; i < (input.Max - input.Min + 1); i++ {
		go brute(input.Password, input.Hash, input.Chars, jobs, result, &status)
	}

	// Gettings results
	for i := input.Min; i <= input.Max; i++ {
		jobs <- i
	}

	close(jobs)

	for {
		if status == 1 {
			return <-result, time.Since(now)
		}
	}
}

// Brute forcer
func brute(password string, htype string, chars []string, jobs <-chan int, result chan<- string, status *int) {
	for currentLength := range jobs {
		counter := make([]int, currentLength)
		curPass := make([]string, currentLength)
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
				curPass[index] = chars[value]
			}
			pw := strings.Join(curPass[:], "")
			pwh := Hash(pw, htype)
			if pwh == password {
				*status = 1
				result <- pw
				return
			}

		}

	}
	*status = 2
}

// hashing function, (Here for faster results)
func Hash(text string, htype string) string {
	switch htype {
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
	case "sha512":
		h := sha512.New()
		h.Write([]byte(text))
		hash := h.Sum(nil)
		return fmt.Sprintf("%x", hash)
	}
	return "Hash type not found"
}
