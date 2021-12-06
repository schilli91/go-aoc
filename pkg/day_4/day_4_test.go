package day_4

import (
	"log"
	"os"
	"testing"
)

func TestDay4(t *testing.T) {
	t.Run("test play bingo, first winning card", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		values, cards := LoadBingoCards(data)

		got := PlayBingo(values, cards, false)
		want := 4512
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test play bingo, last winning card", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		values, cards := LoadBingoCards(data)

		got := PlayBingo(values, cards, true)
		want := 1924
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test bingo card", func(t *testing.T) {
		r := [5]string{
			"22 13 17 11  0",
			"8  2 23  4 24",
			"21  9 14 16  7",
			"6 10  3 18  5",
			"1 12 20 15 19",
		}
		b := NewBingoCard(r)
		got := b.rows

		w := [5][5]string{
			{"22", "13", "17", "11", "0"},
			{"8", "2", "23", "4", "24"},
			{"21", "9", "14", "16", "7"},
			{"6", "10", "3", "18", "5"},
			{"1", "12", "20", "15", "19"},
		}
		want := w
		if want != got {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
	t.Run("test bingo, column", func(t *testing.T) {
		b := new(BingoCard)
		// b.checks=[5][5]bool{
		// 	{false, false, false, false, false},
		// 	{false, false, false, false, false},
		// 	{false, false, false, false, false},
		// 	{false, false, false, false, false},
		// 	{false, false, false, false, false},
		// }
		b.checks = [5][5]bool{
			{false, true, false, false, false},
			{false, true, false, false, false},
			{false, true, false, false, false},
			{false, true, false, false, false},
			{false, true, false, false, false},
		}
		got := b.Bingo()
		want := true
		if want != got {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
	t.Run("test bingo, row", func(t *testing.T) {
		b := new(BingoCard)
		b.checks = [5][5]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{true, true, true, true, true},
			{false, false, false, false, false},
		}
		got := b.Bingo()
		want := true
		if want != got {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
	t.Run("test bingo, no bingo", func(t *testing.T) {
		b := new(BingoCard)
		b.checks = [5][5]bool{
			{false, false, true, false, true},
			{true, true, false, false, false},
			{false, false, true, false, true},
			{true, true, false, true, false},
			{true, true, true, false, false},
		}
		got := b.Bingo()
		want := false
		if want != got {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
	t.Run("test bingo, no bingo", func(t *testing.T) {
		b := new(BingoCard)
		b.checks = [5][5]bool{
			{false, false, true, true, false},
			{false, true, true, true, true},
			{true, true, true, true, true},
			{false, true, false, false, true},
			{false, false, false, false, false},
		}
		got := b.Bingo()
		want := true
		if want != got {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
	t.Run("test bingo, no bingo empty card", func(t *testing.T) {
		b := new(BingoCard)
		b.checks = [5][5]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}
		got := b.Bingo()
		want := false
		if want != got {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
}
