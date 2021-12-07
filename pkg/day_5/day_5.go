package day_5

import (
	"errors"
	"fmt"
	"io"
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

func (l Line) LinePoints(skipDiag bool) ([]Point, error) {
	start, end := l.start, l.end
	if start.x == end.x && start.y == end.y {
		return []Point{start}, nil
	}
	isDiag := false
	if start.x != end.x && start.y != end.y {
		if skipDiag {
			return nil, errors.New("diagonal line")
		}
		isDiag = true
	}

	if start.x > end.x || start.y > end.y {
		start, end = end, start
	}

	points := []Point{}
	points = append(points, start)

	dx, dy := 0, 0
	if start.x == end.x {
		dy = 1
	}
	if start.y == end.y {
		dx = 1
	}

	if isDiag {
		if start.x < end.x && start.y < end.y {
			dx, dy = 1, 1
		}
		if start.x < end.x && start.y > end.y {
			dx, dy = 1, -1
		}
		if start.x > end.x && start.y > end.y {
			dx, dy = -1, -1
		}
		if start.x > end.x && start.y < end.y {
			dx, dy = -1, 1
		}
	}

	p := Point{x: start.x, y: start.y}
	for {
		p = Point{x: p.x + dx, y: p.y + dy}
		points = append(points, p)
		if p.x == end.x && p.y == end.y {
			break
		}
	}

	return points, nil
}

type VentsMap struct {
	vents [][]int
}

func NewVentsMap(size int) VentsMap {
	v := make([][]int, size)
	for i := range v {
		v[i] = make([]int, size)
	}
	return VentsMap{vents: v}
}

func (v *VentsMap) addPoints(points []Point) {
	for _, p := range points {
		v.vents[p.x][p.y] += 1
	}
}

func (v VentsMap) countDangerousPoints() int {
	total := 0
	for i := range v.vents {
		for j := range v.vents[i] {
			if v.vents[i][j] > 1 {
				total += 1
			}
		}
	}
	return total
}

func Run() {
	day := 5
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	dangerousPoints := HydrothermalVents(data, true)
	fmt.Printf("There are a total of %d dangerous points.\n", dangerousPoints)

	data.Seek(0, io.SeekStart)
	dangerousPointsWithDiag := HydrothermalVents(data, false)
	fmt.Printf("There are a total of %d dangerous points incl. diagonals.\n", dangerousPointsWithDiag)
}

func HydrothermalVents(data *os.File, skipDiag bool) int {
	lines := LoadLines(data)
	k := FindMaxDimension(data)
	fmt.Printf("k: %d", k)
	v := NewVentsMap(k)

	for _, l := range lines {
		p, err := l.LinePoints(skipDiag)
		if err != nil {
			fmt.Println(err)
			continue
		}
		v.addPoints(p)
	}

	return v.countDangerousPoints()
}

func LoadLines(data *os.File) []Line {
	rows := common.ReadLinesString(data)

	lines := []Line{}
	for _, row := range rows {
		lines = append(lines, NewLineFromStr(row))
	}

	return lines
}

func FindMaxDimension(data *os.File) int {
	data.Seek(0, io.SeekStart)
	rows := common.ReadLinesString(data)
	pointStrings := []string{}
	for _, r := range rows {
		parts := strings.Split(strings.TrimSpace(r), "->")
		pointStrings = append(pointStrings, strings.TrimSpace(parts[0]))
		pointStrings = append(pointStrings, strings.TrimSpace(parts[1]))
	}

	nums := []int{}
	for _, p := range pointStrings {
		parts := strings.Split(strings.TrimSpace(p), ",")
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, a, b)
	}

	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max + 1
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
