package day_X

import (
	"fmt"
	"log"
	"os"
)

func Run() {
	day := 0
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/2021/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	fmt.Printf("Day %d\n", day)
}
