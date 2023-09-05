// Cracking a hash with a wordlist

package crack

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/adzsx/xcrack/pkg/utils"
)

// Cracking with wordlists
func WlistSet(input utils.Input) (string, time.Duration) {
	fmt.Println("Starting wordlist mode")
	now := time.Now()
	var status int

	if Hash("checking...", input.Hash) == "Hash type not found" {
		utils.Err(errors.New("hash type \"" + input.Hash + "\" not found"))
		os.Exit(0)
	}

	// Jobs for different files
	jobs := make(chan string, len(input.Inputs))
	result := make(chan string)

	for i := 0; i < len(input.Inputs); i++ {
		go wordlist(input.Password, input.Hash, jobs, result, &status)
	}

	for _, path := range input.Inputs {
		err, _ := os.Open(path)
		if err != nil {
			jobs <- path
		} else {
			utils.Err(errors.New(path + " not found"))
			os.Exit(0)
		}
	}

	close(jobs)

	for {
		if status == 1 {
			return <-result, time.Since(now)
		}
	}
}

// Open wordlist and try every password in there.
func wordlist(password string, htype string, jobs <-chan string, result chan<- string, status *int) {
	// Iterate over files available
	for path := range jobs {
		file, err := os.Open(path)
		if err != nil {
			utils.Err(errors.New(path + " not found"))
			return
		}

		// File Scanner
		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		// For each line
		for fileScanner.Scan() {

			// Crack
			data := fileScanner.Text()
			if Hash(data, htype) == password {
				*status = 1
				result <- data
				return

			}
		}
		file.Close()
	}
	*status = 2
}
