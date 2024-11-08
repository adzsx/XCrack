package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	numbers = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	special = [39]string{" ", "^", "\u00b4", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", ".", "~", "’", "–", "·"}
)

// final = [mode, password, path, chars, hash, min, max]

type Input struct {
	// Modes: help, crack, list, hash. test
	Mode     string
	Password string
	File     string
	Threads  string

	// input/output for paths
	Inputs []string
	Output string

	Chars []string

	// Hash modes: md5, sha1, sha256, sha512
	Hash string

	Min int
	Max int

	Raw bool

	// Verbose bool
}

func Args(cmdIn []string) Input {
	// final = [mode, password, path, chars, hash, min, max]

	input := Input{}

	var output string
	var chars []string
	var lists []string

	if InSclice(cmdIn, "help") || InSclice(cmdIn, "-h") || InSclice(cmdIn, "--help") {
		input.Mode = "help"
		return input
	} else if cmdIn[1] == "crack" && input.Mode == "" {
		input.Mode = "crack"

	} else if cmdIn[1] == "list" && input.Mode == "" {
		input.Mode = "list"

	} else if cmdIn[1] == "hash" && input.Mode == "" {
		input.Mode = "hash"

	} else if cmdIn[1] == "--version" && input.Mode == "" {
		input.Mode = "version"
		return input

	} else if cmdIn[1] == "test" && input.Mode == "" {
		input.Mode = "test"
		return input
	}

	for index, element := range cmdIn {
		if element[0:1] == "-" {
			switch element[1:] {

			case "p", "-password":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the password")
					os.Exit(0)
				}

				input.Password = strings.ToLower(cmdIn[index+1])

			case "t", "-type":
				if len(cmdIn) <= index+1 || cmdIn[index+1][0:1] == "-" {
					input.Hash = "auto"
				} else {
					if input.Hash == "" {
						input.Hash = cmdIn[index+1]
					} else {
						input.Hash += " " + cmdIn[index+1]
					}

				}

			case "c", "-characters":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the characters")
					os.Exit(0)
				}

				chars = append(chars, strings.Join(cmdIn[index+1:index+2], ""))

			case "f", "-file":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the file")
					os.Exit(0)
				}

			case "l":
				for _, char := range l_letters {
					chars = append(chars, char)
				}

			case "L":
				for _, char := range u_letters {
					chars = append(chars, char)
				}

			case "n":
				for _, char := range numbers {
					chars = append(chars, char)
				}

			case "s":
				for _, char := range special {
					chars = append(chars, char)
				}

			case "o", "-output":
				output = cmdIn[index+1]

			case "w", "-wordlist":
				lists = append(lists, cmdIn[index+1])

			case "m", "-min":

				min, err := strconv.Atoi(cmdIn[index+1])

				if err != nil {
					fmt.Printf("Failed to convert %v to int", cmdIn[index+1])
				}

				input.Min = min

			case "M", "-max":
				max, err := strconv.Atoi(cmdIn[index+1])

				if err != nil {
					fmt.Printf("Failed to convert %v to int", cmdIn[index+1])
				}

				input.Max = max
			case "r", "-raw":
				input.Raw = true
			}
		}
	}

	if input.Hash == "" {
		input.Hash = GetHashType(input.Password)
		if input.Hash == "" {
			input.Hash = "md5"
		}
	}

	if input.Hash == "auto" {
		input.Hash = GetHashType(input.Password)
	}

	if input.Mode == "" {
		input.Mode = "crack"
	}

	input.Chars = chars

	if input.Mode == "list" {
		if len(lists) > 0 {
			input.Output = output
			input.Inputs = lists
		} else {
			input.Output = output
		}

	} else {
		input.Inputs = lists
	}

	if input.Min == 0 {
		input.Min = 3
	}
	if input.Max == 0 {
		input.Max = 8
	}

	if len(cmdIn) < 3 && input.Mode != "help" {
		fmt.Println("Enter -h for help\n ")
		os.Exit(0)
	}

	if input.Mode == "list" && input.Output == "" {
		if len(chars) > 0 {
			fmt.Println("Pleace specify the output file")
		} else {
			fmt.Println("Please specify the input and output files")

			os.Exit(1)
		}
		return input
	}

	if InSclice(cmdIn, "-d") || InSclice(cmdIn, "--detect") {
		typ := GetHashType(input.Password)

		if typ == "" {
			Err(errors.New("unable to detect hash type"))
			os.Exit(0)
		}

		fmt.Println("Possible hash type: " + typ)
		os.Exit(0)
	}

	return input
}
