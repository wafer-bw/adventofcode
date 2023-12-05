package main

import (
	_ "embed"
	"fmt"
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

func Solve(input string) int {
	s := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(strings.Split(line, ": ")[1])
		parts := strings.Split(line, " | ")
		winners := strings.Split(strings.TrimSpace(parts[0]), " ")
		numbers := strings.Split(strings.TrimSpace(parts[1]), " ")

		fmt.Printf("%#v\n", winners)
		fmt.Printf("%#v\n", numbers)
		fmt.Println()

		pts := 0
		for _, winner := range winners {
			w := strings.TrimSpace(winner)
			if w == "" {
				continue
			}
			if slices.Contains(numbers, winner) {
				if pts == 0 {
					pts = 1
				} else {
					pts *= 2
				}
			}
		}
		s += pts
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
