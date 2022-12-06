package day_1

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 1
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile("pkg/2022/day_1/puzzle_input.txt")
	defer input.Close()

	c := getCaloriesPerElf(*input)
	m := getMaxCalories(c)
	fmt.Printf("The elf carrying max calories carries %d.\n", m)

	maxThree := getThreeMaxCalories(c)
	sumOfMaxThree := sum(maxThree)
	fmt.Printf("The three elves carrying max calories carry %d in total.\n", sumOfMaxThree)
}

func getCaloriesPerElf(caloriesFile os.File) []int {
	caloriesLines := common.ReadLinesString(&caloriesFile)

	caloriesPerElf := make([]int, 1)
	for _, c := range caloriesLines {
		if c == "" {
			caloriesPerElf = append(caloriesPerElf, 0)
			continue
		}
		v, err := strconv.Atoi(c)
		if err != nil {
			log.Fatalln("Failed to parse calories.")
		}
		caloriesPerElf[len(caloriesPerElf)-1] += v
	}

	return caloriesPerElf
}

func getMaxCalories(calories []int) int {
	max := -1
	for _, i := range calories {
		if i > max {
			max = i
		}
	}
	return max
}

func getThreeMaxCalories(calories []int) []int {
	maxThree := make([]int, 3)
	indexOfLowest := 0

	for _, i := range calories {
		if i > maxThree[indexOfLowest] {
			maxThree[indexOfLowest] = i
			min := getMinCalories(maxThree)
			var err error
			indexOfLowest, err = indexOf(maxThree, min)
			if err != nil {
				log.Fatal("Failed to find value in slice.")
			}
			maxThree[indexOfLowest] = min
		}
	}

	return maxThree
}

func getMinCalories(calories []int) int {
	min := calories[0]
	for _, i := range calories {
		if i < min {
			min = i
		}
	}

	return min
}

func indexOf(haystack []int, needle int) (int, error) {
	for k, v := range haystack {
		if v == needle {
			return k, nil
		}
	}
	return -1, fmt.Errorf("value %d not found in slice %v", needle, haystack)
}

func sum(values []int) int {
	sum := 0
	for _, i := range values {
		sum += i
	}
	return sum
}
