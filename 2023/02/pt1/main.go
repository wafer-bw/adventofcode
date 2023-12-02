package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

// var bag = map[string]int{}
var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Solve(input string) int {
	s := 0

Game:
	for _, line := range strings.Split(input, "\n") {
		gameID, _ := strconv.Atoi(strings.TrimPrefix(strings.Split(line, ":")[0], "Game "))
		sets := strings.Split(strings.Split(line, ":")[1], ";")

		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, colorStr := range colors {
				parts := strings.Split(strings.TrimSpace(colorStr), " ")
				numColor, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
				color := parts[1]

				if limits[color] < numColor {
					continue Game
				}
			}
		}
		s += gameID
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	// log.Printf("full: %d", Solve(FullInput))
}
