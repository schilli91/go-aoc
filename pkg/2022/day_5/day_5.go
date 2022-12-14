package day_5

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 5
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile(fmt.Sprintf("pkg/2022/day_%d/puzzle_input.txt", day))
	defer input.Close()
	lines := common.ReadLinesString(input)
	s := performOperations(lines)
	c := s.getTopCrates()
	fmt.Printf("The top crates are: %s\n", c)

	s9001 := performOperations9001(lines)
	o := s9001.getTopCrates()
	fmt.Printf("The top crates with 9001 are: %s\n", o)
}

type Supply struct {
	stacks map[int]Stack
}

type Stack struct {
	crates []Crate
}

func NewStack() *Stack {
	s := new(Stack)
	s.crates = make([]Crate, 0)
	return s
}

func (s *Stack) push(c *Crate) {
	s.crates = append(s.crates, *c)
}

func (s *Stack) pop() (Crate, error) {
	if len(s.crates) == 0 {
		return Crate{}, errors.New("stack is empty")
	}
	c := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return c, nil
}

type Crate struct {
	value rune
}

type Instruction struct {
	amount int
	from   int
	to     int
}

func (s *Supply) exec(instruction Instruction) {
	fmt.Println(instruction)
	for i := 0; i < instruction.amount; i++ {
		f := s.stacks[instruction.from]
		c, err := f.pop()
		if err != nil {
			log.Fatalf("could not pop element from stack: %v", f)
		}
		t := s.stacks[instruction.to]
		t.push(&c)

		s.stacks[instruction.from] = f
		s.stacks[instruction.to] = t
	}
}

func (s *Supply) exec9001(instruction Instruction) {
	f := s.stacks[instruction.from]
	t := s.stacks[instruction.to]
	transfer := f.crates[len(f.crates)-instruction.amount:]
	for _, c := range transfer {
		_, err := f.pop()
		if err != nil {
			log.Fatalf("could not pop element from stack: %v", f)
		}
		t.push(&c)
	}
	s.stacks[instruction.from] = f
	s.stacks[instruction.to] = t
}

func (s Supply) getTopCrates() string {
	r := make([]rune, 0)
	// Iterating over map with `range` does not retain order.
	for i := 1; i <= len(s.stacks); i++ {
		stack := s.stacks[i]
		r = append(r, stack.crates[len(stack.crates)-1].value)
	}

	return string(r)
}

func NewInstruction(l string) *Instruction {
	i := new(Instruction)

	parts := strings.Split(l, "from")
	move := strings.ReplaceAll(parts[0], "move", "")
	from := strings.Split(parts[1], "to")[0]
	to := strings.Split(l, "to")[1]

	move = strings.Trim(move, " ")
	from = strings.Trim(from, " ")
	to = strings.Trim(to, " ")

	var err error
	i.amount, err = strconv.Atoi(move)
	if err != nil {
		log.Fatalf("could not convert string '%v' to int", move)
	}
	i.from, err = strconv.Atoi(from)
	if err != nil {
		log.Fatalf("could not convert string '%v' to int", from)
	}
	i.to, err = strconv.Atoi(to)
	if err != nil {
		log.Fatalf("could not convert string '%v' to int", to)
	}
	return i
}

func NewCrate(planEntry string) *Crate {
	p := strings.Trim(planEntry, " ")
	if p == "" || len(p) != 3 {
		return &Crate{}
	}
	return &Crate{rune(p[1])}
}

func getStackNumbers(inputLines []string) []int {
	_, line := getStackNumberLine(inputLines)
	lineElements := strings.Split(line, " ")

	stackNumbers := make([]int, 0)
	for _, number := range lineElements {
		if number == "" {
			continue
		}
		x, err := strconv.Atoi(number)
		if err != nil {
			log.Fatalf("could not convert string '%s' to int", number)
		}
		stackNumbers = append(stackNumbers, x)
	}
	return stackNumbers
}

func getStackNumberLine(inputLines []string) (int, string) {
	i := -1
	for key, l := range inputLines {
		if strings.TrimSpace(l) == "" {
			i = key - 1
			break
		}
	}

	return i, inputLines[i]
}

func getSupplyLines(inputLines []string) []string {
	i, _ := getStackNumberLine(inputLines)
	return inputLines[:i]
}

func getInstructionLines(inputLines []string) []string {
	i, _ := getStackNumberLine(inputLines)
	return inputLines[i+2:]
}

func NewSupply(stacks []Stack) *Supply {
	supply := new(Supply)
	supply.stacks = make(map[int]Stack)
	for i, s := range stacks {
		supply.stacks[i+1] = s
	}
	return supply
}

func getSupplyStacks(inputLines []string) []Stack {
	stackNumbers := getStackNumbers(inputLines)
	supplyPlan := getSupplyLines(inputLines)

	stacks := make([]Stack, stackNumbers[len(stackNumbers)-1])
	//for _, l := range supplyPlan {
	for j := len(supplyPlan) - 1; j >= 0; j-- {
		l := supplyPlan[j]
		for i := 0; i < len(l); i += 4 {
			var item string
			if i+4 >= len(l) {
				item = l[i:]
			} else {
				item = l[i : i+4]
			}
			if strings.Trim(item, " ") == "" {
				continue
			}
			stackPos := i/4 + 1
			stacks[stackPos-1].push(NewCrate(item))
		}
	}
	return stacks
}

func getInstructions(inputLines []string) []Instruction {
	lines := getInstructionLines(inputLines)
	instructions := make([]Instruction, 0)
	for _, l := range lines {
		instructions = append(instructions, *NewInstruction(l))
	}
	return instructions
}

func performOperations(inputLines []string) *Supply {
	supply := NewSupply(getSupplyStacks(inputLines))
	instructions := getInstructions(inputLines)
	for _, i := range instructions {
		supply.exec(i)
	}
	fmt.Println(len(instructions))
	return supply
}

func performOperations9001(inputLines []string) *Supply {
	supply := NewSupply(getSupplyStacks(inputLines))
	instructions := getInstructions(inputLines)
	for _, i := range instructions {
		supply.exec9001(i)
	}
	fmt.Println(len(instructions))
	return supply
}
