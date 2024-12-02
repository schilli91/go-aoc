package day01

import (
	"testing"

	"github.com/schilli91/go-aoc/pkg/2024/datafile"
)

func TestGetDistanceLists(t *testing.T) {
	lines := datafile.ReadLinesString("test_data")
	expectedL := []int{3, 4, 2, 1, 3, 3}
	expectedR := []int{4, 3, 5, 3, 9, 3}

	l, r := getDistanceLists(lines)
	if !isEqual(l, expectedL) {
		t.Errorf("l != expectedL: %v != %v\n", l, expectedL)
	}
	if !isEqual(r, expectedR) {
		t.Errorf("r != expectedR: %v != %v\n", r, expectedR)
	}
}

func isEqual(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetDistance(t *testing.T) {
	testData := []struct {
		name  string
		given []string
		want  int
	}{
		{
			name:  "simple",
			given: []string{"1   1", "2   2"},
			want:  0,
		},
		{
			name:  "test_input",
			given: datafile.ReadLinesString("test_data"),
			want:  11,
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			got := calcDistance(tt.given)

			if tt.want != got {
				t.Errorf("want != got: %d != %d", tt.want, got)
			}
		})
	}
}
