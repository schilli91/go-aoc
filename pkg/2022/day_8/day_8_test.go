package day_8

import (
	"reflect"
	"testing"

	"github.com/schilli91/go-aoc/pkg/common"
)

func TestNewForest(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	given := common.ReadLinesString(input)

	got := NewForest(given)
	want := &Forest{
		trees: [][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong forest, \ngot \t%v \nwant \t%v", got, want)
	}
}

func TestGetSize(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()

	given := NewForest(common.ReadLinesString(input))
	got := given.getSize()
	want := 25
	if got != want {
		t.Errorf("wrong forest size, \ngot \t%d \nwant \t%d", got, want)
	}
}

func TestFindNumHiddenTrees(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()

	given := NewForest(common.ReadLinesString(input))
	got := given.findNumHiddenTrees()
	want := 4
	if got != want {
		t.Errorf("wrong forest size, \ngot \t%d \nwant \t%d", got, want)
	}
}

func TestFindNumVisibleTrees(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()

	given := NewForest(common.ReadLinesString(input))
	got := given.findNumVisibleTrees()
	want := 21
	if got != want {
		t.Errorf("wrong forest size, \ngot \t%d \nwant \t%d", got, want)
	}
}
