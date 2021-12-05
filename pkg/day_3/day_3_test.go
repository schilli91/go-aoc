package day_3

import (
	"log"
	"os"
	"testing"

	"github.com/schilli91/aoc2021/pkg/common"
)

func TestDay3(t *testing.T) {

	t.Run("power consumption, test data", func(t *testing.T) {
		testInput, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer testInput.Close()
		tr := TransposeReport(common.ReadLinesString(testInput))
		got := PowerConsumption(tr)
		want := 198
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})

	t.Run("power consumption, puzzle input", func(t *testing.T) {
		puzzleInput, err := os.Open("puzzle_input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer puzzleInput.Close()

		tr := TransposeReport(common.ReadLinesString(puzzleInput))
		got := PowerConsumption(tr)
		want := 2595824
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})

	t.Run("life support rating, test data", func(t *testing.T) {
		testInput, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer testInput.Close()
		got := LifeSupportRating(common.ReadLinesString(testInput))
		want := 230
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})

	t.Run("oxygen generator rating, test data", func(t *testing.T) {
		testInput, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer testInput.Close()
		//tr := TransposeReport(common.ReadLinesString(testInput))
		got := OxygenGeneratorRating(common.ReadLinesString(testInput))
		want := 23
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})

	t.Run("oxygen generator rating, test data", func(t *testing.T) {
		testInput, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer testInput.Close()
		got := Co2ScrubberRating(common.ReadLinesString(testInput))
		want := 10
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
