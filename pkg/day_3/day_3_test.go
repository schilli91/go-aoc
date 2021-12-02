package day_3

import (
	"log"
	"os"
	"testing"
)

func TestDay3(t *testing.T) {
	route, err := os.Open("test_data")
	if err != nil {
		log.Fatal(err)
	}
	defer route.Close()

	t.Run("amount of increasing depth, three measurements full workflow", func(t *testing.T) {
		got := 0
		want := 150
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d", got, want)
		}
	})
}
