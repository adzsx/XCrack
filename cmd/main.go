package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adzsx/xcrack/internal/crack"
	"github.com/adzsx/xcrack/internal/list"
	"github.com/adzsx/xcrack/internal/test"
	"github.com/adzsx/xcrack/internal/utils"
)

var (
	help string = `Xcrack usage:
	xcrack [mode] [flags]

Modes:
	crack:	Cracks a given hash with either a wordlist or brute force attack (default)
	list:	Generated a wordlist based on your preferences
	hash:	Generates a hash from a given string

Flags:
	-p, --password 		[string]:   	hashed password/text to be hashed 					
	-w, --wordlist 		[string]:	input wordlist
	-f, --file		[path]:		Crack hashes form file
	-t, --type 		([string]):	hash type, hash detection if left empty
	
	-o, --output 		[string]:	output wordlist

	-c, --characters	[string]:	use CHARS for the password
	-l, --lletters:				lowercase letters				default
	-L, --uletters:				uppercase letters
	-n, --numbers:				numbers						default
	-s, --special:				special Characters
	
	-m, --min 		[int]:		min LENGTH of password				default: 3
	-M, --max 		[int]:		max LENGTH of password 				default: 8
	
	-a, --async		[int]:		number of threads to use

hash-mode:
	-r, --raw: 				Print just the hash(es)
	-d, --detect:				detect the hash type
`

	version = "xcrack v1.2"
)

func main() {
	args := os.Args
	args[0] = "xcrack"

	if len(args) < 2 {
		fmt.Println("Enter -h for help\n ")
		os.Exit(0)
	}

	input := utils.Args(args)
	// input = mode, password, input/output files, hash type, min length,

	if input.Mode == "help" {
		fmt.Println(help)
		os.Exit(0)
		// Crack, if possible with wordlist
	}

	fmt.Printf(`Mode:		%v
Password:	%v
File:		%v
Threads:	%v

Wordlist:	%v
Output: 	%v
Chars:		%v
Hash:		%v
Length:		(%v-%v)

`, input.Mode, input.Password, input.File, input.Threads, strings.Join(input.Inputs, ", "), input.Output, input.Chars, input.Hash, input.Min, input.Max)

	if input.Mode == "version" {
		fmt.Println(version)

	} else if input.Mode == "crack" {
		if len(input.Inputs) == 0 {
			fmt.Println("Starting Brute force mode with hash type " + input.Hash + "...")
			pw, time := crack.BruteSetup(input)

			fmt.Printf("\nPassword: \"%v\"\n[%v]", pw, time)
		} else {
			pw, time := crack.WlistSet(input)

			fmt.Printf("\nPassword: \"%v\"\n[%v]", pw, time)
		}

		// List mode (Generate, clean or merge) wordlists
	} else if input.Mode == "list" {
		if len(input.Chars) == 0 {
			list.WlistClean(input)
		} else {
			list.WgenSetup(input, true)
		}

		// Hash mode (Generate hashes )
	} else if input.Mode == "hash" {
		for _, hash := range strings.Split(input.Hash, " ") {
			if !input.Raw {
				fmt.Printf("\n\"%v\" (%v):			%v\n", input.Password, hash, crack.Hash(input.Password, hash))
			} else {
				fmt.Println(crack.Hash(input.Password, hash))
			}

		}

	} else if input.Mode == "test" {
		test.TestAll()
	}
}
