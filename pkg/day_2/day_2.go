package day_2

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/schilli91/aoc2021/pkg/common"
)

func Run() {
	fmt.Printf("Day 2\n")
	route, err := os.Open("pkg/day_2/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer route.Close()

	positionProduct := Dive(route)
	fmt.Printf("The product of the final position is %d\n", positionProduct)

	// Reset file cursor
	route.Seek(0, io.SeekStart)
	advancedPositionProduct := AdvancedDive(route)
	fmt.Printf("The product of the final position is %d\n", advancedPositionProduct)
}

func Dive(route *os.File) int {
	x, y := 0, 0
	maneuverStrings := common.ReadLinesString(route)

	for _, m := range maneuverStrings {
		maneuver, steps := getManeuver(m)

		switch maneuver {
		case "forward":
			x += steps
		case "up":
			y -= steps
		case "down":
			y += steps
		}
	}

	return x * y
}

func getManeuver(maneuverStr string) (string, int) {
	m := strings.Split(maneuverStr, " ")
	s, err := strconv.Atoi(m[1])
	if err != nil {
		log.Fatal(err)
	}
	return m[0], s
}

func AdvancedDive(route *os.File) int {
	x, y, aim := 0, 0, 0
	maneuverStrings := common.ReadLinesString(route)

	for _, m := range maneuverStrings {
		maneuver, steps := getManeuver(m)

		switch maneuver {
		case "forward":
			x += steps
			y += aim * steps
		case "up":
			aim -= steps
		case "down":
			aim += steps
		}
	}

	return x * y
}
