package main

import (
	_ "embed"
	"log"
	"slices"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/test"
)

//go:embed inputs.txt
var Inputs string

var Cases test.Cases = test.GetCases(Inputs)

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
	filter := []int{}
	for i, c := range Cases {
		if slices.Contains(filter, i) {
			continue
		}
		log.Printf("case=%d expect=%d got=%d", i, c.Answer, Solve(c.Input))
	}
}
