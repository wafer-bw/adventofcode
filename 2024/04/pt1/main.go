package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type Grid [][]string

func (g Grid) Pos(v vector.V2) string {
	if v.Y < 0 || v.Y >= len(g) {
		return ""
	} else if v.X < 0 || v.X >= len(g[v.Y]) {
		return ""
	}
	return g[v.Y][v.X]
}

func (g Grid) Search(word string, at vector.V2) int {
	searchers := []func(string, vector.V2) int{
		g.searchRight,
		g.searchDown,
		g.searchLeft,
		g.searchUp,
		g.searchUpRight,
		g.searchDownRight,
		g.searchDownLeft,
		g.searchUpLeft,
	}

	total := 0
	for _, searcher := range searchers {
		total += searcher(word, at)
	}
	return total
}

func (g Grid) searchRight(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X + i, Y: at.Y}) != string(char) {
			return 0
		}
	}
	return 1
}

func (g Grid) searchDown(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X, Y: at.Y + i}) != string(char) {
			return 0
		}
	}
	return 1
}

func (g Grid) searchLeft(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X - i, Y: at.Y}) != string(char) {
			return 0
		}
	}
	return 1
}

func (g Grid) searchUp(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X, Y: at.Y - i}) != string(char) {
			return 0
		}
	}
	return 1
}

func (g Grid) searchUpRight(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X + i, Y: at.Y - i}) != string(char) {
			return 0
		}
	}
	return 1
}

func (g Grid) searchDownRight(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X + i, Y: at.Y + i}) != string(char) {
			return 0
		}
	}
	return 1
}

func (g Grid) searchDownLeft(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X - i, Y: at.Y + i}) != string(char) {
			return 0
		}
	}
	return 1
}

func (g Grid) searchUpLeft(word string, at vector.V2) int {
	for i, char := range word {
		if g.Pos(vector.V2{X: at.X - i, Y: at.Y - i}) != string(char) {
			return 0
		}
	}
	return 1
}

func Solve(input string) int {
	s := 0
	word := "XMAS"

	lines := strings.Split(input, "\n")
	grid := make(Grid, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	for y, row := range grid {
		for x := range row {
			s += grid.Search(word, vector.V2{X: x, Y: y})
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
