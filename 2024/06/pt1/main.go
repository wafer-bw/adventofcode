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

type Map [][]Tile

func (m Map) String() string {
	str := ""
	for _, row := range m {
		strs := make([]string, len(row))
		for i, tile := range row {
			strs[i] = string(tile)
		}
		str += strings.Join(strs, "") + "\n"
	}
	return str
}

type Tile string

const (
	TileEmpty      Tile = "."
	TileObstacle   Tile = "#"
	TileGuardStart Tile = "^"
	TileVisited    Tile = "X"
)

func Solve(input string) int {
	s := 0
	lines := strings.Split(input, "\n")

	dir := vector.Cardinal2North // It's facing north to pos in both my inputs.
	pos := vector.V2{X: 0, Y: 0}
	m := make(Map, len(lines))
	for y, line := range lines {
		tiles := strings.Split(line, "")
		m[y] = make([]Tile, len(tiles))
		for x, tile := range tiles {
			m[y][x] = Tile(tile)
			if m[y][x] == TileGuardStart {
				pos = vector.V2{X: x, Y: y}
				// could set dir here if we needed to determine it from the input.
			}
		}
	}

	for {
		if m[pos.Y][pos.X] != TileVisited {
			m[pos.Y][pos.X] = TileVisited
			s++
		}
		next := pos.Translate(1, dir)
		if next.Y < 0 || next.Y >= len(m) || next.X < 0 || next.X >= len(m[next.Y]) {
			break
		} else if m[next.Y][next.X] == TileObstacle {
			dir = dir.RotateRight()
		} else {
			pos = next
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
