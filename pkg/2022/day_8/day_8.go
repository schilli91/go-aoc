package day_8

import (
	"fmt"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 8
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile(fmt.Sprintf("pkg/2022/day_%d/puzzle_input.txt", day))
	defer input.Close()
	f := NewForest(common.ReadLinesString(input))

	o := f.findNumVisibleTrees()
	fmt.Printf("The value of part one is %d\n", o)

	// 6852; your answer is too high

	t := -1
	fmt.Printf("The value of part two is %d\n", t)
}

type Forest struct {
	trees [][]int
}

func NewForest(grid []string) *Forest {
	f := new(Forest)
	f.trees = make([][]int, len(grid))
	for i, row := range grid {
		f.trees[i] = make([]int, 0)
		for _, tree := range row {
			h := int(tree - '0')
			f.trees[i] = append(f.trees[i], h)
		}
	}

	return f
}

func (f Forest) getSize() int {
	s := 0
	for _, t := range f.trees {
		s += len(t)
	}
	return s
}

func (f Forest) findNumHiddenTrees() int {
	//A tree is visible if all of the other trees between it and an edge of the grid are shorter than it.

	num := 0
	for i := 1; i < len(f.trees)-1; i++ {
		for j := 1; j < len(f.trees[i])-1; j++ {
			c := f.trees[i][j]
			if c > f.trees[i-1][j] {
				continue
			}
			if c > f.trees[i+1][j] {
				continue
			}
			if c > f.trees[i][j-1] {
				continue
			}
			if c > f.trees[i][j+1] {
				continue
			}
			//fmt.Printf("hidden (%d/%d): %d\n", j, i, c)
			num++
		}
	}
	return num
}

func (f Forest) findNumVisibleTrees() int {
	hidden := f.findNumHiddenTrees()
	return f.getSize() - hidden
}
