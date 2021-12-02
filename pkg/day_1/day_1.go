package day_1

import (
	"fmt"
	"log"
	"os"

	"github.com/schilli91/aoc2021/pkg/common"
)

func Run() {
	fmt.Println("Day 1")

	measurements, err := os.Open("pkg/day_1/puzzle_input")
	if err != nil {
		log.Fatal(err)
	}
	defer measurements.Close()

	measurement_lines := common.ReadLinesInt(measurements)

	c := CountDepthIncreases(measurement_lines)
	fmt.Printf("Depth increased %d times.\n", c)

	t := CountThreeWindowDepthIncreases(measurement_lines)
	fmt.Printf("Depth increased %d times using three measurement sums.\n", t)
}

func CountDepthIncreases(measurements []int) int {
	count := 0
	first := true
	var prev int
	for _, m := range measurements {
		if first {
			first = false
			prev = m
			continue
		}

		if prev < m {
			count++
		}
		prev = m
	}

	return count
}

func CountThreeWindowDepthIncreases(measurements []int) int {
	threeMeasurements := ThreeMeasurementSums(measurements)
	return CountDepthIncreases(threeMeasurements)
}

func ThreeMeasurementSums(measurements []int) []int {
	var three_measurements []int
	var a, b, c []int

	for _, m := range measurements {
		if len(a) == len(b) && len(b) == len(c) {
			a = append(a, m)
			continue
		}
		if len(b) == len(c) && len(a) == 1 {
			a = append(a, m)
			b = append(b, m)
			continue
		}

		a = append(a, m)
		b = append(b, m)
		c = append(c, m)

		if len(a) == 3 {
			three_measurements = append(three_measurements, sumElements(a))
			a = a[:0]
		}
		if len(b) == 3 {
			three_measurements = append(three_measurements, sumElements(b))
			b = b[:0]
		}
		if len(c) == 3 {
			three_measurements = append(three_measurements, sumElements(c))
			c = c[:0]
		}
	}

	return three_measurements
}

func sumElements(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return total
}
