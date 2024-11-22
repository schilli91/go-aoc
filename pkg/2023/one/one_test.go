package one

import (
	"fmt"
	"testing"

	"github.com/schilli91/go-aoc/pkg/2023/datafile"
)

func TestDecodeLine(t *testing.T) {
	f := datafile.OpenInputFile("test_data")
	defer f.Close()
	lines := datafile.ReadLinesString(f)
	expected := []int{12, 38, 15, 77}

	for i, l := range lines {
		t.Run(fmt.Sprintf("Line: %s", l), func(t *testing.T) {
			got := decodeLine(l)
			want := expected[i]

			if got != want {
				t.Fatalf("\t\t\t\tgot: %d, want: %d", got, want)
			}
		})
	}
}
