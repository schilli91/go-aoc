package day_6

import (
	"fmt"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 6
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile(fmt.Sprintf("pkg/2022/day_%d/puzzle_input.txt", day))
	defer input.Close()
	s := common.ReadLinesString(input)[0]

	c := findPacketMarker(s)
	fmt.Printf("The packets start after %d characters\n", c)

	o := findMessageMarker(s)
	fmt.Printf("The packets start after %d characters\n", o)
}

func findPacketMarker(dataStream string) int {
	return findMarker(dataStream, 4)
}

func findMessageMarker(dataStream string) int {
	return findMarker(dataStream, 14)
}

func findMarker(dataStream string, numberOfUniques int) int {
	pos := numberOfUniques
	candidate := dataStream[:pos]
	for hasDuplicates(candidate) {
		pos++
		candidate = dataStream[pos-numberOfUniques : pos]
	}

	return pos
}

func hasDuplicates(marker string) bool {
	set := make(map[rune]rune)
	for _, i := range marker {
		set[i] = i
	}
	return len(set) != len(marker)
}
