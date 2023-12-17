package main

import (
	_ "embed"
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/m"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

var (
	startPattern = regexp.MustCompile(`A$`)
	endPattern   = regexp.MustCompile(`Z$`)
	instructMap  = map[rune]int{'L': 0, 'R': 1}
)

type node struct {
	key    string
	nexts  []string
	period int
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	instructions := []int{}
	for _, ch := range strings.TrimSpace(lines[0]) {
		instructions = append(instructions, instructMap[ch])
	}

	starts := []*node{}
	nodes := map[string]*node{}
	for _, ln := range lines[2:] {
		parts := strings.Split(ln, " = ")
		nodesStr := strings.TrimPrefix(parts[1], "(")
		nodesStr = strings.TrimSuffix(nodesStr, ")")
		node := &node{
			key:   parts[0],
			nexts: strings.Split(nodesStr, ", "),
		}
		nodes[parts[0]] = node
		if startPattern.MatchString(parts[0]) {
			starts = append(starts, node)
		}
	}

	for _, start := range starts {
		current := *nodes[start.key]
		for i := 0; i < math.MaxInt16; i++ { // run a long loop instead of an infinite one
			instruction := instructions[i%len(instructions)]
			current = *nodes[current.nexts[instruction]]
			start.period++
			if endPattern.MatchString(current.key) {
				break
			}
		}
	}

	lcm := starts[0].period
	for _, start := range starts[1:] {
		lcm = m.LeastCommonMultiple(lcm, start.period)
	}

	return lcm
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
