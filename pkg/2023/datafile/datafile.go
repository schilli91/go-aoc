package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func OpenInputFile(path string) *os.File {
	input, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func ReadLinesInt(f *os.File) []int {
	var lines []int
	scanner := bufio.NewScanner(f)
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

func ReadLinesString(f *os.File) []string {
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

func ReadCommaSeparatedLineInt(f *os.File) []int {
	s := ReadLinesString(f)
	numsStr := strings.Split(s[0], ",")
	nums := []int{}

	for _, f := range numsStr {
		val, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, val)
	}
	return nums
}
