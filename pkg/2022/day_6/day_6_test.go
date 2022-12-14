package day_6

import (
	"fmt"
	"testing"
)

var testStartOfPacket = []struct {
	dataStream  string
	packetStart int
}{
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
}

var testStartOfMessage = []struct {
	dataStream   string
	messageStart int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
}

func TestFindPacketMarker(t *testing.T) {
	for _, d := range testStartOfPacket {
		t.Run(fmt.Sprintf("%v", d), func(t *testing.T) {
			got := findPacketMarker(d.dataStream)
			want := d.packetStart
			if got != want {
				t.Errorf("wrong packet start, got %d, want %d", got, want)
			}
		})
	}
}

func TestFindMessageMarker(t *testing.T) {
	for _, d := range testStartOfMessage {
		t.Run(fmt.Sprintf("%v", d), func(t *testing.T) {
			got := findMessageMarker(d.dataStream)
			want := d.messageStart
			if got != want {
				t.Errorf("wrong packet start, got %d, want %d", got, want)
			}
		})
	}
}
