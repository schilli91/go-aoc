package day_5

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/schilli91/aoc2021/pkg/common"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func (l Line) LinePoints() []Point {
	start, end := l.start, l.end
	if start.x == end.x && start.y == end.y {
		return []Point{start}
	}

	if l.start.x > l.end.x || l.start.y > l.end.y {
		start, end = end, start
	}

	return []Point{}
}

type VentsMap struct {
	vents []Point
}

func Run() {
	day := 5
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	fmt.Printf("Day %d\n", day)
}

func HydrothermalVents(data *os.File) int {

	return 0
}

func LoadLines(data *os.File) []Line {
	rows := common.ReadLinesString(data)

	lines := []Line{}
	for _, row := range rows {
		lines = append(lines, NewLineFromStr(row))
	}

	return lines
}

func NewLineFromStr(s string) Line {
	parts := strings.Split(strings.TrimSpace(s), "->")

	start := NewPointFromStr(strings.TrimSpace(parts[0]))
	end := NewPointFromStr(strings.TrimSpace(parts[1]))

	return Line{start: start, end: end}
}

func NewPointFromStr(s string) Point {
	p := strings.Split(strings.TrimSpace(s), ",")
	x, err := strconv.Atoi(p[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(p[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{x: x, y: y}
}
