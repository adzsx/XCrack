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
	help string = `Xcrack
a tool for offline password attacks.


Modes:
-------------------------------------------------------------------------------------

crack:   Cracks a given hash with either a wordlist or brute force attack (default)
list:   Generated a wordlist based on your preferences
hash:    Generates a hash from a given string

-------------------------------------------------------------------------------------




hash mode:
Hash cracking
-------------------------------------------------------------------------------------

Syntax:			xcrack (hash) [flags]


-p HASH:	   	Specify the hashed password 					required
-t TYPE:		specify the hash-TYPE 							default: md5

flags:

	-n:			numbers											default
 	-l:			lowercase letters								default
	-L:			uppercase letters
	-s:			special Characters
	-c CHARS:	Only uses CHARS for the password
 
	-m LENGTH:	min LENGTH of password							default: 1
	-M LENGTH:	max LENGTH of password 							default: 8

	-w PATH:	uses a wordlist in PATH

-------------------------------------------------------------------------------------




list mode:
Wordlist operations
-------------------------------------------------------------------------------------

Syntax:        	xcrack list [flags]

flags:
	-n:    		numbers											
	-l:    		lowercase letters								
	-L:    		uppercase letters						
	-s:    		special Characters
	-c CHARS:	Only uses CHARS for the password

	-m LENGTH:  min LENGTH of password							
	-M LENGTH:  max LENGTH of password							

	-w PATH:	input file at PATH for merning and cleaning
	-o PATH:	output file at PATH for generating, merning and cleaning

-------------------------------------------------------------------------------------




gen mode:
Hash generation
-------------------------------------------------------------------------------------

Syntax:			xcrack gen flags]

flags:
	-t TYPE:	Specifies the type of the hash 					default: md5
	-p STRING:  Argument will be hashed with TYPE

-------------------------------------------------------------------------------------
`

	start string = `
############################################

▀▄▀ █▀▀ █▀█ █▀█ █▀▀ █▄▀
█ █ █▄▄ █▀▄ █▀█ █▄▄ █ █

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

  fmt.Println(query)

	if query.Mode == "help" {
		fmt.Println(help)

		// Crack, if possible with wordlist
	} else if query.Mode == "crack" {
		if len(query.Inputs) == 0 {
			crack.BruteSetup(query)
		} else {
			crack.WlistSet(query)
		}

		// List mode (Generate, clean or merge) wordlists
	} else if query.Mode == "list" {
		if len(query.Chars) == 0 {
			list.WlistClean(query)
		} else {
			list.WgenSetup(query)
		}

		// Hash mode (Generate hashes )
	} else if query.Mode == "hash" {
		fmt.Printf("\n\"%v\" (%v):			%v\n", query.Password, query.Hash, crack.Hash(query.Password, query.Hash))

	} else if query.Mode == "test" {
		test.TestAll()
	}
}
