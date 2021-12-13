package day_9

import (
	"fmt"
	"log"
	"os"

	"github.com/schilli91/aoc2021/pkg/common"
)

type Point struct {
	x, y, value int
}

func Run() {
	day := 9
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	m := LocalMinima(data)
	riskLevel := RiskLevel(m)
	fmt.Printf("The risk level is %d.\n", riskLevel)
}

func Basin(minima []Point) {

}

func LocalMinima(f *os.File) []Point {
	heightMap := HeightMap(f)

	minima := []Point{}
	for i := range heightMap {
		for j, d := range heightMap[i] {
			up, down, left, right := d+1, d+1, d+1, d+1
			if i-1 >= 0 {
				up = heightMap[i-1][j]
			}
			if i+1 < len(heightMap) {
				down = heightMap[i+1][j]
			}
			if j-1 >= 0 {
				left = heightMap[i][j-1]
			}
			if j+1 < len(heightMap[0]) {
				right = heightMap[i][j+1]
			}

			if d < up && d < down && d < left && d < right {
				minima = append(minima, Point{x: j, y: i, value: d})
			}
		}
	}

	return minima
}

func HeightMap(f *os.File) [][]int {
	lines := common.ReadLinesString(f)
	heightMap := make([][]int, len(lines))
	for i := range heightMap {
		heightMap[i] = make([]int, len(lines[0]))
	}

	for i, l := range lines {
		for j, d := range l {
			heightMap[i][j] = int(d - '0')
		}
	}

	return heightMap
}

func RiskLevel(minima []Point) int {
	total := 0
	for _, m := range minima {
		total += 1 + m.value
	}
	return total
}
