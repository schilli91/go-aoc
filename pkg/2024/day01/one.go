package day01

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/schilli91/go-aoc/pkg/2024/datafile"
)

func Run() {
	day := "01"
	fmt.Printf("Day %s of AOC 2024.\n", day)

	input := datafile.ReadLinesString(fmt.Sprintf("pkg/2024/day%s/puzzle_input", day))

	o := calcDistance(input)
	fmt.Printf("The value of part one is %d\n", o)

	t := -1
	fmt.Printf("The value of part two is %d\n", t)
}

func calcSimilarity() int {

	return -1
}

func calcDistance(lines []string) int {
	l, r := getDistanceLists(lines)

	// fmt.Printf("l (unsorted): %v\n", l)
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	// fmt.Printf("l (sorted):   %v\n", l)

	// fmt.Printf("r (unsorted): %v\n", r)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	// fmt.Printf("r (sorted):   %v\n", r)

	var distances []int
	for i := range l {
		d := l[i] - r[i]
		if d < 0 {
			d = d * -1
		}
		distances = append(distances, d)
	}

	sum := 0
	for _, d := range distances {
		sum += d
	}
	return sum
}

func getDistanceLists(lines []string) ([]int, []int) {
	var l, r []int
	for _, row := range lines {
		parts := strings.Split(row, " ")
		lVal, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Could not convert string '%s' to int: %v\n", parts[0], err)
		}
		rVal, err := strconv.Atoi(parts[len(parts)-1])
		if err != nil {
			log.Fatalf("Could not convert string '%s' to int: %v\n", parts[len(parts)-1], err)
		}
		l = append(l, lVal)
		r = append(r, rVal)
	}
	return l, r
}
