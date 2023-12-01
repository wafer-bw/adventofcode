package main

import (
	_ "embed"
	"fmt"
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
	result := 0

	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
	}

	return result
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
