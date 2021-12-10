package day_8

import (
	"log"
	"os"
	"testing"

	"github.com/schilli91/aoc2021/pkg/common"
)

func TestDay8(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		lines := common.ReadLinesString(data)

		d := LoadDisplays(lines)
		got := CountUniqueDigits(d)
		want := 26
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test decode display", func(t *testing.T) {
		t.Skip("skip")
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		lines := []string{
			"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
		}

		d := LoadDisplays(lines)
		got := d[0].Decode()
		want := 5353
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test decode display", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		lines := []string{
			"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		}

		d := LoadDisplays(lines)
		got := d[0].Decode()
		want := 8394
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test decode displays", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		lines := common.ReadLinesString(data)

		d := LoadDisplays(lines)
		got := DecodeDisplays(d)
		want := 61229
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
