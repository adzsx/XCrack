// Cracking a hash with a wordlist

package crack

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	now = time.Now()
)

func WlistSet(password string, htype string, paths []string) {
	fmt.Println("Starting wordlist mode")
	if Hash("checking...", htype) == "Hash type not found" {
		fmt.Println("The hash type was not found")
		os.Exit(0)
	}

	jobs := make(chan string, len(paths))
	result := make(chan bool, len(paths))

	for i := 0; i < len(paths); i++ {
		go wordlist(password, htype, jobs, result)
	}

	for _, path := range paths {
		jobs <- path
	}

	var finished []bool
	for i := range result {
		finished = append(finished, i)
		if len(finished) >= len(paths) {
			fmt.Println("Password not found")
			fmt.Printf("\n[%v]\n", time.Since(now))
			os.Exit(0)
		}
	}
}

func wordlist(password string, htype string, jobs <-chan string, response chan<- bool) {
	for path := range jobs {
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Path \"%v\" found. Plase enter a valid path!\n", path)
			os.Exit(0)
		}

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			data := fileScanner.Text()
			if Hash(data, htype) == password {
				fmt.Printf("Password: %v \n", data)
				fmt.Printf("\n[%v]\n", time.Since(now))
				os.Exit(0)
			}
		}
		response <- false
		file.Close()
	}
}
