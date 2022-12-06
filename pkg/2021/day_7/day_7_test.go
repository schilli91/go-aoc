package day_7

import (
	"log"
	"os"
	"testing"
)

func TestDay7(t *testing.T) {
	t.Run("test cheapest position", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		p := LoadCrabs(data)
		got := CheapestPosition(p, false)
		want := 37
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test cheapest position, dynamic costs", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		p := LoadCrabs(data)
		got := CheapestPosition(p, true)
		want := 168
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test limits", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		p := LoadCrabs(data)
		gotMin, gotMax := FindLimits(p)
		wantMin, wantMax := 0, 16
		if wantMin != gotMin || wantMax != gotMax {
			t.Errorf("Incorrect value: got [%d, %d], want [%d, %d].\n\n", gotMin, gotMax, wantMin, wantMax)
		}
	})
}
