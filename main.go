package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

var (
	type_    string
	password string
	islist   bool
	path     string
)

func main() {
	args := os.Args
	help := "\nBrute forcer cracks hashed passwords:\n------------------------------------------------------------------------------------------\n-h: 			Shows this message (ignores other arguments)\n-p HASH:		(required) Sets the HASH\n-t HASH-TYPE:		(required) specify the HASH-TYPE: (md5, sha1, sha256, sha512)\n-n: 			numbers\n-l: 			lowercase letters\n-L: 			uppercase letters\n-s: 			special Characters\n-m LENGTH: 		min LENGTH of password\n-M LENGTH: 		max LENGTH of password\n-w PATH:		uses a wordlist in PATH (ignores other arguments)\n------------------------------------------------------------------------------------------\n"
	args[0] = "Hash-Cracker"
	for n, element := range args {
		if element == "-h" {
			fmt.Println(help)
			break
		} else if element == "-t" {
			type_ = args[n+1]
		} else if element == "-p" {
			password = args[n+1]
		} else if element == "-w" {
			path = args[n+1]
			islist = true
		}
	}
	if islist {
		wordlist(password, type_, path)
	}
}

func wordlist(password string, type_ string, path string) {
	fmt.Println(password, type_, path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Path \"%v\" found. Plase enter a valid path!\n", path)
		os.Exit(1)
	} else {
		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			data := fileScanner.Text()
			if hash(type_, data) == password {
				fmt.Printf("Password: %v \n", data)
				os.Exit(0)
			}
		}
		fmt.Println("Password not found")
	}
	file.Close()
}

func hash(form string, text string) string {
	if form == "md5" {
		hash := md5.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	}
	return ""
}

//530ea1472e71035353d32d341ecf6343
