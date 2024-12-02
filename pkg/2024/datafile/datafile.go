package datafile

import (
	"bufio"
	"log"
	"os"
)

func ReadLinesString(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i := scanner.Text()
		lines = append(lines, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
