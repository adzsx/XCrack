package utils

import (
	"fmt"
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
		fmt.Printf("\033[91m%v", err)
	}
}

func Sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
