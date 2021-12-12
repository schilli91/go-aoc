package day_8

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/schilli91/aoc2021/pkg/common"
)

func Run() {
	day := 8
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	lines := common.ReadLinesString(data)

	d := LoadDisplays(lines)
	n := CountUniqueDigits(d)

	fmt.Printf("There are %d unique digits (1, 4, 7, 8) in the outputs.\n", n)

	o := DecodeDisplays(d)
	fmt.Printf("The sum of outputs is %d.\n", o)
}

type Display struct {
	digits  []string
	outputs []string
}

func NewDisplay(l string) Display {
	parts := strings.Split(l, "|")
	d := strings.TrimSpace(parts[0])
	o := strings.TrimSpace(parts[1])
	return Display{
		digits:  strings.Split(d, " "),
		outputs: strings.Split(o, " "),
	}
}

func LoadDisplays(lines []string) []Display {
	displays := []Display{}
	for _, l := range lines {
		displays = append(displays, NewDisplay(l))
	}
	return displays
}

func CountUniqueDigits(displays []Display) int {
	total := 0
	for _, d := range displays {
		for _, o := range d.outputs {
			if len(o) == 2 || len(o) == 3 || len(o) == 4 || len(o) == 7 {
				total++
			}
		}
	}

	return total
}

func DecodeDisplays(displays []Display) int {
	total := 0

	for _, d := range displays {
		total += d.Decode()
	}

	return total
}

func (display Display) Decode() int {
	wiring := wiring(display.digits)
	outputDigits := []int{}

	for _, o := range display.outputs {
		for i, d := range display.digits {
			if countSegmentsInDigit(o, d) == len(o) && countSegmentsInDigit(o, d) == len(d) {
				outputDigits = append(outputDigits, wiring[i])
				break
			}
		}
	}

	return digitsToNumber(outputDigits)
}

func digitsToNumber(digits []int) int {
	s := []string{}
	for _, d := range digits {
		s = append(s, strconv.Itoa(d))
	}
	numberString := strings.Join(s, "")
	n, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatalf("Failed to convert: %s", numberString)
	}
	return n
}

func getValue(idx int, wiring []int) (int, error) {
	for j, v := range wiring {
		if idx == v {
			return j, nil
		}
	}
	return 0, errors.New("Failed.")
}

func wiring(digits []string) []int {
	wiring := uniqueWirings(digits)
	for i, d := range digits {
		// 1, 4, 7, 8
		if len(d) == 5 {
			// 2, 3, 5
			if numInDigit(digits[wiring[7]], d) {
				wiring[3] = i
				continue
			}
			if countSegmentsInDigit(digits[wiring[4]], d) == 3 {
				wiring[5] = i
				continue
			}
			wiring[2] = i
			continue
		}
		if len(d) == 6 {
			// 0, 6, 9
			if numInDigit(digits[wiring[4]], d) {
				wiring[9] = i
				continue
			}
			if numInDigit(digits[wiring[7]], d) {
				wiring[0] = i
				continue
			}
			wiring[6] = i
			continue
		}
	}
	return indexNumberWiring(wiring)
}

func indexNumberWiring(wiring []int) []int {
	w := make([]int, 10)
	for i, v := range wiring {
		w[v] = i
	}
	return w
}

func uniqueWirings(digits []string) []int {
	wiring := make([]int, 10)
	for i, d := range digits {
		if len(d) == 2 {
			wiring[1] = i
		}
		if len(d) == 3 {
			wiring[7] = i
		}
		if len(d) == 4 {
			wiring[4] = i
		}
		if len(d) == 7 {
			wiring[8] = i
		}
	}
	return wiring
}

func numInDigit(num string, digit string) bool {
	for _, c := range num {
		if !strings.ContainsRune(digit, c) {
			return false
		}
	}
	return true
}

func countSegmentsInDigit(num string, digit string) int {
	total := 0
	for _, c := range num {
		if strings.ContainsRune(digit, c) {
			total++
		}
	}
	return total
}
