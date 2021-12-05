package day_4

import (
	"fmt"
	"log"
	"os"
)

func Run() {
	day := 4
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	score := PlayBingo(data)
	fmt.Printf("The final Bingo score is %d\n", score)
}

func PlayBingo(data *os.File) int {

	return 0
}
