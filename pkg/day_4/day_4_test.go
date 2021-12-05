package day_4

import (
	"log"
	"os"
	"testing"
)

func TestDay4(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		got := PlayBingo(data)
		want := 4512
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
