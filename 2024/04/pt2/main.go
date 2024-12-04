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
	total := 0
	if g.searchDownRight(word, at) == 1 && g.searchUpRight(word, vector.V2{X: at.X, Y: at.Y + 2}) == 1 {
		total++
	}

	if g.searchDownLeft(word, at) == 1 && g.searchUpLeft(word, vector.V2{X: at.X, Y: at.Y + 2}) == 1 {
		total++
	}

	if g.searchDownRight(word, at) == 1 && g.searchDownLeft(word, vector.V2{X: at.X + 2, Y: at.Y}) == 1 {
		total++
	}

	if g.searchUpRight(word, at) == 1 && g.searchUpLeft(word, vector.V2{X: at.X + 2, Y: at.Y}) == 1 {
		total++
	}
	return total
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
	word := "MAS"

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
