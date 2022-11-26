// Cracking a hash with a wordlist

package crack

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Wordlist(password string, htype string, path string) {
	if Hash("checking...", htype) == "Hash type not found" {
		fmt.Println("The hash type was not found")
		os.Exit(1)
	}
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
