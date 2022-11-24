package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adzsx/xcrack/pkg/list"
)

var (
	// some strings to display
	help string = `Xcrack
a tool for offline password attacks and functions


Modes:
-------------------------------------------------------------------------------------

hash:   Cracks a given hash with either a wordlist or brute force attack (default)
list:   Generated a wordlist based on your preferences
gen:    Generates a hash from a given string
file:	Combine wordlists and generate a new list, with duplicates removed

(mode must be the first argument)

-------------------------------------------------------------------------------------




hash mode:
-------------------------------------------------------------------------------------

Syntax:		xcrack (hash) <HASH> <TYPE> [OPTIONS]


HASH:	   		Specify the hashed password 	(required)
TYPE:			specify the hash-TYPE 			(default: md5)

Options:

	-n:			numbers							(default)
 	-l:			lowercase letters				(default)
	-L:			uppercase letters
	-s:			special Characters
	-c CHARS:	Only uses CHARS for the password
 
	-m LENGTH:	min LENGTH of password			(default: 1)
	-M LENGTH:	max LENGTH of password 			(default: 8)

	-w PATH:	uses a wordlist in PATH instead of character preferences

-------------------------------------------------------------------------------------




list mode:
-------------------------------------------------------------------------------------

Syntax:        xcrack list <path> [OPTIONS]

PATH:	The location where the list is created		 	(required)
		(if lsit exists, new element will be appended)

Options:
	-n:    numbers										(default)
	-l:    lowercase letters							(default)
	-L:    uppercase letters					
	-s:    special Characters
	-c CHARS:	Only uses CHARS for the password

	-m LENGTH:   min LENGTH of password					(default: 1)
	-M LENGTH:   max LENGTH of password					(default: 8)

-------------------------------------------------------------------------------------




gen mode:
-------------------------------------------------------------------------------------

Syntax:		xcrack gen [OPTIONS]

Options:
	-t TYPE:   Specifies the type of the hash 				(default: md5)
	STRING:    Every other argument will be hashed with the specified TYPE

-------------------------------------------------------------------------------------




file mode:
-------------------------------------------------------------------------------------

Syntax:		xcrack file <OUTPUT-FILE> <FILE1> <FILE2> <...>

OUTPUT-FILE: 	Path to file where the new wordlist will be stored
				(If path already exists the file will be overwritten)

FILE:			File with elements to be sorted

-------------------------------------------------------------------------------------

`

	start string = `
############################################

▀▄▀ █▀▀ █▀█ █▀█ █▀▀ █▄▀
█ █ █▄▄ █▀▄ █▀█ █▄▄ █ █

############################################
`

	// storing for flags in declaration of app (eg. -l -s, -m x -M y)
	// password
	hashed string

	// Arguments entered in the command line
	flags [6]string

	// some arguments
	mode       string = "hash"
	type_      string = "md5"
	types      []string
	isWordlist bool   = false
	path       string = "./wordlist.txt"
	toHash     []string
	files      []string

	// characters for brute force mode
	L_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	U_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	Numbers   = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	Special   = [39]string{" ", "^", "´", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", ".", "~", "’", "–", "·"}

	// Characters defined by flags
	Chars []string
)

var Now = time.Now()

// setup and checking for arguments
func main() {
	fmt.Println(start)
	args := os.Args
	args[0] = "Hash-Cracker"
	args = append(args, "")

	if args[1] != "hash" && args[1] != "list" && args[1] != "gen" && args[1] != "file" && args[1] != "-h" {
		mode = "hash"
	} else if args[1] == "-h" {
		fmt.Println(help)
		os.Exit(0)
	} else {
		mode = args[1]
	}

	// specifies the default length
	flags[4] = "1"
	flags[5] = "8"

	switch mode {
	case "hash":

		// check for command line arguments
		hashed = args[1]

		if args[2] == "sha256" || args[2] == "md5" || args[2] == "sha1" {
			type_ = args[2]
		}

		for index, element := range args {
			switch element {
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

		if flags[0] == "" && flags[1] == "" && flags[2] == "" && flags[3] == "" {
			flags[0] = "-l"
			flags[2] = "-n"
		}

		// Display error message when hashed or type_ if not given
		if len(args) < 2 {
			fmt.Println("Enter -h for help")
			os.Exit(0)
		}
		if hashed == "" {
			fmt.Printf("You need to specify the hashed password\n")
			os.Exit(0)
		}

		// start wordlist mode when -w is given
		if isWordlist {
			wordlist(hashed, type_, path)
			fmt.Printf("\n[%v]\n", time.Since(Now))
		} else if len(args) > 3 {
			brute_force(flags, hashed, type_)
		} else {
			fmt.Println("Enter -h for help")
		}

	case "list":
		path = args[1]
		for index, element := range args {
			switch element {
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

		if flags[0] == "" && flags[1] == "" && flags[2] == "" && flags[3] == "" {
			flags[0] = "-l"
			flags[2] = "-n"
		}

		if path == "" {
			fmt.Println("You need to enter a path")
		} else {
			list.WgenSetup(flags, path)
		}

	case "gen":

		for index, element := range args {
			if element == "-t" {
				types = append(types, args[index+1])
			}
		}
		for index, element := range args {
			if index > 1 && element != "-t" && args[index-1] != "-t" {
				toHash = append(toHash, element)
			}
		}

		if len(types) == 0 {
			types = append(types, "md5")
		}

		for _, str := range toHash {
			for _, type_ := range types {
				if str != "" {
					fmt.Printf("\"%v\" (%v):    		%v\n", str, type_, hash(str, type_))
				}
			}
		}

		fmt.Printf("\n[%v]\n", time.Since(Now))

	case "file":
		output := args[2]
		for _, j := range args[2:] {
			files = append(files, j)
		}
		fmt.Printf("Output: %v, Input: %v", output, files)
	}
}

// starting wordlist mode
func wordlist(password string, type_ string, path string) {
	fmt.Println("Starting wordlist mode")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Path \"%v\" found. Plase enter a valid path!\n", path)
		fmt.Printf("\n[%v]\n", time.Since(Now))
		os.Exit(1)
	} else {
		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			data := fileScanner.Text()
			if hash(data, type_) == password {
				fmt.Printf("Password: %v \n", data)
				fmt.Printf("\n[%v]\n", time.Since(Now))
				os.Exit(0)
			}
		}
		fmt.Println("Password not found")
	}
	file.Close()
}

// setting up brute force mode
func brute_force(args [6]string, password string, type_ string) {
	fmt.Println("Starting brute force mode")
	min, err := strconv.Atoi(args[4])
	max, err2 := strconv.Atoi(args[5])

	check(err)
	check(err2)

	if contains(args, "-n") {
		for _, v := range Numbers {
			Chars = append(Chars, v)
		}
	}
	if contains(args, "-l") {
		for _, v := range L_letters {
			Chars = append(Chars, v)
		}
	}
	if contains(args, "-L") {
		for _, v := range U_letters {
			Chars = append(Chars, v)
		}
	}
	if contains(args, "-s") {
		for _, v := range Special {
			Chars = append(Chars, v)
		}
	}

	if min > max {
		fmt.Println("min length cant be longer than max length!")
		os.Exit(1)
	}

	// chars: all chars used in password
	// password: hashed password
	// type_: type of hash
	// jobs: length to generate password

	jobs := make(chan int, max-min)
	result := make(chan bool)

	for i := 0; i < (max - min + 1); i++ {
		go brute(Chars, password, jobs, result)
	}

	for i := min; i <= max; i++ {
		jobs <- i
	}

	close(jobs)

	var finished []bool
	for i := range result {
		finished = append(finished, i)
		if len(finished) >= max-min {
			fmt.Println("Password not found")
			fmt.Printf("\n[%v]\n", time.Since(Now))
			os.Exit(0)
		}
	}
}

// Brute forcer
func brute(chars []string, hashed string, jobs <-chan int, response chan<- bool) {
	// chars = characters for password
	// hashed = hashed password to crack
	// length = length of characters in chars
	// jobs = jobs for lengths for multiple gorutines

	for currentLength := range jobs {
		// if len(jobs) == 0 {
		//  fmt.Println("Password not found! Password probably longer than specified length")
		//  fmt.Printf("\n[%v]\n", time.Since(Now))
		//  os.Exit(1)
		// }
		counter := make([]int, currentLength)
		password := make([]string, currentLength)
		counter[0] = -1
		total := len(counter) * (len(chars) - 1)
		for sum(counter) < total {

			counter[0] += 1

			for index, value := range counter {
				if value > len(chars)-1 {
					counter[index] = 0

					if len(counter) > index+1 {

						counter[index+1] += 1
						continue

					} else {
						break
					}
				}
			}

			for index, value := range counter {
				password[index] = chars[value]
			}
			pw := strings.Join(password[:], "")
			pwh := hash(pw, type_)
			if pwh == hashed {
				fmt.Printf("Password: %v\n", pw)
				fmt.Printf("\n[%v]\n", time.Since(Now))
				os.Exit(0)
			}

		}

	}
	response <- false
}

// hashing function
func hash(text string, type_ string) string {
	switch type_ {
	case "md5":
		hash := md5.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	case "sha1":
		hash := sha1.Sum([]byte(text))
		return hex.EncodeToString(hash[:])
	case "sha256":
		h := sha256.New()
		h.Write([]byte(text))
		hash := h.Sum(nil)
		return fmt.Sprintf("%x", hash)
	}
	return ""
}

// in array checker
func contains(s [6]string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
