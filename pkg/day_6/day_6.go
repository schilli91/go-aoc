package day_6

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/schilli91/aoc2021/pkg/common"
)

type Swarm struct {
	fish []int
}

func Run() {
	day := 6
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	s := LoadSwarm(data)
	days := 80
	numFish := LanternFishLife(s, days)

	fmt.Printf("After %d days, there are %d fish.\n", days, numFish)
	moreDays := 256
	numImmortalFish := LanternFishLife(s, moreDays)

	fmt.Printf("After %d days, there are %d fish.\n", moreDays, numImmortalFish)
}

func LanternFishLife(swarm []int, days int) int64 {
	fish := LoadFish(swarm)
	for i := 0; i < days; i++ {
		b := fish[0]
		fish = fish[1:]
		fish[6] += b
		fish = append(fish, b)
	}
	return CountFish(fish)
}

func LoadFish(swarm []int) []int64 {
	s := make([]int64, 9)
	for _, f := range swarm {
		s[f] += 1
	}
	return s
}

func CountFish(fish []int64) int64 {
	total := int64(0)
	for _, f := range fish {
		total += f
	}
	return total
}

func LoadSwarm(data *os.File) []int {
	s := common.ReadLinesString(data)
	fishStr := strings.Split(s[0], ",")
	swarm := []int{}

	for _, f := range fishStr {
		fish, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		swarm = append(swarm, fish)
	}
	return swarm
}
