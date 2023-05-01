package check

import (
	"fmt"
	"os"
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
		panic(err)
	}
}

func ErrPrint(err error, text string, exit bool) {
	if err != nil {
		fmt.Println(text)

		if exit {
			os.Exit(0)
		}
	}
}

func Sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
