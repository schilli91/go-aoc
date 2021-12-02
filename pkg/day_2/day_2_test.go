package day_2

import (
	"log"
	"os"
	"testing"
)

func TestDive(t *testing.T) {
	route, err := os.Open("test_data")
	if err != nil {
		log.Fatal(err)
	}
	defer route.Close()

	t.Run("amount of increasing depth, three measurements full workflow", func(t *testing.T) {
		got := Dive(route)
		want := 150
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d", got, want)
		}
	})
}

func TestAdvancedDive(t *testing.T) {
	route, err := os.Open("test_data")
	if err != nil {
		log.Fatal(err)
	}
	defer route.Close()

	t.Run("amount of increasing depth, three measurements full workflow", func(t *testing.T) {
		got := AdvancedDive(route)
		want := 900
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d", got, want)
		}
	})
}
