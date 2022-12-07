package day_2

import (
	"fmt"
	"strings"

	"github.com/schilli91/go-aoc/pkg/common"
)

func Run() {
	day := 1
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile("pkg/2022/day_2/puzzle_input.txt")
	defer input.Close()

	strategies := common.ReadLinesString(input)
	tournamentScore := getTournamentScore(getMatches(strategies))
	fmt.Printf("The tournament will end with a score of %d for me.\n", tournamentScore)

	correctedScore := getTournamentScore(getCorrectedMatches(strategies))
	fmt.Printf("The tournament will end with a score of %d for me after correcting my move.\n", correctedScore)
}

const (
	opponentRock    = "A"
	opponentPaper   = "B"
	opponentScissor = "C"
	myRock          = "X"
	myPaper         = "Y"
	myScissor       = "Z"
	expectedWin     = "Z"
	expectedDraw    = "Y"
	expectedLose    = "X"
)

type Match struct {
	myMove         string
	opponentMove   string
	expectedResult string
}

func NewMatch(strategy string) *Match {
	m := new(Match)
	f := strings.Fields(strategy)
	m.myMove = f[1]
	m.opponentMove = f[0]
	m.expectedResult = f[1]
	return m
}

func (m *Match) correctMyMove() {
	switch m.expectedResult {
	case expectedWin:
		switch m.opponentMove {
		case opponentRock:
			m.myMove = myPaper
		case opponentPaper:
			m.myMove = myScissor
		case opponentScissor:
			m.myMove = myRock
		}
	case expectedDraw:
		switch m.opponentMove {
		case opponentRock:
			m.myMove = myRock
		case opponentPaper:
			m.myMove = myPaper
		case opponentScissor:
			m.myMove = myScissor
		}
	case expectedLose:
		switch m.opponentMove {
		case opponentRock:
			m.myMove = myScissor
		case opponentPaper:
			m.myMove = myRock
		case opponentScissor:
			m.myMove = myPaper
		}
	}
}

func (m Match) getScore() int {
	switch m.myMove {
	case myRock:
		return 1 + evalRockOpponent(m.opponentMove)
	case myPaper:
		return 2 + evalPaperOpponent(m.opponentMove)
	case myScissor:
		return 3 + evalScissorOpponent(m.opponentMove)
	}

	return 0
}

func evalRockOpponent(opponentMove string) int {
	switch opponentMove {
	case opponentRock:
		return 3
	case opponentPaper:
		return 0
	case opponentScissor:
		return 6
	}
	return 0
}

func evalPaperOpponent(opponentMove string) int {
	switch opponentMove {
	case opponentRock:
		return 6
	case opponentPaper:
		return 3
	case opponentScissor:
		return 0
	}
	return 0
}

func evalScissorOpponent(opponentMove string) int {
	switch opponentMove {
	case opponentRock:
		return 0
	case opponentPaper:
		return 6
	case opponentScissor:
		return 3
	}
	return 0
}

func getMatches(strategies []string) []Match {
	matches := make([]Match, 0)

	for _, s := range strategies {
		matches = append(matches, *NewMatch(s))
	}
	return matches
}

func getCorrectedMatches(strategies []string) []Match {
	matches := make([]Match, 0)

	for _, s := range strategies {
		m := *NewMatch(s)
		m.correctMyMove()
		matches = append(matches, m)
	}
	return matches
}

func getTournamentScore(matches []Match) int {
	sum := 0
	for _, m := range matches {
		sum += m.getScore()
	}
	return sum
}
