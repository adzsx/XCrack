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
	chars     [101]string
	length    int
)

func main() {
	args := os.Args
	help := "\nBrute forcer cracks hashed passwords:\n------------------------------------------------------------------------------------------\n-h: 			Shows this message (ignores other arguments)\n-p HASH:		(required) Sets the HASH\n-t HASH-TYPE:		(required) specify the HASH-TYPE: (md5, sha1)\n-n: 			numbers\n-l: 			lowercase letters\n-L: 			uppercase letters\n-s: 			special Characters\n-m LENGTH: 		min LENGTH of password\n-M LENGTH: 		max LENGTH of password\n-w PATH:		uses a wordlist in PATH (ignores other arguments)\n------------------------------------------------------------------------------------------\n"
	args[0] = "Hash-Cracker"

	// check for command line arguments
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
			included[2] = args[n]
		} else if element == "-s" {
			included[3] = args[n]
		} else if element == "-m" {
			included[4] = args[n+1]
		} else if element == "-M" {
			included[5] = args[n+1]
		}

	}
	//start wordlist mode when -w is true
	if islist {
		wordlist(password, type_, path)
		os.Exit(0)
	} else {

		fmt.Println(brute_force(included, password))
	}
}

func brute_force(args [6]string, password string) string {
	if contains(args, "-n") {
		for n, v := range numbers {
			chars[n] = v
		}
		length += 10
	}
	if contains(args, "-l") {
		for n, v := range l_letters {
			chars[n+10] = v
		}
		length += 26
	}
	if contains(args, "-L") {
		for n, v := range u_letters {
			chars[n+36] = v
		}
		length += 26
	}
	if contains(args, "-s") {
		for n, v := range special {
			chars[n+62] = v
		}
		length += 39
	}
	jobs := make(chan int, 50)
	result := make(chan string, 1)

	go brute(chars, password, length, jobs, result)

	for i := 1; i <= 50; i++ {
		jobs <- i
	}
	close(jobs)

	return <-result
}

func brute(chars [101]string, password string, length int, jobs <-chan int, result chan<- string) {
	//chars = characters for password
	//password = hashed password to crack
	//length = length of characters in chars
	//jobs = jobs for lengths for multiple gorutines
	//result = channel to send password if found
	for n := range jobs {
		fmt.Println(n)
	}
	result <- password
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

func contains(s [6]string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}
