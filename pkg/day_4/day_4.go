package day_X

import (
	"fmt"
	"log"
	"os"
)

func Run() {
	day := 4
	fmt.Printf("Day %d\n", day)
	route, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer route.Close()

	fmt.Printf("Day %d\n", day)
}
