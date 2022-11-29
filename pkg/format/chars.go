package format

import (
	"strings"

	"github.com/adzsx/xcrack/pkg/check"
)

var (
	chars []string

	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers   = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	special   = [39]string{" ", "^", "\u00b4", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", ".", "~", "’", "–", "·"}
)

func charArr(args []string) []string {
	if check.InSclice(args, "n") {
		for _, v := range numbers {
			chars = append(chars, v)
		}
	}
	if check.InSclice(args, "l") {
		for _, v := range l_letters {
			chars = append(chars, v)
		}
	}
	if check.InSclice(args, "L") {
		for _, v := range u_letters {
			chars = append(chars, v)
		}
	}
	if check.InSclice(args, "s") {
		for _, v := range special {
			chars = append(chars, v)
		}
	}
	return chars
}

func Args(cmdIn []string) [6]string {
	// final = [mode, password, hash, chars/path, min, max]
	var final [6]string
	var lists []string
	modeCount := 0

	if check.InSclice(cmdIn, "help") || check.InSclice(cmdIn, "-h") || check.InSclice(cmdIn, "--help") && modeCount < 1 {
		final[0] = "help"
		return final
	} else if check.InSclice(cmdIn, "hash") && modeCount < 1 {
		final[0] = "hash"
		modeCount++
	} else if check.InSclice(cmdIn, "gen") && modeCount < 1 {
		final[0] = "gen"
		modeCount++
	}

	for index, element := range cmdIn {
		if element[0:1] == "-" && len(element) > 2 {
			flags := strings.Split(element[1:], "")

			chars = charArr(flags)
		} else if element[0:1] == "-" {
			switch element[1:2] {
			case "p":
				final[1] = cmdIn[index+1]

			case "t":
				final[2] = cmdIn[index+1]

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

			case "w":
				lists = append(lists, cmdIn[index+1])

			case "m":
				final[4] = cmdIn[index+1]

			case "M":
				final[5] = cmdIn[index+1]

			case "c":
				chars = append(chars, strings.Join(cmdIn[index+1:index+2], ""))
			}
		}
	}
	if final[2] == "" {
		final[2] = "md5"
	}

	if final[0] == "" {
		final[0] = "hash"
	}

	if len(chars) == 0 {
		for _, char := range l_letters {
			chars = append(chars, char)
		}

		for _, char := range numbers {
			chars = append(chars, char)
		}
	}

	if len(lists) < 1 {
		final[3] = strings.Join(chars, "")
	} else {
		final[3] = strings.Join(lists, ",")
		final[0] = "list"
	}

	if final[4] == "" {
		final[4] = "1"
	}
	if final[5] == "" {
		final[5] = "8"
	}

	return final
}
