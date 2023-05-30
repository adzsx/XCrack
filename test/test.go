package test

import (
	"log"
	"os"
	"strings"

	"github.com/adzsx/xcrack/pkg/crack"
	"github.com/adzsx/xcrack/pkg/format"
	"github.com/adzsx/xcrack/pkg/list"
)

var (
	inp   string
	query format.Query
)

func TestAll() {
	log.Println("Starting tests...\n\n ")

	// Wordlist generation
	log.Println("Testing list mode")
	inp = "xcrack list -o ./tempWlist1.txt -M 4 -m 2 -l -n"
	query = format.Args(strings.Split(inp, " "))
	list.WgenSetup(query, false)

	inp = "xcrack list -o ./tempWlist2.txt -M 3 -n -L"
	query = format.Args(strings.Split(inp, " "))
	list.WgenSetup(query, false)

	//Wordlist merging and cleaning
	inp = "xcrack list -w ./tempWlist1.txt -w ./tempWlist2.txt -o ./tempWlist.txt"
	query = format.Args(strings.Split(inp, " "))
	list.WlistClean(query)

	log.Println("Passed list mode test\n ")

	// Hash cracking
	log.Println("Testing cracking mode")
	inp = "xcrack crack -p a94a8fe5ccb19ba61c4c0873d391e987982fbbd3 -t sha1 -l -M 4"
	query = format.Args(strings.Split(inp, " "))
	password, _ := crack.BruteSetup(query)
	if password != "test" {
		log.Fatalf("Expected \"test\", got \"%v\". Brute force cracking", password)
	}

	inp = "xcrack crack -p a94a8fe5ccb19ba61c4c0873d391e987982fbbd3 -t sha1 -w tempWlist.txt -w tempWlist1.txt"
	query = format.Args(strings.Split(inp, " "))
	password, _ = crack.WlistSet(query)
	if password != "test" {
		log.Fatalf("Expected \"test\", got \"%v\". Wordlist cracking", password)
	}

	log.Println("Passed crack mode test\n ")

	// Hash generation
	log.Println("Testing hashing mode")
	inp = "xcrack hash -p test -t sha1"
	query = format.Args(strings.Split(inp, " "))
	hash := crack.Hash(query.Password, query.Hash)
	if hash != "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" {
		log.Fatalf("Expected \"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3\", got \"%v\". Hash generation", hash)
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
