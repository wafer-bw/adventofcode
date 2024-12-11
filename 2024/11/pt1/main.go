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

type Stones []Stone

func (ss Stones) Next() Stones {
	newList := make(Stones, 0, len(ss))
	for _, s := range ss {
		newList = append(newList, s.Next()...)
	}
	return newList
}

type Stone int

func (s Stone) Next() Stones {
	if s == 0 {
		return Stones{1}
	}

	sStr := strconv.Itoa(int(s))
	if len(sStr)%2 == 0 {
		half := len(sStr) / 2
		left, _ := strconv.Atoi(sStr[:half])
		right, _ := strconv.Atoi(sStr[half:])
		return Stones{Stone(left), Stone(right)}
	}

	return Stones{s * 2024}
}

func Solve(input string) int {
	stones := Stones{}
	for _, stoneStr := range strings.Split(input, " ") {
		n, _ := strconv.Atoi(stoneStr)
		stones = append(stones, Stone(n))
	}

	blinks := 25
	for blink := 0; blink < blinks; blink++ {
		stones = stones.Next()
	}

	return len(stones)
}

func main() {
	log.Printf("sample 1: %d", Solve(SampleInput1))
	log.Printf("sample 2: %d", Solve(SampleInput2))
	log.Printf("full: %d", Solve(FullInput))
}
