package main

import (
	_ "embed"
	"log"
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
	for i, line := range lines {
		_ = i
		_ = line
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
