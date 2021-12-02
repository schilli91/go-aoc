package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadLinesInt(measurements *os.File) []int {
	var lines []int
	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func ReadLinesString(measurements *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		i := scanner.Text()
		lines = append(lines, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
