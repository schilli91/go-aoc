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
	numFish := LanternfishLife(s, days)

	fmt.Printf("After %d days, there are %d fish.\n", days, numFish)
}

func LanternfishLife(swarm []int, days int) int {
	numFish := len(swarm)
	for d := 0; d < days; d++ {
		babyFish := []int{}
		for i, f := range swarm {
			if f == 0 {
				babyFish = append(babyFish, 8)
				swarm[i] = 6
			} else {
				swarm[i] = f - 1
			}
		}
		if len(babyFish) > 0 {
			// TODO: Vlt als goroutine?
			numFish += LanternfishLife(babyFish, days-d-1)
		}
		// if d == 150 || d == 200 || d == 220 {
		// 	fmt.Printf("After %d days, there are %d fish.\n", d+1, numFish)
		// }
	}

	return numFish
}

func GoLanternfishLife(swarm []int, days int, ch chan int) {
	numFish := len(swarm)
	for d := 0; d < days; d++ {
		babyFish := []int{}
		for i, f := range swarm {
			if f == 0 {
				babyFish = append(babyFish, 8)
				swarm[i] = 6
			} else {
				swarm[i] = f - 1
			}
		}
		if len(babyFish) > 0 {
			// TODO: Vlt als goroutine?
			go LanternfishLife(babyFish, days-d-1)
			numBabies := <-ch
			numFish += numBabies
		}
		// if d == 150 || d == 200 || d == 220 {
		// 	fmt.Printf("After %d days, there are %d fish.\n", d+1, numFish)
		// }
	}
	ch <- numFish
	return
}

//1171199621
//1073741823.5
//2147483647

func OLDLanternfishLife(swarm Swarm, days int) int {
	for d := 0; d < days; d++ {
		babyFish := []int{}
		for i, f := range swarm.fish {
			if f == 0 {
				babyFish = append(babyFish, 8)
				swarm.fish[i] = 6
			} else {
				swarm.fish[i] = f - 1
			}
		}
		swarm.fish = append(swarm.fish, babyFish...)
		if d%50 == 0 {
			fmt.Printf("After %d days, there are %d fish.\n", d+1, len(swarm.fish))
		}
	}

	return len(swarm.fish)
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
