package test

import (
	"fmt"

	"github.com/adzsx/xcrack/pkg/check"
)

func TestAll() {
	fmt.Println("Running tests...")
	// Tests to be configured to test all funcitons at once.
}

func TestMode(mode string) {
	if !check.InSclice([]string{"crack", "hash", "list", "format"}, mode) {
		fmt.Printf("Mode %v not found", mode)
	} else {
		fmt.Printf("Testing %v mode...", mode)
	}

	if mode == "crack" {
		// crack.BruteSetup()
	}
}

// func rmFile(name string) {
// 	e := os.Remove(name)
// 	if e != nil {
// 		panic(e)
// 	}
// }
