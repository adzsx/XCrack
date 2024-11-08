// Cracking a hash with a wordlist

package crack

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/adzsx/xcrack/internal/utils"
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
		defer file.Close()

		fileInfo, _ := file.Stat()
		fileSize := fileInfo.Size()

		pieceSize := fileSize / int64(runtime.NumCPU())

		var wg sync.WaitGroup
		wg.Add(runtime.NumCPU())

		for i := 0; i < runtime.NumCPU(); i++ {
			// Calculate offset and limit for each piece
			offset := int64(i) * pieceSize
			limit := pieceSize
			if i == runtime.NumCPU()-1 {
				// For the last piece, read until the end of the file
				limit = fileSize - offset
			}

			go func(offset, limit int64) {
				defer wg.Done()

				// Seek to the beginning of the piece
				_, err := file.Seek(offset, 0)
				if err != nil {
					fmt.Println("Error seeking file:", err)
					return
				}

				// Create a buffer to read the piece
				buffer := make([]byte, limit)

				// Read the piece
				_, err = file.Read(buffer)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}

				time.Sleep(time.Second * 5)
				// Process the piece (in this example, just printing it)
				fmt.Printf("Piece %d:\n%s\n", offset/pieceSize, string(buffer))

				fileScanner := bufio.NewScanner(bytes.NewReader(buffer))

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
			}(offset, limit)
		}

		// Wait for all goroutines to finish
		wg.Wait()

		file.Close()
	}
	*status = 2
}
