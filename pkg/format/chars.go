package format

import (
	"github.com/adzsx/xcrack/pkg/check"
)

var (
	chars []string

	l_letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	u_letters = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	numbers   = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	special   = [39]string{" ", "^", "´", "+", "#", "-", "+", ".", "\"", "<", "°", "!", "§", "$", "%", "&", "/", "(", ")", "=", "?", "`", "*", "'", "_", ":", ";", "′", "{", "[", "]", "}", "\\", ".", "~", "’", "–", "·"}
)

func CharList(args [6]string) []string {
	if check.InArr(args, "-n") {
		for _, v := range numbers {
			chars = append(chars, v)
		}
	}
	if check.InArr(args, "-l") {
		for _, v := range l_letters {
			chars = append(chars, v)
		}
	}
	if check.InArr(args, "-L") {
		for _, v := range u_letters {
			chars = append(chars, v)
		}
	}
	if check.InArr(args, "-s") {
		for _, v := range special {
			chars = append(chars, v)
		}
	}
	return chars
}
