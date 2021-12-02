package day_X

import (
	"fmt"
	"log"
	"os"
)

func Run() {
	fmt.Printf("Day X\n")
	route, err := os.Open("pkg/day_2/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer route.Close()

	fmt.Printf("Day %d\n", 3)
}
