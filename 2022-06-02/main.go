package main

import (
	"log"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"github.com/wafer-bw/adventofcode/tools/stack"
)

const (
	puzzleID  string = "2022-06"
	markerLen int    = 14
)

func solve(lines []string) int {
	ln := lines[0]
	marker := []string{}
	for i, ch := range ln {
		if len(marker) <= markerLen-1 {
			marker = append(marker, string(ch))
		} else {
			stack.Push(&marker, string(ch))
			stack.PopBack(&marker)
		}

		if len(marker) == markerLen && allUnique(marker) {
			return i + 1
		}
	}

	return -1
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func allUnique(list []string) bool {
	m := map[string]struct{}{}
	for _, v := range list {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
