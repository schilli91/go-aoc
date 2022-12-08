package day_3

import (
	"fmt"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 1
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile("/Users/moritzschillinger/dev/go/advent-of-code/pkg/2022/day_3/puzzle_input.txt")
	defer input.Close()

	p := getPriorityOfSharedItems(common.ReadLinesString(input))
	fmt.Printf("The total priority of shared item types is %d.\n", p)
}

type Rucksack struct {
	firstCompartment  string
	secondCompartment string
}

func NewRucksack(content string) *Rucksack {
	r := new(Rucksack)
	r.firstCompartment = content[:len(content)/2]
	r.secondCompartment = content[len(content)/2:]
	return r
}

func (r Rucksack) findSharedItemType() []rune {
	common := make([]rune, 0)
	for _, i := range r.firstCompartment {
		for _, j := range r.secondCompartment {
			if i == j && !contains(common, i) {
				common = append(common, i)
			}
		}
	}
	return common
}

func contains(haystack []rune, needle rune) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
}

func getItemPriority(items []rune) int {
	p := 0

	for _, i := range items {
		k := int(i - '0')
		if 17 <= k && k <= 42 {
			// rune in A - Z
			k += 10
		} else if 49 <= k && k <= 74 {
			// rune in a - z
			k -= 48
		}
		p += k
	}
	return p
}

func getPriorityOfSharedItems(rucksackContents []string) int {
	prio := 0
	for _, content := range rucksackContents {
		r := NewRucksack(content)
		s := r.findSharedItemType()
		p := getItemPriority(s)
		prio += p
	}
	return prio
}
