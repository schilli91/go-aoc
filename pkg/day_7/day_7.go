package day_7

import (
	"fmt"
	"log"
	"os"

	"github.com/schilli91/aoc2021/pkg/common"
)

func Run() {
	day := 7
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	p := LoadCrabs(data)
	fuel := CheapestPosition(p, false)
	fuelDynamic := CheapestPosition(p, true)

	fmt.Printf("The crabs need %d fuel for the cheapest position.\n", fuel)
	fmt.Printf("The crabs need %d fuel for the cheapest position with dynamic fuel cost.\n", fuelDynamic)
}

func CheapestPosition(positions []int, dynamicCosts bool) int {
	min, max := FindLimits(positions)
	cheapest := Fuel(positions, min, dynamicCosts)

	for i := min; i <= max; i++ {
		c := Fuel(positions, i, dynamicCosts)
		if c < cheapest {
			cheapest = c
		}
	}

	return cheapest
}

func Fuel(positions []int, targetPosition int, dynamicCosts bool) int {
	total := 0
	for _, p := range positions {
		if dynamicCosts {
			total += DynamicCosts(abs(p - targetPosition))
		} else {
			total += abs(p - targetPosition)
		}
	}
	return total
}

func LoadCrabs(data *os.File) []int {
	return common.ReadCommaSeparatedLineInt(data)
}

func FindLimits(positions []int) (int, int) {
	return min(positions), max(positions)
}

func max(v []int) int {
	max := 0
	for _, i := range v {
		if i > max {
			max = i
		}
	}
	return max
}

func min(v []int) int {
	min := v[0]
	for _, i := range v {
		if i < min {
			min = i
		}
	}
	return min
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func DynamicCosts(i int) int {
	n := abs(i)
	return n * (n + 1) / 2
}
