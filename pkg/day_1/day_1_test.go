package day_1

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/schilli91/aoc2021/pkg/common"
)

func TestCountDepthIncreases(t *testing.T) {
	t.Run("amount of increasing depth", func(t *testing.T) {
		input := openInputFile("test_data")
		defer input.Close()

		count := CountDepthIncreases(common.ReadLinesInt(input))
		want := 7
		if want != count {
			t.Errorf("Miscounted increases of depth: got %d, want %d", count, want)
		}
	})
	t.Run("amount of increasing depth, real data", func(t *testing.T) {
		input := openInputFile("puzzle_input")
		defer input.Close()

		count := CountDepthIncreases(common.ReadLinesInt(input))
		want := 1121
		if want != count {
			t.Errorf("Miscounted increases of depth: got %d, want %d", count, want)
		}
	})
	t.Run("amount of increasing depth, three measurements", func(t *testing.T) {
		input := openInputFile("three_measurements")
		defer input.Close()

		count := CountDepthIncreases(common.ReadLinesInt(input))
		want := 5
		if want != count {
			t.Errorf("Miscounted increases of depth: got %d, want %d", count, want)
		}
	})
}

func TestCountThreeWindowDepthIncreases(t *testing.T) {
	t.Run("amount of increasing depth, three measurements full workflow", func(t *testing.T) {
		input := openInputFile("test_data")
		defer input.Close()

		count := CountThreeWindowDepthIncreases(common.ReadLinesInt(input))
		want := 5
		if want != count {
			t.Errorf("Miscounted increases of depth: got %d, want %d", count, want)
		}
	})
}

func TestThreeMeasurementSums(t *testing.T) {
	t.Run("amount of increasing depth", func(t *testing.T) {
		input := openInputFile("test_data")
		defer input.Close()

		got := ThreeMeasurementSums(common.ReadLinesInt(input))
		want := common.ReadLinesInt(openInputFile("three_measurements"))
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wrong three measurement sums: got %v, want %v", got, want)
		}
	})
}

func openInputFile(path string) *os.File {
	input, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return input
}
