package day_2

import (
	"fmt"
	"testing"

	"github.com/schilli91/go-aoc/pkg/common"
)

func TestNewMatch(t *testing.T) {
	m := NewMatch("A Y")
	got := m.myMove
	want := myPaper
	if got != want {
		t.Errorf("Mismatch of my move: got %s, want %s", got, want)
	}
	got = m.opponentMove
	want = opponentRock
	if got != want {
		t.Errorf("Mismatch of my move: got %s, want %s", got, want)
	}

	got = m.expectedResult
	want = expectedDraw
	if got != want {
		t.Errorf("Mismatch of my move: got %s, want %s", got, want)
	}
}

func TestGetScore(t *testing.T) {
	m := NewMatch("A Y")
	got := m.getScore()
	want := 8
	if got != want {
		t.Errorf("Mismatch of match score: got %d, want %d", got, want)
	}
}

func TestGetCorrectedScore(t *testing.T) {
	m := NewMatch("A Y")
	m.correctMyMove()
	got := m.getScore()
	want := 4
	if got != want {
		t.Errorf("Mismatch of match score: got %d, want %d", got, want)
	}
}

func TestGetTournamentScore(t *testing.T) {
	day := 1
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	strategies := common.ReadLinesString(input)
	got := getTournamentScore(getMatches(strategies))
	want := 15
	if got != want {
		t.Errorf("Mismatch of number of elves: got %d, want %d", got, want)
	}
}

func TestGetCorrectedTournamentScore(t *testing.T) {
	day := 1
	fmt.Printf("Day %d of AOC 2022.\n", day)

	input := common.OpenInputFile("test_data.txt")
	defer input.Close()
	strategies := common.ReadLinesString(input)
	got := getTournamentScore(getCorrectedMatches(strategies))
	want := 12
	if got != want {
		t.Errorf("Mismatch of number of elves: got %d, want %d", got, want)
	}
}
