package day_5

import (
	"log"
	"os"
	"testing"
)

func TestDayX(t *testing.T) {
	t.Run("test hydro thermal vents", func(t *testing.T) {
		data, err := os.Open("test_data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()

		got := HydrothermalVents(data)
		want := 5
		if want != got {
			t.Errorf("Incorrect value: got %d, want %d\n\n", got, want)
		}
	})
	t.Run("test line from string", func(t *testing.T) {
		s := "0,9 -> 5,9"

		got := NewLineFromStr(s)

		want := Line{
			start: Point{x: 0, y: 9},
			end:   Point{x: 5, y: 9},
		}
		if want != got {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
	t.Run("test points from line", func(t *testing.T) {
		s := "3,4 -> 1,4"
		l := NewLineFromStr(s)
		got := l.LinePoints()

		want := []Point{
			{x: 1, y: 4},
			{x: 2, y: 4},
			{x: 3, y: 4},
		}

		// fmt.Println(got)
		// fmt.Println(want)

		if !comparePointSlices(want, got) {
			t.Errorf("Incorrect value: got %v, want %v\n\n", got, want)
		}
	})
}

func comparePointSlices(a []Point, b []Point) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		ai, bi := a[i], b[i]
		if ai.x != bi.x {
			return false
		}
		if ai.y != bi.y {
			return false
		}
	}
	return true
}
