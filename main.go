package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

var (
	password string
	included [6]string

	type_  string
	islist bool
	path   string

	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers   = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	special   = [39]string{"^", "´", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", "¸", "~", "’", "–", "·"}
)

func main() {
	args := os.Args
	help := "\nBrute forcer cracks hashed passwords:\n------------------------------------------------------------------------------------------\n-h: 			Shows this message (ignores other arguments)\n-p HASH:		(required) Sets the HASH\n-t HASH-TYPE:		(required) specify the HASH-TYPE: (md5, sha1)\n-n: 			numbers\n-l: 			lowercase letters\n-L: 			uppercase letters\n-s: 			special Characters\n-m LENGTH: 		min LENGTH of password\n-M LENGTH: 		max LENGTH of password\n-w PATH:		uses a wordlist in PATH (ignores other arguments)\n------------------------------------------------------------------------------------------\n"
	args[0] = "Hash-Cracker"

	for n, element := range args {
		if element == "-h" {
			fmt.Println(help)
			os.Exit(0)

		} else if element == "-t" {
			type_ = args[n+1]

		} else if element == "-p" {
			password = args[n+1]

		} else if element == "-w" {
			path = args[n+1]
			islist = true
		} else if element == "-l" {
			included[0] = args[n]
		} else if element == "-L" {
			included[1] = args[n]
		} else if element == "-n" {
			included[0] = args[n]
		} else if element == "-s" {
			included[0] = args[n]
		} else if element == "-m" {
			included[4] = args[n+1]
		} else if element == "-M" {
			included[5] = args[n+1]
		} else if element != "Hash-Cracker" {
			fmt.Printf("Unknown flag: %v \n", element)
		}

	}
	if islist {
		wordlist(password, type_, path)
		os.Exit(0)
	} else {
		brute(included)
	}
}

func brute(args [6]string) string {
	for n, arg := range args {
		fmt.Printf("%v: %v", n, arg)
	}
	return ""
}

func wordlist(password string, type_ string, path string) {
	fmt.Println("Starting Wordlist mode...")
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
	} else if form == "sha1" {
		hash := sha1.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	}
	return ""
}
