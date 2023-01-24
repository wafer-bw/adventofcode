// https://adventofcode.com/2022/day/4#part2

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
		var overlaps bool
		filter := map[int]struct{}{}
		pairParts := strings.Split(ln, ",")
		for _, part := range pairParts {
			rangeParts := strings.Split(part, "-")
			begin, _ := strconv.Atoi(rangeParts[0])
			end, _ := strconv.Atoi(rangeParts[1])
			filter, overlaps = hasOverlaps(filter, begin, end)
			if overlaps {
				overlaps = false
				count++
				break
			}
		}
	}

	return count
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func hasOverlaps(filter map[int]struct{}, s, e int) (map[int]struct{}, bool) {
	for i := s; i <= e; i++ {
		if _, ok := filter[i]; !ok {
			filter[i] = struct{}{}
		} else {
			return nil, true
		}
	}
	return filter, false
}
