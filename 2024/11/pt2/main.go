package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

var (
	//go:embed input-sample1.txt
	SampleInput1 string
	//go:embed input-sample2.txt
	SampleInput2 string
	//go:embed input.txt
	FullInput string
)

// Stones maps each [Stone] to the number of occurrences of it.
type Stones map[Stone]int

func (ss Stones) Next() Stones {
	next := Stones{}

	if _, ok := ss[0]; ok {
		next[1] += ss[0]
		delete(ss, 0)
	}

	for s, count := range ss {
		sStr := strconv.Itoa(int(s))
		if len(sStr)%2 == 0 {
			half := len(sStr) / 2
			left, _ := strconv.Atoi(sStr[:half])
			right, _ := strconv.Atoi(sStr[half:])
			next[Stone(left)] += count
			next[Stone(right)] += count
		} else {
			next[s*2024] += count
		}
	}
	return next
}

type Stone int

func Solve(input string) int {
	stones := Stones{}
	for _, stoneStr := range strings.Split(input, " ") {
		n, _ := strconv.Atoi(stoneStr)
		stones[Stone(n)] += 1
	}

	s := 0
	blinks := 75
	for blink := 0; blink < blinks; blink++ {
		stones = stones.Next()
	}

	for _, count := range stones {
		s += count
	}

	return s
}

func main() {
	log.Printf("sample 1: %d", Solve(SampleInput1))
	log.Printf("sample 2: %d", Solve(SampleInput2))
	log.Printf("full: %d", Solve(FullInput))
}
