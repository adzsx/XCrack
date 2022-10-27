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
	mode string

	hashed   string
	included [6]string

	type_  string
	islist bool
	path   string

	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers   = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	special   = [39]string{" ", "^", "´", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", "¸", "~", "’", "–", "·"}
	chars     [101]string
	length    int

	help string = `
Ptevo Cracker,
a tool for common password attacks

Modes:
-----------------------------------------------------------------------------

hash:			Cracks a given hash with a wordlist or brute force attack
gen:			Generated a wordlist based on your preferences
login:			Cracks a login on a website

(mode has to be the first argument)

-----------------------------------------------------------------------------


Hash mode:
	Presets:
		-p HASH:		(required) Sets the HASH
		-t HASH-TYPE:	(required) specify the HASH-TYPE: (md5, sha1)

	Character preferences:
		-n: 			numbers
		-l: 			lowercase letters
		-L: 			uppercase letters
		-s: 			special Characters

	Length preferences
		-m LENGTH: 		min LENGTH of password
		-M LENGTH: 		max LENGTH of password

	Wordlist Preferences:
		-w PATH:		uses a wordlist in PATH instead of character preferences





Wordlist generation mode:
	Presets:
		-f PATH:		Stores wordlist in PATH, cwd is the default

	Character preferences:
		-n: 			numbers
		-l: 			lowercase letters
		-L: 			uppercase letters
		-s: 			special Characters

	Length preferences
		-m LENGTH: 		min LENGTH of password
		-M LENGTH: 		max LENGTH of password

`

	start string = `
############################################

▀▄▀ █▀▀ █▀█ █▀█ █▀▀ █▄▀
█ █ █▄▄ █▀▄ █▀█ █▄▄ █ █

############################################
`
)

func main() {
	fmt.Println(start)
	args := os.Args
	args[0] = "Hash-Cracker"

	// check for command line arguments
	for n, element := range args {
		switch element {
		case "-h":
			fmt.Println(help)
			os.Exit(0)
		case "-t":
			type_ = args[n+1]
		case "-p":
			hashed = args[n+1]
		case "-w":
			path = args[n+1]
			islist = true
		case "-l":
			included[0] = args[n]
		case "-L":
			included[1] = args[n]
		case "-n":
			included[2] = args[n]
		case "-s":
			included[3] = args[n]
		case "-m":
			included[4] = args[n+1]
		case "-M":
			included[5] = args[n+1]
		}
	}
	//start wordlist mode when -w is true
	if islist {
		wordlist(hashed, type_, path)
		os.Exit(0)
	} else if len(args) > 1 {
		fmt.Println(len(args), args)
		fmt.Println(brute_force(included, hashed))
	} else {
		fmt.Println("Enter -h for help\n ")
		os.Exit(0)
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
	go brute(chars, password, length, jobs, result)
	go brute(chars, password, length, jobs, result)
	go brute(chars, password, length, jobs, result)

	for i := 1; i <= 50; i++ {
		jobs <- i
	}
	close(jobs)

	return <-result
}

func brute(chars [101]string, hashed string, length int, jobs <-chan int, result chan<- string) {
	//chars = characters for password
	//hashed = hashed password to crack
	//length = length of characters in chars
	//jobs = jobs for lengths for multiple gorutines
	//result = channel to send password if found

	fmt.Println("Starting brute force mode")
	for currentLength := range jobs {
		fmt.Printf("Starting with length: %v\n", currentLength)

	}

	result <- hashed
}

func wordlist(password string, type_ string, path string) {
	fmt.Println("Starting Wordlist mode")
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
	switch form {
	case "md5":
		hash := md5.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	case "sha1":
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
