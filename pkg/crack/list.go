// Cracking a hash with a wordlist

package crack

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Wordlist(password string, type_ string, path string) {
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
			if Hash(data, type_) == password {
				fmt.Printf("Password: %v \n", data)
				fmt.Printf("\n[%v]\n", time.Since(now))
				os.Exit(0)
			}
		}
		fmt.Println("Password not found")
	}
	file.Close()
}
