package day_10

import (
	"log"
	"os"
	"testing"
)

func TestDay10(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		got := 0
		want := 26397
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
