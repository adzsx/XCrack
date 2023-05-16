package list

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/adzsx/xcrack/pkg/check"
	"github.com/adzsx/xcrack/pkg/format"
)

var (
	items []string
)

func WlistClean(query format.Query) time.Duration {
	now := time.Now()

	for _, file := range query.Inputs {
		readList(file)
	}

	items = rmDupl(items)

	os.Remove(query.Output)

	outfile, err := os.Create(query.Output)

	check.Err(err)

	for _, item := range items {
		_, _ = io.WriteString(outfile, item+"\n")
	}

	return time.Since(now)
}

func readList(fileName string) {
	_, err := os.ReadFile(fileName)

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
