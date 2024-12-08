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

var operators map[string]func(int, int) int = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

func Solve(input string) int {
	s := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		x, _ := strconv.Atoi(strings.TrimSuffix(parts[0], ":"))
		valueStrings := parts[1:]
		values := make([]int, len(valueStrings))
		for i, valueString := range valueStrings {
			values[i], _ = strconv.Atoi(valueString)
		}

		if resolve(x, values[0], 1, values) {
			s += x
		}
	}

	return s
}

func resolve(ans, cumulative, idx int, values []int) bool {
	if idx == len(values) {
		return ans == cumulative
	}
	for _, opf := range operators {
		c := opf(cumulative, values[idx])
		if c > ans {
			continue
		} else if resolve(ans, c, idx+1, values) {
			return true
		}
	}
	return false
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
