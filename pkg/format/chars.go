package format

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adzsx/xcrack/pkg/check"
)

var (
	chars []string

	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	numbers = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	special = [39]string{" ", "^", "\u00b4", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", ".", "~", "’", "–", "·"}
)

// final = [mode, password, path, chars, hash, min, max]

type Query struct {
	// Modes: help, crack, list, hash. test
	Mode     string
	Password string

	// input/output for paths
	Inputs []string
	Output string

	Chars []string

	// Hash modes: md5, sha1, sha256
	Hash string

	Min int
	Max int

	// Verbose bool
}

func Args(cmdIn []string) Query {
	// final = [mode, password, path, chars, hash, min, max]
	query := Query{}
	var output string

	var lists []string

	if check.InSclice(cmdIn, "help") || check.InSclice(cmdIn, "-h") || check.InSclice(cmdIn, "--help") {
		query.Mode = "help"
		return query
	} else if cmdIn[1] == "crack" && query.Mode == "" {
		query.Mode = "crack"

	} else if cmdIn[1] == "list" && query.Mode == "" {
		query.Mode = "list"
	} else if cmdIn[1] == "hash" && query.Mode == "" {
		query.Mode = "hash"
	} else if cmdIn[1] == "test" && query.Mode == "" {
		query.Mode = "test"

		return query
	}

	for index, element := range cmdIn {
		if element[0:1] == "-" {
			switch element[1:2] {

			case "p":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the password")
					os.Exit(0)
				}

				query.Password = cmdIn[index+1]

			case "t":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the type")
					os.Exit(0)
				}

				query.Hash = cmdIn[index+1]

			case "c":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the characters")
					os.Exit(0)
				}

				chars = append(chars, strings.Join(cmdIn[index+1:index+2], ""))

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

			case "o":
				output = cmdIn[index+1]

			case "w":
				lists = append(lists, cmdIn[index+1])

			case "m":

				min, err := strconv.Atoi(cmdIn[index+1])

				if err != nil {
					fmt.Printf("Failed to convert %v to int", cmdIn[index+1])
				}

				query.Min = min

			case "M":
				max, err := strconv.Atoi(cmdIn[index+1])

				if err != nil {
					fmt.Printf("Failed to convert %v to int", cmdIn[index+1])
				}

				query.Max = max
			}
		}
	}

	if query.Hash == "" {
		query.Hash = "md5"
	}

	if query.Mode == "" {
		query.Mode = "crack"
	}

	query.Chars = chars

	if query.Mode == "list" {
		if len(lists) > 0 {
			query.Output = output
			query.Inputs = lists
		} else {
			query.Output = output
		}

	} else {
		query.Inputs = lists
	}

	if query.Min == 0 {
		query.Min = 1
	}
	if query.Max == 0 {
		query.Max = 8
	}

	if len(cmdIn) < 3 && query.Mode != "help" {
		fmt.Println("Enter -h for help\n ")
		os.Exit(0)
	}

	if query.Mode == "list" && query.Output == "" {
		if len(chars) > 0 {
			fmt.Println("Pleace specify the output file")
		} else {
			fmt.Println("Please specify the input and output files")

			os.Exit(1)
		}
		return query
	}

	return query
}
