package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/adzsx/xcrack/pkg/crack"
	"github.com/adzsx/xcrack/pkg/list"
	"github.com/adzsx/xcrack/pkg/utils"
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
	-p, --password 		[HASH]:   	hashed password/text to be hashed 					
	-t, --type 		([TYPE]):	hash type, hash detection if left empty		default: md5
	-f, --file		[FILE]:		Crack hashes form file
	-n, --numbers:				numbers						default
 	-l, --lletters:				lowercase letters				default
	-L, --uletters:				uppercase letters
	-s, --special:				special Characters
	-c, --characters	[CHARS]:	use CHARS for the password
	-m, --min 		[LENGTH]:	min LENGTH of password				default: 3
	-M, --max 		[LENGTH]:	max LENGTH of password 				default: 8
	-w, --wordlist 		[PATH]:		input wordlist
	-o, -output 		[PATH]:		output wordlist

Mode Flags:
	hash:
		-r, --raw: 		Prints just the hash(es)
		-d, --detect:	(Tries) to detect the hash type
`

	version = "xcrack v1.3"
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

		// Crack, if possible with wordlist
	} else if input.Mode == "version" {
		fmt.Println(version)

	} else if input.Mode == "crack" {
		if len(input.Inputs) == 0 {
			fmt.Println("Starting Brute force mode with hash type " + input.Hash + "...")
			pw, time := crack.BruteSetup(input)

			fmt.Printf("\nPassword: \"%v\"\n[%v]", pw, time)
		} else {
			fmt.Println("Starting wordlist mode")
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
