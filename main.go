package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	help := "\nBrute forcer cracks hashed passwords:\n------------------------------------------------------------------------------------------\n-h: 			Shows this message (ignores other arguments)\n-t HASH:		specify the type of HASH: (md5 (Default), sha1, sha256, sha512)\n-n: 			numbers\n-l: 			lowercase letters\n-L: 			uppercase letters\n-s: 			special Characters\n-m LENGTH: 		min LENGTH of password\n-M LENGTH: 		max LENGTH of password\n-w PATH:		uses a wordlist in PATH (ignores other arguments)\n------------------------------------------------------------------------------------------\n"
	args[0] = "Hash-Cracker"
	for n, element := range args {
		if element == "-h" {
			fmt.Println(help)
			break
		} else if element == "-w" {
			list := args[n+1]
			wordlist("test", list)
		}
	}
}

func wordlist(preHash string, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Path \"%v\" found. Plase enter a valid path!\n", path)
		os.Exit(1)
	} else {
		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			data := fileScanner.Text()
			if hash(data) == preHash {
				fmt.Printf("Password: %v \n", data)
				os.Exit(0)
			}
		}
		fmt.Println("Password not found")
	}
	file.Close()
}

func hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
