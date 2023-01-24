// https://adventofcode.com/2022/day/3#part2

package main

import (
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID    string = "2022-03"
	upperOffset int    = 38
	lowerOffset int    = 96
)

func priority(c rune) int {
	if strings.ToLower(string(c)) == string(c) {
		return int(c) - lowerOffset
	}

	return int(c) - upperOffset
}

func shared(a, b, c string) *rune {
	for _, r := range a {
		if strings.Contains(b, string(r)) && strings.Contains(c, string(r)) {
			return &r
		}
	}

	return nil
}

func solve(lines []string) int {
	sum := 0
	group := []string{}
	for _, ln := range lines {
		group = append(group, ln)
		if len(group) == 3 {
			sh := shared(group[0], group[1], group[2])
			if sh == nil {
				log.Fatal("no shared found")
			}
			sum += priority(*sh)
			group = []string{}
		}
	}

	return sum
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
