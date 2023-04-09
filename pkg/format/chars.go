package format

import (
	"fmt"
	"os"
	"strings"

	"github.com/adzsx/xcrack/pkg/check"
)

var (
	chars  []string
	output string

	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers   = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	special   = [39]string{" ", "^", "\u00b4", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", ".", "~", "’", "–", "·"}
)

/*
final = [mode, password, path, chars, hash, min, max]

modes:

	help:	Shows the help message
	crack:	For cracking passwords (default)
	list:	List operations:
							Wordlist generation
							Wordlist merging
							Wordlist cleaning
	hash:	Generate hashes

password:

	hashed password

path:

	path for wordlist for:
							Cracking
							Generation
							Cleaning

							Merging (Different paths seperated by space, output is 1.)

Chars:

	Chars for cracking/wordlist generation
*/
func Args(cmdIn []string) [7]string {
	// final = [mode, password/test function, path, chars, hash, min, max]
	var final [7]string
	var lists []string

	if check.InSclice(cmdIn, "help") || check.InSclice(cmdIn, "-h") || check.InSclice(cmdIn, "--help") {
		final[0] = "help"
		return final
	} else if cmdIn[1] == "crack" && final[0] == "" {
		final[0] = "crack"

	} else if cmdIn[1] == "list" && final[0] == "" {
		final[0] = "list"
	} else if cmdIn[1] == "hash" && final[0] == "" {
		final[0] = "hash"
	} else if cmdIn[1] == "test" && final[0] == "" {
		final[0] = "test"
		if len(cmdIn) > 2 {
			final[1] = cmdIn[2]
		}

		return final
	}

	for index, element := range cmdIn {
		if element[0:1] == "-" {
			switch element[1:2] {

			case "p":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the password")
					os.Exit(0)
				}

				final[1] = cmdIn[index+1]

			case "t":
				if len(cmdIn) <= index+1 {
					fmt.Println("Please specify the type")
					os.Exit(0)
				}

				final[4] = cmdIn[index+1]

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
				final[5] = cmdIn[index+1]

			case "M":
				final[6] = cmdIn[index+1]
			}
		}
	}

	if final[4] == "" {
		final[4] = "md5"
	}

	if final[0] == "" {
		final[0] = "crack"
	}

	if len(chars) == 0 && final[0] != "list" {
		for _, char := range l_letters {
			chars = append(chars, char)
		}

		for _, char := range numbers {
			chars = append(chars, char)
		}
	}

	final[3] = strings.Join(chars, "")

	if final[0] == "list" {
		if len(lists) > 0 {
			final[2] = output + " " + strings.Join(lists, " ")
		} else {
			final[2] = output
		}

	} else {
		strings.Join(lists, " ")
	}

	if final[5] == "" {
		final[5] = "1"
	}
	if final[6] == "" {
		final[6] = "8"
	}

	if len(cmdIn) < 3 && final[0] != "help" {
		fmt.Println("Enter -h for help\n ")
		os.Exit(0)
	}

	if final[0] == "list" && final[2] == "" {
		if len(chars) > 0 {
			fmt.Println("Pleace specify the output file")
		} else {
			fmt.Println("Please specify the input and output files")
		}
		os.Exit(1)
	}
	return final
}
