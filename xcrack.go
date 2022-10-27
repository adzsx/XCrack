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
	//some strings to display
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

	//mode (help message)
	mode string

	//storing for flags in declaration of app (eg. -l -s, -m x -M y)
	//password
	hashed string

	//flags
	flags [6]string

	//type of hash
	type_      string
	isWordlist bool
	path       string

	//characters for brute force mode
	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers   = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	special   = [39]string{" ", "^", "´", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", "¸", "~", "’", "–", "·"}

	//Characters defined by flags
	chars []string
)

// setup and checking for arguments
func main() {
	fmt.Println(start)
	args := os.Args
	args[0] = "Hash-Cracker"

	// check for command line arguments
	for index, element := range args {
		switch element {
		case "-h":
			fmt.Println(help)
			os.Exit(0)
		case "-t":
			type_ = args[index+1]
		case "-p":
			hashed = args[index+1]
		case "-w":
			path = args[index+1]
			isWordlist = true
		case "-l":
			flags[0] = args[index]
		case "-L":
			flags[1] = args[index]
		case "-n":
			flags[2] = args[index]
		case "-s":
			flags[3] = args[index]
		case "-m":
			flags[4] = args[index+1]
		case "-M":
			flags[5] = args[index+1]
		}
	}

	//start wordlist mode when -w is given
	if isWordlist {
		wordlist(hashed, type_, path)
		os.Exit(0)
	} else if len(args) > 1 {
		fmt.Println(len(args), args)
		fmt.Println(brute_force(flags, hashed))
	} else {
		fmt.Println("Enter -h for help\n ")
		os.Exit(0)
	}
}

// starting wordlist mode
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

// setting up wordlist mode
func brute_force(args [6]string, password string) string {
	if contains(args, "-n") {
		for _, v := range numbers {
			chars = append(chars, v)
		}
	}
	if contains(args, "-l") {
		for _, v := range l_letters {
			chars = append(chars, v)
		}
	}
	if contains(args, "-L") {
		for _, v := range u_letters {
			chars = append(chars, v)
		}
	}
	if contains(args, "-s") {
		for _, v := range special {
			chars = append(chars, v)
		}
	}
	jobs := make(chan int, 50)
	result := make(chan string, 1)

	go brute(chars, password, jobs, result)
	go brute(chars, password, jobs, result)
	go brute(chars, password, jobs, result)
	go brute(chars, password, jobs, result)

	for i := 1; i <= 50; i++ {
		jobs <- i
	}
	close(jobs)

	return <-result
}

func brute(chars []string, hashed string, jobs <-chan int, result chan<- string) {
	//chars = characters for password
	//hashed = hashed password to crack
	//length = length of characters in chars
	//jobs = jobs for lengths for multiple gorutines
	//result = channel to send password if found

	fmt.Println("Starting brute force mode")
	counter := make([]int, len(chars))
	counter[0] = -1
	for currentLength := range jobs {
		fmt.Printf("Starting with length: %v\n%v\n", currentLength, chars)

	}

	result <- hashed
}

// hashing function
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

// in list checker
func contains(s [6]string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}

// error checker
func check(err error) {
	if err != nil {
		panic(err)
	}
}
