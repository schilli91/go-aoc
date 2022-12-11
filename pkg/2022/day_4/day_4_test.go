package day_4

import (
	"strings"
	"testing"

	"github.com/schilli91/go-aoc/pkg/common"
)

func TestNewSection(t *testing.T) {
	s := NewSection("2-4")
	if s.start != 2 {
		t.Errorf("wrong section start: got %d, want %d", s.start, 2)
	}
	if s.end != 4 {
		t.Errorf("wrong section end: got %d, want %d", s.end, 4)
	}
}

func TestContains(t *testing.T) {
	testData := []struct {
		sectionString string
		isContained   bool
	}{
		{sectionString: "2-4,6-8", isContained: false},
		{sectionString: "2-3,4-5", isContained: false},
		{sectionString: "5-7,7-9", isContained: false},
		{sectionString: "2-8,3-7", isContained: true},
		{sectionString: "6-6,4-6", isContained: true},
		{sectionString: "2-6,4-8", isContained: false},
	}

	for _, i := range testData {
		t.Run(i.sectionString, func(t *testing.T) {
			p := strings.Split(i.sectionString, ",")
			a := NewSection(p[0])
			b := NewSection(p[1])
			got := a.contains(*b)
			want := i.isContained
			if got != want {
				t.Errorf("wrong containment value: got %v, want %v", got, want)
			}
		})
	}
}

func TestOverlap(t *testing.T) {
	testData := []struct {
		sectionString string
		isContained   bool
	}{
		{sectionString: "2-4,6-8", isContained: false},
		{sectionString: "2-3,4-5", isContained: false},
		{sectionString: "5-7,7-9", isContained: true},
		{sectionString: "2-8,3-7", isContained: true},
		{sectionString: "6-6,4-6", isContained: true},
		{sectionString: "2-6,4-8", isContained: true},
	}

	for _, i := range testData {
		t.Run(i.sectionString, func(t *testing.T) {
			p := strings.Split(i.sectionString, ",")
			a := NewSection(p[0])
			b := NewSection(p[1])
			got := a.overlap(*b)
			want := i.isContained
			if got != want {
				t.Errorf("wrong containment value: got %v, want %v", got, want)
			}
		})
	}
}

func TestGetNumberOfContainedSections(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	got := getNumberOfContainedSections(common.ReadLinesString(input))
	want := 2
	if got != want {
		t.Errorf("wrong number of contained sections: got %d, want %d", got, want)
	}
}

func TestGetNumberOfOverlappingSections(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	got := getNumberOfOverlappingSections(common.ReadLinesString(input))
	want := 4
	if got != want {
		t.Errorf("wrong number of overlapping sections: got %d, want %d", got, want)
	}
}
