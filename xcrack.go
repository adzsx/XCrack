package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	//some strings to display
	help string = `
Xcrack,
a tool for common password attacks

Modes:
-------------------------------------------------------------------------------------

hash:			Cracks a given hash with a wordlist or brute force attack
gen:			Generated a wordlist based on your preferences

(mode has to be the first argument)

-------------------------------------------------------------------------------------


Hash mode:
	Presets:
		-p HASH:		(required) Sets the HASH
		-t HASH-TYPE:		(required) specify the HASH-TYPE: (md5, sha1)

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
		-f PATH:		(required) Stores wordlist in PATH
					If file already exists, password will be appended
					duplicates will not be removed

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
	special   = [39]string{" ", "^", "´", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", ".", "~", "’", "–", "·"}

	//Characters defined by flags
	chars []string
)

var now = time.Now()

// setup and checking for arguments
func main() {
	fmt.Println(start)
	args := os.Args
	args = append(args, "")
	args[0] = "Hash-Cracker"
	flags[4] = "1"
	flags[5] = "50"
	switch args[1] {
	case "hash":

		// check for command line arguments
		for index, element := range args {
			switch element {
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

		//Display error message when hashed or type_ if not given
		if len(args) < 4 {
			fmt.Printf("Enter -h for help\n")
			os.Exit(0)
		}
		if hashed == "" && type_ == "" {
			fmt.Printf("You need to specity the hashed password and the type of the hash\n")
			os.Exit(0)
		} else if hashed == "" {
			fmt.Printf("You need to specify the hashed password\n")
			os.Exit(0)
		} else if type_ == "" {
			fmt.Printf("You need to specify the type of the hash\n")
			os.Exit(0)
		}

		//start wordlist mode when -w is given
		if isWordlist {
			wordlist(hashed, type_, path)
			fmt.Printf("\n[%v]\n", time.Since(now))
		} else if len(args) > 6 {
			brute_force(flags, hashed, type_)
		} else {
			fmt.Println("Enter -h for help")
		}
	case "gen":
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
			case "-p":
				path = args[index+1]
			}
		}
		if path == "" {
			fmt.Println("You need to enter a path")
		} else {
			wgenSetup(flags, path)
		}
	case "-h":
		fmt.Println(help)
	}
}

// starting wordlist mode
func wordlist(password string, type_ string, path string) {
	fmt.Println("Starting Wordlist mode")
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Path \"%v\" found. Plase enter a valid path!\n", path)
		fmt.Printf("\n[%v]\n", time.Since(now))
		os.Exit(1)
	} else {
		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			data := fileScanner.Text()
			if hash(data, type_) == password {
				fmt.Printf("Password: %v \n", data)
				fmt.Printf("\n[%v]\n", time.Since(now))
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

	if min > max {
		fmt.Println("min length cant be longer than max length!")
		os.Exit(1)
	}

	//chars: all chars used in password
	//password: hashed password
	//type_: type of hash
	//jobs: length to generate password

	jobs := make(chan int, max-min)
	result := make(chan bool)

	for i := 0; i < max-min+1; i++ {
		go brute(chars, password, jobs, result)
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
			fmt.Printf("\n[%v]\n", time.Since(now))
			os.Exit(0)
		}
	}

}

// Brute forcer
func brute(chars []string, hashed string, jobs <-chan int, result chan<- bool) {
	//chars = characters for password
	//hashed = hashed password to crack
	//length = length of characters in chars
	//jobs = jobs for lengths for multiple gorutines

	for currentLength := range jobs {
		// if len(jobs) == 0 {
		// 	fmt.Println("Password not found! Password probably longer than specified length")
		// 	fmt.Printf("\n[%v]\n", time.Since(now))
		// 	os.Exit(1)
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
				fmt.Printf("\n[%v]\n", time.Since(now))
				os.Exit(0)
			}

		}

	}
	result <- false

}

func wgenSetup(args [6]string, path string) {
	fmt.Println("Starting wordlist generation mode.")

	//some variables for generating the wordlist
	file, _ := os.Create(path)
	min, err := strconv.Atoi(args[4])
	max, err2 := strconv.Atoi(args[5])

	//check errors in string conversion
	check(err)
	check(err2)

	//Create list with characters included in the password
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

	//fmt.Println(chars)

	//length of passwords to be generated
	jobs := make(chan int, max-min)
	response := make(chan bool, max-min)

	for i := 0; i < max-min+1; i++ {
		go gen(chars, jobs, response, file)
	}

	for i := min; i <= max; i++ {
		jobs <- i
	}

	close(jobs)

	var finished []bool
	for i := range response {
		finished = append(finished, i)
		if len(finished) >= max-min {
			fmt.Println("Done")
			fmt.Printf("\n[%v]\n", time.Since(now))
			os.Exit(0)
		}
	}
}

// Wordlist generation mode
func gen(chars []string, jobs <-chan int, response chan<- bool, file *os.File) {
	//chars = characters for password
	//hashed = hashed password to crack
	//length = length of characters in chars
	//jobs = jobs for lengths for multiple gorutines

	for currentLength := range jobs {
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
			io.WriteString(file, pw+"\n")
		}

	}
	response <- true

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
		bs := h.Sum(nil)
		return string(bs[:])
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
