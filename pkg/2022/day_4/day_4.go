package day_4

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 4
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile("pkg/2022/day_4/puzzle_input.txt")
	defer input.Close()

	sectionDefinitions := common.ReadLinesString(input)
	c := getNumberOfContainedSections(sectionDefinitions)
	fmt.Printf("There are %d sections contained.\n", c)

	o := getNumberOfOverlappingSections(sectionDefinitions)
	fmt.Printf("There are %d sections overlapping.\n", o)
}

type Section struct {
	start int
	end   int
}

func NewSection(sectionString string) *Section {
	s := new(Section)
	p := strings.Split(sectionString, "-")
	var err error
	s.start, err = strconv.Atoi(p[0])
	if err != nil {
		log.Fatalf("Couldn't init section with string '%s'", sectionString)
	}
	s.end, err = strconv.Atoi(p[1])
	if err != nil {
		log.Fatalf("Couldn't init section with string '%s'", sectionString)
	}
	return s
}

func (s Section) contains(other Section) bool {
	if s.start <= other.start && s.end >= other.end {
		return true
	}
	if other.start <= s.start && other.end >= s.end {
		return true
	}
	return false
}

func (s Section) overlap(other Section) bool {
	if s.contains(other) {
		return true
	}
	if other.start <= s.start && s.start <= other.end {
		return true
	}
	if other.start <= s.end && s.end <= other.end {
		return true
	}
	return false
}

func getNumberOfContainedSections(sectionDefinitions []string) int {
	n := 0
	for _, i := range sectionDefinitions {
		p := strings.Split(i, ",")
		a := NewSection(p[0])
		b := NewSection(p[1])
		if a.contains(*b) {
			n++
		}
	}
	return n
}

func getNumberOfOverlappingSections(sectionDefinitions []string) int {
	n := 0
	for _, i := range sectionDefinitions {
		p := strings.Split(i, ",")
		a := NewSection(p[0])
		b := NewSection(p[1])
		if a.overlap(*b) {
			n++
		}
	}
	return n
}
