package main

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/wafer-bw/always"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

var cardValues map[rune]int = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type HandType int

const (
	HighCard HandType = iota
	TwoOfAKind
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards []int
	Bid   int
}

func (h *Hand) Parse(line string) {
	parts := strings.Split(line, " ")
	h.Bid = always.Accept(strconv.Atoi(parts[1]))
	for _, card := range parts[0] {
		h.Cards = append(h.Cards, cardValues[card])
	}
}

func (h Hand) bucket() map[int]int {
	bucket := map[int]int{}
	for _, card := range h.Cards {
		if _, ok := bucket[card]; !ok {
			bucket[card] = 0
		}
		bucket[card]++
	}
	return bucket
}

func (h *Hand) Type() int {
	bucket := h.bucket()
	bcounts := []int{}
	for _, count := range bucket {
		bcounts = append(bcounts, count)
	}

	if len(bucket) == 1 {
		return int(FiveOfAKind)
	}

	if len(bucket) == 5 {
		return int(HighCard)
	}

	if len(bucket) == 2 && (bcounts[0] == 2 && bcounts[1] == 3) || (bcounts[0] == 3 && bcounts[1] == 2) {
		return int(FullHouse)
	}

	if len(bucket) == 3 && slices.Contains(bcounts, 1) && slices.Contains(bcounts, 2) {
		return int(TwoPair)
	}

	t := 0
	for _, count := range bucket {
		t = max(t, count)
	}

	switch t {
	case 2:
		return int(TwoOfAKind)
	case 3:
		return int(ThreeOfAKind)
	case 4:
		return int(FourOfAKind)
	}

	panic("unknown type for hand")
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0, len(lines))
	for _, line := range lines {
		hand := Hand{}
		hand.Parse(line)
		hands = append(hands, hand)
	}

	slices.SortStableFunc(hands, func(a, b Hand) int {
		if a.Type() == b.Type() {
			for i := 0; i < len(a.Cards); i++ {
				if a.Cards[i] == b.Cards[i] {
					continue
				}
				return a.Cards[i] - b.Cards[i]
			}
		}
		return a.Type() - b.Type()
	})

	s := 0
	for i, hand := range hands {
		rank := i + 1
		s += hand.Bid * rank
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
