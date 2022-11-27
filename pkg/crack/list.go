// Cracking a hash with a wordlist

package crack

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func WlistSet(password string, htype string, paths []string) {
	if Hash("checking...", htype) == "Hash type not found" {
		fmt.Println("The hash type was not found")
		os.Exit(1)
	}

	jobs := make(chan string, len(paths))
	result := make(chan bool)

	for i := 0; i < len(paths); i++ {
		go wordlist(password, htype, jobs)
	}

	for i, path := range paths {
		jobs <- path
	}
}

func wordlist(password string, htype string, jobs chan string) {
	now := time.Now()

	fmt.Println("Starting wordlist mode")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Path \"%v\" found. Plase enter a valid path!\n", path)
		os.Exit(1)
	} else {
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
		fmt.Println("Password not found")
	}
	file.Close()
}
