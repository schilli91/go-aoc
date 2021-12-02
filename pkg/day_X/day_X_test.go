package day_X

import (
	"log"
	"os"
	"testing"
)

func TestDayX(t *testing.T) {
	route, err := os.Open("test_data")
	if err != nil {
		log.Fatal(err)
	}
	defer route.Close()

	t.Run("test", func(t *testing.T) {
		got := 0
		want := 150
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d", got, want)
		}
	})
}
