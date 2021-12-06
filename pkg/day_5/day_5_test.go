package day_5

import (
	"log"
	"os"
	"testing"
)

func TestDayX(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		got := 0
		want := 150
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
