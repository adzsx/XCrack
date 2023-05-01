package test

import (
	"fmt"
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
	fmt.Println("Running tests...")

	// Testing Wordlist generation
	fmt.Println("Testing wordlist generation mode")

	// First wordlist
	inp = "xcrack list -o ./tempWlist1.txt -M 4 -l"
	query = format.Args(strings.Split(inp, " "))
  list.WgenSetup(query)

	// Second wordlist for merging
	inp = "xcrack list -o ./tempWlist2.txt -M 3 -l -n"
	query = format.Args(strings.Split(inp, " "))
  list.WgenSetup(query)

	// Wordlist merging
	inp = "xcrack list -w ./tempWlist1.txt -w ./tempWlist2.txt -o ./tempWlist.txt"
	query = format.Args(strings.Split(inp, " "))

	rmFile("./tempWlist1.txt")
	rmFile("./tempWlist2.txt")

	// Hash generation testing
	fmt.Println("Testing hash generation:")
	inp = "xcrack hash -p example -t md5"
	query = format.Args(strings.Split(inp, " "))

	fmt.Println(crack.Hash(query.Password, query.Hash))
}

func rmFile(name string) {
	e := os.Remove(name)
	if e != nil {
		panic(e)
	}
}
