// Cracking a hash with a wordlist

package crack

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adzsx/xcrack/pkg/format"
)

var (
	now      = time.Now()
	finished bool
)

// Cracking with wordlists
func WlistSet(query format.Query) {

	fmt.Println("Starting wordlist mode")
	if Hash("checking...", query.Hash) == "Hash type not found" {
		fmt.Println("The hash type was not found")
		os.Exit(0)
	}

	// Jobs for different files
	jobs := make(chan string, len(query.Inputs))

	for i := 0; i < len(query.Inputs); i++ {
		go wordlist(query.Password, query.Hash, jobs, &finished)
	}

	for _, path := range query.Inputs {
		err, _ := os.Open(path)
		if err != nil {
			jobs <- path
		}
	}

	for {
		if !finished {
			fmt.Println("Password not found")
			fmt.Printf("\n[%v]\n", time.Since(now))
			return
		} else if finished {
			return
		}
	}
}

// Open wordlist and try every password in there.
func wordlist(password string, htype string, jobs <-chan string, finished *bool) {

	// Iterate over files available
	for path := range jobs {
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Path \"%v\" found. Plase enter a valid path!\n", path)
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
				fmt.Printf("Password: %v \n", data)
				fmt.Printf("\n[%v]\n", time.Since(now))
				*finished = true
				return

			}
		}
		file.Close()
	}
	*finished = false
}
