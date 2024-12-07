package main

// < 1821

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

type Tile string

const (
	TileEmpty      Tile = "."
	TileObstacle   Tile = "#"
	TileGuardStart Tile = "^"
	TileVisited    Tile = "X"
)

type Step struct {
	Position  vector.V2
	Direction vector.V2
}

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
				m[y][x] = TileEmpty
				// could set dir here if we needed to determine it from the input.
			}
		}
	}

	steps := Traverse(m, pos, dir)
	obstructions := map[vector.V2]struct{}{}
	for i, step := range steps {
		log.Printf("step %d / %d", i, len(steps))
		if i == len(steps)-1 {
			continue
		} else if _, ok := obstructions[steps[i+1].Position]; ok {
			continue
		}
		l := m[steps[i+1].Position.Y][steps[i+1].Position.X]
		m[steps[i+1].Position.Y][steps[i+1].Position.X] = TileObstacle
		obstructions[steps[i+1].Position] = struct{}{}
		s += LoopTraverse(m, s, step.Position, step.Direction)
		m[steps[i+1].Position.Y][steps[i+1].Position.X] = l
	}

	return s
}

func Traverse(m Map, pos, dir vector.V2) []Step {
	startTracking := false
	steps := []Step{{Position: pos, Direction: dir}}
	for {
		next := pos.Translate(1, dir)
		if next.Y < 0 || next.Y >= len(m) || next.X < 0 || next.X >= len(m[next.Y]) {
			return steps
		} else if m[next.Y][next.X] == TileObstacle {
			dir = dir.RotateRight()
			startTracking = true
		} else {
			if startTracking {
				steps = append(steps, Step{Position: next, Direction: dir})
			}
			pos = next
		}
	}
}

func LoopTraverse(m Map, n int, pos, dir vector.V2) int {
	seens := map[vector.V2]map[vector.V2]struct{}{}
	for {
		if _, ok := seens[pos][dir]; ok {
			return 1
		}
		if _, ok := seens[pos]; !ok {
			seens[pos] = map[vector.V2]struct{}{}
		}
		seens[pos][dir] = struct{}{}
		next := pos.Translate(1, dir)
		if next.Y < 0 || next.Y >= len(m) || next.X < 0 || next.X >= len(m[next.Y]) {
			return 0
		} else if m[next.Y][next.X] == TileObstacle {
			dir = dir.RotateRight()
		} else {
			pos = next
		}
	}
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
