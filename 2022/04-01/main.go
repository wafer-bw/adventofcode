// https://adventofcode.com/2022/day/4

package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID string = "2022-04"
)

func solve(lines []string) int {
	count := 0

	for _, ln := range lines {
		pair := [][]int{}
		pairParts := strings.Split(ln, ",")
		for _, part := range pairParts {
			rangeParts := strings.Split(part, "-")
			begin, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			pair = append(pair, []int{begin, end})
		}

		if pair[0][0] <= pair[1][0] && pair[0][1] >= pair[1][1] || pair[0][0] >= pair[1][0] && pair[0][1] <= pair[1][1] {
			count++
		}
	}

	return count
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
