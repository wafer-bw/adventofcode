package main

import (
	_ "embed"
	"fmt"
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
		numbers := []int{}
		for _, ch := range line {
			if n, err := strconv.Atoi(string(ch)); err == nil {
				numbers = append(numbers, n)
			}
		}
		value := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
		numericalValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		s += numericalValue
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
