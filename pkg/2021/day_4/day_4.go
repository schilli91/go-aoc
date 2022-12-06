package day_4

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/schilli91/go-aoc/pkg/common"
)

type BingoCard struct {
	rows   [5][5]string
	checks [5][5]bool
}

func NewBingoCard(rows [5]string) BingoCard {
	b := *new(BingoCard)
	for i, row := range rows {
		row = strings.ReplaceAll(strings.TrimSpace(row), "  ", " ")
		b.rows[i] = *(*[5]string)(strings.Split(row, " "))
	}

	return b
}

func (card *BingoCard) checkCard(value string) {
	for i := 0; i < len(card.rows); i++ {
		for j := 0; j < len(card.rows[i]); j++ {
			if value == card.rows[i][j] {
				card.checks[i][j] = true
			}
		}
	}
	//fmt.Printf("%v\n", card.checks)
}

func (card BingoCard) Bingo() bool {
	colBingo := []bool{true, true, true, true, true}

	for row := range card.checks {
		rowBingo := true
		for col, item := range card.checks[row] {
			if !item {
				rowBingo = false
				colBingo[col] = false
			}
		}
		if rowBingo {
			return true
		}
	}
	for _, bingo := range colBingo {
		if bingo {
			return true
		}
	}

	return false
}

func Run() {
	day := 4
	fmt.Printf("Day %d\n", day)
	data, err := os.Open(fmt.Sprintf("pkg/2021/day_%d/puzzle_input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	values, cards := LoadBingoCards(data)

	firstScore := PlayBingo(values, cards, false)
	fmt.Printf("The final Bingo score of the first winning card is %d\n", firstScore)

	lastScore := PlayBingo(values, cards, true)
	fmt.Printf("The final Bingo score of the last winning card is %d\n", lastScore)
}

func LoadBingoCards(data *os.File) ([]string, []BingoCard) {
	var values []string
	var cards []BingoCard

	lines := common.ReadLinesString(data)
	counter := 0
	for i, line := range lines {
		if i == 0 {
			values = strings.Split(line, ",")
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		if counter > 0 {
			counter--
			continue
		}
		rows := [5]string{
			lines[i],
			lines[i+1],
			lines[i+2],
			lines[i+3],
			lines[i+4],
		}
		c := NewBingoCard(rows)
		cards = append(cards, c)
		counter = 4
	}

	return values, cards
}

func (card BingoCard) finalScore(lastValue int) int {
	total := 0
	for i := 0; i < len(card.checks); i++ {
		for j := 0; j < len(card.checks[i]); j++ {
			if !card.checks[i][j] {
				item, err := strconv.Atoi(card.rows[i][j])
				if err != nil {
					log.Fatal(err)
				}
				total += item
			}
		}
	}
	return total * lastValue
}

func PlayBingo(values []string, cards []BingoCard, lastCard bool) int {
	won := []int{}
	for _, value := range values {
		for idx := range cards {
			card := &cards[idx]
			card.checkCard(value)

			if card.Bingo() {
				if lastCard {
					if contains(won, idx) {
						continue
					}
					won = append(won, idx)
					if len(won) != len(cards) {
						continue
					}
				}
				lastValue, err := strconv.Atoi(value)
				if err != nil {
					log.Fatal(err)
				}
				return card.finalScore(lastValue)
			}
		}
	}
	return 0
}

func contains(haystack []int, needle int) bool {
	for _, i := range haystack {
		if i == needle {
			return true
		}
	}
	return false
}
