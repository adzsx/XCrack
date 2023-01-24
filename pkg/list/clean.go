package list

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/adzsx/xcrack/pkg/check"
)

var (
	items []string
)

func WlistClean(files []string, output string) {
	outfile, err := os.Create(output)

	check.Err(err)

	_, err = io.WriteString(outfile, "This is a test")

	check.Err(err)

	for _, file := range files {
		readList(file)
	}
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
