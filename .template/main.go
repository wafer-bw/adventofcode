// https://adventofcode.com/XXXX/day/X

package main

import (
	"log"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID string = "XXXX-XX"
)

func solve(lines []string) int {
	return 0
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
