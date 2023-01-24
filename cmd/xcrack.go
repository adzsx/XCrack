package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adzsx/xcrack/pkg/check"
	"github.com/adzsx/xcrack/pkg/crack"
	"github.com/adzsx/xcrack/pkg/format"
)

var (
	help string = `Xcrack
a tool for offline password attacks
For the entire documentation visit: https://adzsx.github.io/docs/xcrack


Modes:
-------------------------------------------------------------------------------------

hash:   Cracks a given hash with either a wordlist or brute force attack (default)
list:   Generated a wordlist based on your preferences
gen:    Generates a hash from a given string
file:	Combine wordlists and generate a new list, with duplicates removed

-------------------------------------------------------------------------------------




hash mode:
Hash cracking
-------------------------------------------------------------------------------------

Syntax:			xcrack (hash) [OPTIONS]


-p HASH:	   	Specify the hashed password 					required
-t TYPE:		specify the hash-TYPE 							default: md5

Options:

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

Syntax:        	xcrack list [OPTIONS]

-p PATH:		The location where the list is created		 	required
					New element will be appended

Options:
	-n:    		numbers											default
	-l:    		lowercase letters								default
	-L:    		uppercase letters						
	-s:    		special Characters
	-c CHARS:	Only uses CHARS for the password

	-m LENGTH:  min LENGTH of password							default: 1
	-M LENGTH:  max LENGTH of password							default: 8

	-i PATH:	input file at PATH for merging and cleaning		
	-o PATH:	output file at PATH for merning and cleaning

-------------------------------------------------------------------------------------




gen mode:
Hash generation
-------------------------------------------------------------------------------------

Syntax:			xcrack gen [OPTIONS]

Options:
	-t TYPE:	Specifies the type of the hash 					default: md5
	-p STRING:  Argument will be hashed with TYPE

-------------------------------------------------------------------------------------
`

	start string = `
############################################

▀▄▀ █▀▀ █▀█ █▀█ █▀▀ █▄▀
█ █ █▄▄ █▀▄ █▀█ █▄▄ █ █

############################################
`
)

func main() {
	fmt.Println(start)
	args := os.Args
	args[0] = "Hash-Cracker"

	if check.InSclice(args, "-h") || check.InSclice(args, "help") || check.InSclice(args, "--help") {
		fmt.Println(help)
	}

	sets := format.Args(args)

	// sets = [mode, password, hash type, chars, min, max]

	if sets[0] == "hash" {
		min, err := strconv.Atoi(sets[4])
		check.Err(err)

		max, err := strconv.Atoi(sets[5])
		check.Err(err)

		crack.BruteSetup(sets[1], sets[2], strings.Split(sets[3], ""), min, max)
	} else if sets[0] == "list" {
		crack.WlistSet(sets[1], sets[2], strings.Split(sets[3], ","))
	} else if sets[0] == "gen" {
		fmt.Printf("\n\"%v\" (%v):			%v\n", sets[1], sets[2], crack.Hash(sets[1], sets[2]))
	}
}
