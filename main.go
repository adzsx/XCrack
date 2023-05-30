package main

import (
	"fmt"
	"os"

	"github.com/adzsx/xcrack/pkg/crack"
	"github.com/adzsx/xcrack/pkg/format"
	"github.com/adzsx/xcrack/pkg/list"
	"github.com/adzsx/xcrack/test"
)

var (
	help string = `Xcrack usage:
	xcrack [mode] [flags]

Modes:
	crack:	Cracks a given hash with either a wordlist or brute force attack (default)
	list:	Generated a wordlist based on your preferences
	hash:	Generates a hash from a given string

Flags:
	-p HASH:   	hashed password/text to be hashed 					
	-t TYPE:	hash type				default: md5
	-n:		numbers					default
 	-l:		lowercase letters			default
	-L:		uppercase letters
	-s:		special Characters
	-c CHARS:	use CHARS for the password
	-m LENGTH:	min LENGTH of password			default: 1
	-M LENGTH:	max LENGTH of password 			default: 8
	-w PATH:	input wordlist
	-o PATH		output wordlist
`

	start string = `
############################################

▀▄▀ █▀▀ █▀█ █▀█ █▀▀ █▄▀
█ █ █▄▄ █▀▄ █▀█ █▄▄ █ █

############################################
`
)

func main() {
	fmt.Println(start)
	args := os.Args
	args[0] = "xcrack"

	if len(args) < 2 {
		fmt.Println("Enter -h for help\n ")
		os.Exit(0)
	}

	query := format.Args(args)
	// query = mode, password, input/output files, hash type, min length,

	if query.Mode == "help" {
		fmt.Println(help)

		// Crack, if possible with wordlist
	} else if query.Mode == "crack" {
		if len(query.Inputs) == 0 {
			pw, time := crack.BruteSetup(query)

			fmt.Printf("\nPassword: \"%v\"\n[%v]", pw, time)
		} else {
			pw, time := crack.WlistSet(query)

			fmt.Printf("\nPassword: \"%v\"\n[%v]", pw, time)
		}

		// List mode (Generate, clean or merge) wordlists
	} else if query.Mode == "list" {
		if len(query.Chars) == 0 {
			list.WlistClean(query)
		} else {
			list.WgenSetup(query, true)
		}

		// Hash mode (Generate hashes )
	} else if query.Mode == "hash" {
		fmt.Printf("\n\"%v\" (%v):			%v\n", query.Password, query.Hash, crack.Hash(query.Password, query.Hash))

	} else if query.Mode == "test" {
		test.TestAll()
	}
}
