package utils

import (
	"fmt"
	"runtime"
)

func InArr(s [6]string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}

func InSclice(s []string, element string) bool {
	for _, v := range s {
		if element == v {
			return true
		}
	}
	return false
}

func Err(err error) {
	if err != nil {
		Ansi("\033[91m")
		fmt.Println(err)
		Ansi("\033[0m")
	}
}

func Ansi(str string) {
	if runtime.GOOS != "windows" {
		fmt.Print(str)
	}
}

func Sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
