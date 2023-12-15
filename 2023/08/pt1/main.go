package main

import (
	_ "embed"
	"log"
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

var instructMap = map[rune]int{
	'L': 0,
	'R': 1,
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	instructions := []int{}
	for _, ch := range strings.TrimSpace(lines[0]) {
		instructions = append(instructions, instructMap[ch])
	}

	nodes := map[string][]string{}
	for _, ln := range lines[2:] {
		parts := strings.Split(ln, " = ")
		nodesStr := strings.TrimPrefix(parts[1], "(")
		nodesStr = strings.TrimSuffix(nodesStr, ")")
		nodes[parts[0]] = strings.Split(nodesStr, ", ")
	}

	i := 0
	s := 0
	cur := "AAA"
	end := "ZZZ"
	for {
		instruction := instructions[i%len(instructions)]
		cur = nodes[cur][instruction]
		s++
		i++

		if cur == end {
			return s
		}
	}
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput1))
	log.Printf("sample: %d", Solve(SampleInput2))
	log.Printf("full: %d", Solve(FullInput))
}
