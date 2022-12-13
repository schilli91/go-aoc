package day_5

import (
	"reflect"
	"testing"

	"github.com/schilli91/go-aoc/pkg/common"
)

func TestGetStackNumbers(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	lines := common.ReadLinesString(input)
	got := getStackNumbers(lines)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong number of stacks: got %v, want %v", got, want)
	}
}

func TestGetSupplyPlan(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	lines := common.ReadLinesString(input)
	got := getSupplyLines(lines)
	want := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong wrong supply plan: got %v, want %v", got, want)
	}
}

func TestGetInstructionLines(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	lines := common.ReadLinesString(input)
	got := getInstructionLines(lines)
	want := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong wrong supply plan: got %v, want %v", got, want)
	}
}

func TestGetSupplyStacks(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	lines := common.ReadLinesString(input)
	//s := getSupplyPlan(lines)
	got := getSupplyStacks(lines)

	a := NewStack()
	a.push(&Crate{value: 'Z'})
	a.push(&Crate{value: 'N'})

	b := NewStack()
	b.push(&Crate{value: 'M'})
	b.push(&Crate{value: 'C'})
	b.push(&Crate{value: 'D'})

	c := NewStack()
	c.push(&Crate{value: 'P'})

	want := []Stack{*a, *b, *c}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong supply plan: got %v, want %v", got, want)
	}
}

func TestNewInstruction(t *testing.T) {
	given := "move 1 from 2 to 1"
	got := NewInstruction(given)
	want := &Instruction{amount: 1, from: 2, to: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong instruction: got %v, want %v", got, want)
	}
}

func TestGetInstructions(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	given := common.ReadLinesString(input)
	got := getInstructions(given)
	want := []Instruction{
		{amount: 1, from: 2, to: 1},
		{amount: 3, from: 1, to: 3},
		{amount: 2, from: 2, to: 1},
		{amount: 1, from: 1, to: 2},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong instruction: got %v, want %v", got, want)
	}
}

func TestExecInstruction(t *testing.T) {
	a := NewStack()
	a.push(&Crate{value: 'Z'})
	a.push(&Crate{value: 'N'})

	b := NewStack()
	b.push(&Crate{value: 'M'})
	b.push(&Crate{value: 'C'})
	b.push(&Crate{value: 'D'})

	c := NewStack()
	c.push(&Crate{value: 'P'})

	given := NewSupply([]Stack{*a, *b, *c})
	// fmt.Println(given)

	given.exec(Instruction{amount: 1, from: 2, to: 1})
	got := given

	d := NewStack()
	d.push(&Crate{value: 'Z'})
	d.push(&Crate{value: 'N'})
	d.push(&Crate{value: 'D'})

	e := NewStack()
	e.push(&Crate{value: 'M'})
	e.push(&Crate{value: 'C'})

	f := NewStack()
	f.push(&Crate{value: 'P'})

	want := NewSupply([]Stack{*d, *e, *f})

	// fmt.Println(got)
	// fmt.Println(want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong instruction execution: got %v, want %v", got, want)
	}
}

func TestPerformOperations(t *testing.T) {
	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	got := performOperations(common.ReadLinesString(input))

	a := NewStack()
	a.push(&Crate{value: 'C'})

	b := NewStack()
	b.push(&Crate{value: 'M'})

	c := NewStack()
	c.push(&Crate{value: 'P'})
	c.push(&Crate{value: 'D'})
	c.push(&Crate{value: 'N'})
	c.push(&Crate{value: 'Z'})

	want := NewSupply([]Stack{*a, *b, *c})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong supply: \ngot \t%v, \nwant \t%v", got, want)
	}
}

func TestGetTopCrates(t *testing.T) {
	a := NewStack()
	a.push(&Crate{value: 'C'})

	b := NewStack()
	b.push(&Crate{value: 'M'})

	c := NewStack()
	c.push(&Crate{value: 'P'})
	c.push(&Crate{value: 'D'})
	c.push(&Crate{value: 'N'})
	c.push(&Crate{value: 'Z'})

	given := NewSupply([]Stack{*a, *b, *c})
	got := given.getTopCrates()
	want := "CMZ"
	if got != want {
		t.Errorf("wrong top crates: got %s, want %s", got, want)
	}
}
