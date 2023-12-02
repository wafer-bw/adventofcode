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

func Solve(input string) int {
	s := 0

	for _, line := range strings.Split(input, "\n") {
		_, _ = strconv.Atoi(strings.TrimPrefix(strings.Split(line, ":")[0], "Game "))
		sets := strings.Split(strings.Split(line, ":")[1], ";")

		var max = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range sets {
			colors := strings.Split(set, ",")
			for _, colorStr := range colors {
				parts := strings.Split(strings.TrimSpace(colorStr), " ")
				numColor, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
				color := parts[1]

				if max[color] < numColor {
					max[color] = numColor
				}
			}
		}

		s += max["red"] * max["green"] * max["blue"]
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
