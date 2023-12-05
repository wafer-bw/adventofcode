package main

import (
	_ "embed"
	"log"
	"slices"
	"strings"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type Card struct {
	amount  int
	matches int
}

func Solve(input string) int {
	s := 0
	lines := strings.Split(input, "\n")
	cards := []*Card{}
	for _, line := range lines {
		line = strings.TrimSpace(strings.Split(line, ": ")[1])
		parts := strings.Split(line, " | ")
		winners := strings.Split(strings.TrimSpace(parts[0]), " ")
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

		matches := 0
		for _, winner := range winners {
			w := strings.TrimSpace(winner)
			if w == "" {
				continue
			}
			if slices.Contains(numbers, w) {
				matches++
			}
		}

		cards = append(cards, &Card{amount: 1, matches: matches})
	}

	for cn, card := range cards {
		cardNumber := cn + 1
		for j := 1; j <= card.amount; j++ {
			for i := 1 + cardNumber; i <= cardNumber+card.matches; i++ {
				cards[i-1].amount++
			}
		}
	}

	for _, card := range cards {
		s += card.amount
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
