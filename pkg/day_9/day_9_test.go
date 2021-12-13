package day_9

import (
	"log"
	"os"
	"testing"
)

func TestDay9(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		m := LocalMinima(data)
		got := RiskLevel(m)
		want := 15
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test", func(t *testing.T) {
		data, err := os.Open("puzzle_input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		m := LocalMinima(data)
		got := RiskLevel(m)
		want := 425
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
