package list

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/adzsx/xcrack/pkg/check"
)

var (
	items []string
)

func WlistClean(files []string, output string) {
	now := time.Now()

	for _, file := range files {
		readList(file)
	}

	items = rmDupl(items)

	os.Remove(output)

	outfile, err := os.Create(output)

	check.Err(err)

	for _, item := range items {
		_, _ = io.WriteString(outfile, item+"\n")
	}

	fmt.Printf("\n[%v]\n", time.Since(now))
	os.Exit(0)
}

func readList(fileName string) {
	_, err := ioutil.ReadFile(fileName)

	check.Err(err)

	file, err := os.Open(fileName)

	check.Err(err)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		items = append(items, fileScanner.Text())
	}
}

func rmDupl(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
