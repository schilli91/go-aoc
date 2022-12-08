package day_3

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/schilli91/go-aoc/pkg/common"
)

func TestNewRucksack(t *testing.T) {
	r := NewRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	got := r.firstCompartment
	want := "vJrwpWtwJgWr"
	if got != want {
		t.Errorf("wrong content in first compartment: got %s, want %s", got, want)
	}
	got = r.secondCompartment
	want = "hcsFMMfFFhFp"
	if got != want {
		t.Errorf("wrong content in second compartment: got %s, want %s", got, want)
	}
}

func TestFindSharedItemType(t *testing.T) {
	r := NewRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	got := r.findSharedItemType()
	want := []rune{'p'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong shared item type: got %v, want %v", got, want)
	}
}

var testItems = []struct {
	item             []rune
	expectedPriority int
}{
	{item: []rune{'p'}, expectedPriority: 16},
	{item: []rune{'L'}, expectedPriority: 38},
	{item: []rune{'P'}, expectedPriority: 42},
	{item: []rune{'v'}, expectedPriority: 22},
	{item: []rune{'t'}, expectedPriority: 20},
	{item: []rune{'s'}, expectedPriority: 19},
}

func TestGetItemPriority(t *testing.T) {
	for _, ii := range testItems {
		testName := fmt.Sprintf("item=%v,expected=%d", ii.item, ii.expectedPriority)

		t.Run(testName, func(t *testing.T) {
			got := getItemPriority(ii.item)
			want := ii.expectedPriority
			if got != want {
				t.Errorf("wrong priority: got %d, want %d", got, want)
			}
		})

	}
}

func TestGetPriorityOfSharedItems(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	got := getPriorityOfSharedItems(common.ReadLinesString(input))
	want := 157
	if got != want {
		t.Errorf("wrong priority: got %d, want %d", got, want)
	}
}
