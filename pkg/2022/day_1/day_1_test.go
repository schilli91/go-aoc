package day_1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/schilli91/go-aoc/pkg/common"
)

func TestGetCaloriesPerElf(t *testing.T) {
	day := 1
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	got := getCaloriesPerElf(*input)
	want := []int{6000, 4000, 11000, 24000, 10000}
	if len(got) != len(want) {
		t.Errorf("Mismatch of number of elves: got %d, want %d", len(got), len(want))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Mismatch of calories: got %v, want %v", got, want)
	}
}

func TestGetMaxCalories(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	calories := getCaloriesPerElf(*input)

	got := getMaxCalories(calories)
	want := 24000
	if got != want {
		t.Errorf("Mismatch of max calories: got %v, want %v", got, want)
	}
}

func TestGetThreeMaxCalories(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	calories := getCaloriesPerElf(*input)

	maxThreeCalories := getThreeMaxCalories(calories)
	got := sum(maxThreeCalories)
	want := 45000
	if got != want {
		t.Errorf("Mismatch of max calories: got %v, want %v", got, want)
	}
}
