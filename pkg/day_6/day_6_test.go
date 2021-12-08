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
		got := LanternfishLife(s, 80)
		want := 5934
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test lanternfish life, goroutine", func(t *testing.T) {
		t.Skip("skip")
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		s := LoadSwarm(data)
		ch := make(chan int)
		go GoLanternfishLife(s, 80, ch)
		got := <-ch
		close(ch)
		want := 5934
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test lanternfish life, infinite life", func(t *testing.T) {
		//t.Skip("skip")
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		s := LoadSwarm(data)
		got := LanternfishLife(s, 200)
		want := 200
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test lanternfish life, infinite life", func(t *testing.T) {
		//t.Skip("skip")
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		s := LoadSwarm(data)
		got := LanternfishLife(s, 256)
		want := 26984457539
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
}
