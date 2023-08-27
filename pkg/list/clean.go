package list

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/adzsx/xcrack/pkg/utils"
)

var (
	items []string
)

func WlistClean(query utils.Input) time.Duration {
	now := time.Now()

	for _, file := range query.Inputs {
		readList(file)
	}

	items = rmDupl(items)

	os.Remove(query.Output)

	outfile, err := os.Create(query.Output)

	utils.Err(err)

	for _, item := range items {
		_, _ = io.WriteString(outfile, item+"\n")
	}

	return time.Since(now)
}

func readList(fileName string) {
	_, err := os.ReadFile(fileName)

	utils.Err(err)

	file, err := os.Open(fileName)

	utils.Err(err)

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
