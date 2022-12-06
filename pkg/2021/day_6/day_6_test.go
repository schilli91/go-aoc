package day_6

import (
	"log"
	"os"
	"testing"
)

func TestDay6(t *testing.T) {
	t.Run("test lanternfish life", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		s := LoadSwarm(data)
		got := LanternFishLife(s, 80)
		want := int64(5934)
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test lanternfish life, infinite life", func(t *testing.T) {
		//t.Skip("skip\n")
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		s := LoadSwarm(data)
		got := LanternFishLife(s, 200)
		want := int64(200)
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test lanternfish life, infinite life", func(t *testing.T) {
		//t.Skip("skip\n")
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		s := LoadSwarm(data)
		got := LanternFishLife(s, 256)
		want := int64(26984457539)
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
