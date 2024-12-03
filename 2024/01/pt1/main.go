package main

import (
	_ "embed"
	"log"
	"math"
	"slices"
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

	rows := 0
	locationIDs := make([][]int, 2)
	for _, line := range strings.Split(input, "\n") {
		for i, locationIDStr := range strings.Split(line, "   ") {
			idx := i % 2
			if idx == 0 {
				rows++
			}
			locationID, _ := strconv.Atoi(locationIDStr)
			locationIDs[idx] = append(locationIDs[idx], locationID)
		}
	}
	slices.Sort(locationIDs[0])
	slices.Sort(locationIDs[1])

	for i := range rows {
		s += int(math.Abs(float64(locationIDs[0][i]) - float64(locationIDs[1][i])))
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
