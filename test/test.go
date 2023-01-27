package test

import (
	"fmt"

	"github.com/adzsx/xcrack/pkg/format"
)

func Test() {
	args := []string{"-l", "-L", "-n", "-p", "e1b849f9631ffc1829b2e31402373e3c", "-M", "5"}
	argF := format.Args(args)
	fmt.Println(argF)
}
