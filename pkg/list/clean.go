package list

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/adzsx/xcrack/pkg/check"
)

var (
	items []string
)

func WlistClean(files []string, output string) {
	for _, file := range files {
		readList(file)
	}

	items = rmDupl(items)

	err := os.Remove(output)

	fmt.Println(err)

	outfile, err := os.Create(output)

	check.Err(err)

	for _, item := range items {
		_, _ = io.WriteString(outfile, item+"\n")
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
