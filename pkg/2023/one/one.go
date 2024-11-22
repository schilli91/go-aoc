package one

import (
	"fmt"
	"log"
	"strconv"

	"github.com/schilli91/go-aoc/pkg/2023/datafile"
)

func Run() {
	day := "one"
	fmt.Printf("Day %s of AOC 2023.\n", day)

	input := datafile.OpenInputFile(fmt.Sprintf("pkg/2023/%s/puzzle_input", day))
	defer input.Close()

	o := DecodeCalibration(datafile.ReadLinesString(input))
	fmt.Printf("The value of part one is %d\n", o)

	t := -1
	fmt.Printf("The value of part two is %d\n", t)
}

func DecodeCalibration(calLines []string) int {
	s := 0
	for _, l := range calLines {
		s += decodeLine(l)
	}
	return s
}

func decodeLine(l string) int {
	a, b := 0, 0

	for _, v := range l {
		val := int(v - '0')
		if 0 <= val && val <= 9 {
			a = val
			break
		}
	}
	for i := range l {
		val := int(l[len(l)-i-1] - '0')
		if 0 <= val && val <= 9 {
			b = val
			break
		}
	}
	i, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	if err != nil {
		log.Fatalf("Error during Atoi: %v\n", err)
	}
	return i
}
