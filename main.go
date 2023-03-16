package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adzsx/xcrack/pkg/check"
	"github.com/adzsx/xcrack/pkg/crack"
	"github.com/adzsx/xcrack/pkg/format"
	"github.com/adzsx/xcrack/pkg/list"
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

	sets := format.Args(args)
	// new = [mode, password, path, chars, hash, min, max]

	if sets[0] == "help" {
		fmt.Println(help)

	} else if sets[0] == "crack" {
		min, err := strconv.Atoi(sets[5])
		check.Err(err)

		max, err := strconv.Atoi(sets[6])
		check.Err(err)

		if sets[2] != "" {
			crack.WlistSet(sets[1], sets[4], strings.Split(sets[2], " "))
		} else {
			crack.BruteSetup(sets[1], sets[4], strings.Split(sets[3], ""), min, max)
		}

	} else if sets[0] == "list" {
		min, err := strconv.Atoi(sets[5])
		check.Err(err)

		max, err := strconv.Atoi(sets[6])
		check.Err(err)

		paths := strings.Split(sets[2], " ")
		output := paths[0]
		paths = paths[1:]

		if sets[2] == "" {
			list.WlistClean(paths, output)
		} else {
			fmt.Println(strings.Split(sets[3], ""))
			list.WgenSetup(strings.Split(sets[3], ""), output, min, max)
		}

	} else if sets[0] == "hash" {
		fmt.Printf("\n\"%v\" (%v):			%v\n", sets[1], sets[4], crack.Hash(sets[1], sets[4]))
	}
}
