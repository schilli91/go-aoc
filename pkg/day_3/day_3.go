package day_3

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/schilli91/aoc2021/pkg/common"
)

func Run() {
	fmt.Printf("Day 3\n")
	puzzleInput, err := os.Open("pkg/day_3/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer puzzleInput.Close()

	reportLines := common.ReadLinesString(puzzleInput)

	transposedReport := TransposeReport(reportLines)

	p := PowerConsumption(transposedReport)
	fmt.Printf("The power consumption is %d.\n", p)

	l := LifeSupportRating(reportLines)
	fmt.Printf("The life support rating is %d.\n", l)
}

func PowerConsumption(transposedReport [][]int) int {
	var g []string
	var e []string
	for _, b := range transposedReport {
		g = append(g, gammaRate(b))
		e = append(e, epsilonRate(b))
	}

	gamma, err := strconv.ParseInt(strings.Join(g, ""), 2, 16)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(strings.Join(e, ""), 2, 16)
	if err != nil {
		log.Fatal(err)
	}

	return int(gamma) * int(epsilon)
}

func LifeSupportRating(report []string) int {
	o := OxygenGeneratorRating(report)
	c := Co2ScrubberRating(report)

	return o * c
}

func TransposeReport(reportLines []string) [][]int {
	numBits := len(reportLines[1])
	transposedReport := make([][]int, numBits)

	for _, line := range reportLines {
		for i := 0; i < len(line); i++ {
			bit, err := strconv.Atoi(string(line[i]))
			if err != nil {
				log.Fatal(err)
			}

			transposedReport[i] = append(transposedReport[i], bit)
		}
	}
	return transposedReport
}

func OxygenGeneratorRating(report []string) int {
	return lifeSupportRating(report, 0, true)
}

func Co2ScrubberRating(report []string) int {
	return lifeSupportRating(report, 0, false)
}

func lifeSupportRating(report []string, idx int, mostCommon bool) int {
	if len(report) == 1 {
		//fmt.Println(report)
		v, err := strconv.ParseInt(report[0], 2, 16)
		if err != nil {
			log.Fatal(err)
		}
		return int(v)
	}

	ones, zeros := 0, 0
	for _, i := range report {
		k, err := strconv.Atoi(string(i[idx]))
		if err != nil {
			log.Fatal(err)
		}
		if k == 1 {
			ones += 1
		} else {
			zeros += 1
		}
	}

	oneBit := "1"
	zeroBit := "0"
	if !mostCommon {
		oneBit, zeroBit = zeroBit, oneBit
	}
	var bit string
	if ones >= zeros {
		bit = oneBit
	} else {
		bit = zeroBit
	}

	keep := []string{}
	for _, row := range report {
		if string(row[idx]) == bit {
			keep = append(keep, row)
		}
	}

	return lifeSupportRating(keep, idx+1, mostCommon)
}

func gammaRate(bits []int) string {
	total := 0
	for _, i := range bits {
		total += i
		if total >= len(bits)/2 {
			return "1"
		}
	}
	return "0"
}

func epsilonRate(bits []int) string {
	total := 0
	for _, i := range bits {
		total += i
		if total >= len(bits)/2 {
			return "0"
		}
	}
	return "1"
}
