package test

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/adzsx/xcrack/pkg/crack"
	"github.com/adzsx/xcrack/pkg/list"
	"github.com/adzsx/xcrack/pkg/utils"
	format "github.com/adzsx/xcrack/pkg/utils"
)

var (
	inp   string
	input format.Input
)

func TestAll() {
	log.Println("Starting tests...\n\n ")

	// Wordlist generation
	log.Println("Testing list mode")
	inp = "xcrack list -o ./tempWlist1.txt -M 4 -m 2 -l -n"
	input = format.Args(strings.Split(inp, " "))
	list.WgenSetup(input, false)

	inp = "xcrack list -o ./tempWlist2.txt -M 3 -n -L"
	input = format.Args(strings.Split(inp, " "))
	list.WgenSetup(input, false)

	//Wordlist merging and cleaning
	inp = "xcrack list -w ./tempWlist1.txt -w ./tempWlist2.txt -o ./tempWlist.txt"
	input = format.Args(strings.Split(inp, " "))
	list.WlistClean(input)

	log.Println("Passed list mode test\n ")

	// Hash cracking
	log.Println("Testing cracking mode")
	inp = "xcrack crack -p a94a8fe5ccb19ba61c4c0873d391e987982fbbd3 -t sha1 -l -M 4"
	input = format.Args(strings.Split(inp, " "))
	password, _ := crack.BruteSetup(input)
	if password != "test" {
		utils.Err(errors.New("expected \"test\", got \"" + password + "\". Brute force cracking"))
	}

	inp = "xcrack crack -p a94a8fe5ccb19ba61c4c0873d391e987982fbbd3 -t sha1 -w tempWlist.txt -w tempWlist1.txt"
	input = format.Args(strings.Split(inp, " "))
	password, _ = crack.WlistSet(input)
	if password != "test" {
		utils.Err(errors.New("expected \"test\", got \"" + password + "\". wordlist cracking"))
	}

	log.Println("Passed crack mode test\n ")

	// Hash generation
	log.Println("Testing hashing mode")
	inp = "xcrack hash -p test -t sha512"
	input = format.Args(strings.Split(inp, " "))
	hash := crack.Hash(input.Password, input.Hash)
	if hash != "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff" {
		utils.Err(errors.New("expected \"ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff\", got \"" + hash + "\". Hash generation"))
	}

	log.Println("Passed hash generation test\n\n ")

	log.Println("Tests completed successfully")

	rmFile("tempWlist1.txt")
	rmFile("tempWlist2.txt")
	rmFile("tempWlist.txt")
}

func rmFile(name string) {
	e := os.Remove(name)
	if e != nil {
		panic(e)
	}
}
