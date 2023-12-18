package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
	//go:embed input-sample1.txt
	SampleInput1 string
	//go:embed input-sample2.txt
	SampleInput2 string
	//go:embed input-sample3.txt
	SampleInput3 string
	//go:embed input-sample4.txt
	SampleInput4 string
	//go:embed input.txt
	FullInput string
)

var (
	North vector.V2 = vector.V2{X: 0, Y: -1}
	East  vector.V2 = vector.V2{X: 1, Y: 0}
	South vector.V2 = vector.V2{X: 0, Y: 1}
	West  vector.V2 = vector.V2{X: -1, Y: 0}
)

var pipes = map[string][]vector.V2{
	".": nil,
	"|": {North, South},
	"-": {East, West},
	"L": {North, East},
	"J": {North, West},
	"7": {South, West},
	"F": {South, East},
	"S": {North, East, South, West},
}

func Solve(input string) int {
	startVec := vector.V2{X: 0, Y: 0}
	lines := strings.Split(input, "\n")
	pipeSketch := make([][]string, len(lines))
	for y, ln := range lines {
		for x, ch := range ln {
			chs := string(ch)
			if x == 0 {
				pipeSketch[y] = make([]string, len(ln))
			}
			pipeSketch[y][x] = string(chs)
			if chs == "S" {
				startVec = vector.V2{X: x, Y: y}
			}
		}
	}

	startConnections := getStartConnections(startVec, pipeSketch)

	at := vector.V2{X: startVec.X, Y: startVec.Y}
	heading := startConnections[0]
	path := []vector.V2{}
	for {
		from := vector.V2{X: at.X, Y: at.Y}
		at = at.Add(heading)
		path = append(path, at)
		for _, dir := range pipes[pipeSketch[at.Y][at.X]] {
			if at.Add(dir) == from {
				continue
			}
			heading = dir
		}

		if at == startVec {
			break
		}
	}

	pipeSketch = growGrid(pipeSketch)
	path = shiftPath(path)
	for _, row := range pipeSketch {
		fmt.Println(row)
	}
	fmt.Println()

	pipeSketch = floodFill(pipeSketch, path, vector.V2{}, "O")
	for _, row := range pipeSketch {
		fmt.Println(row)
	}
	fmt.Println()

	for y, row := range pipeSketch {
		for x, cell := range row {
			if cell == "*" {
				pipeSketch[y][x] = "O"
			} else if slices.Contains(path, vector.V2{X: x, Y: y}) {
				pipeSketch[y][x] = "O"
			}
		}
	}
	for _, row := range pipeSketch {
		fmt.Println(row)
	}
	fmt.Println()

	s := 0
	for _, row := range pipeSketch {
		for _, tile := range row {
			if tile != "O" {
				s++
			}
		}
	}

	return s
}

func floodFill(grid [][]string, path []vector.V2, fill vector.V2, reachedCh string) [][]string {
	if fill.X < 0 || fill.X >= len(grid[0]) || fill.Y < 0 || fill.Y >= len(grid) {
		return grid
	} else if slices.Contains(path, vector.V2{X: fill.X, Y: fill.Y}) {
		return grid
	} else if grid[fill.Y][fill.X] == reachedCh {
		return grid
	}

	grid[fill.Y][fill.X] = reachedCh

	grid = floodFill(grid, path, vector.V2{X: fill.X + 1, Y: fill.Y}, reachedCh)
	grid = floodFill(grid, path, vector.V2{X: fill.X - 1, Y: fill.Y}, reachedCh)
	grid = floodFill(grid, path, vector.V2{X: fill.X, Y: fill.Y + 1}, reachedCh)
	grid = floodFill(grid, path, vector.V2{X: fill.X, Y: fill.Y - 1}, reachedCh)

	return grid
}

func shiftPath(path []vector.V2) []vector.V2 {
	newPath := make([]vector.V2, len(path))
	for i := range path {
		newPath[i] = path[i].Add(vector.V2{X: 1, Y: 1})
	}
	return newPath
}

func growGrid(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid)+2)
	for y := range newGrid {
		newGrid[y] = make([]string, len(grid[0])+2)
		for x := range newGrid[y] {
			newGrid[y][x] = "X"
		}
	}

	for y, row := range grid {
		for x, cell := range row {
			newGrid[y+1][x+1] = cell
		}
	}

	return newGrid
}

func getStartConnections(start vector.V2, sketch [][]string) []vector.V2 {
	connections := []vector.V2{}
	for _, dir := range pipes["S"] {
		if start.Y+dir.Y < 0 || start.Y+dir.Y >= len(sketch) || start.X+dir.X < 0 || start.X+dir.X >= len(sketch[start.Y+dir.Y]) {
			continue
		}

		targetKey := sketch[start.Y+dir.Y][start.X+dir.X]
		target := pipes[targetKey]

		for _, targetDir := range target {
			if connected(dir, targetDir) {
				connections = append(connections, dir)
			}
		}
	}
	return connections
}

func connected(a, b vector.V2) bool {
	sub := a.Add(b)
	return sub.X == 0 && sub.Y == 0
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput1))
	log.Printf("sample: %d", Solve(SampleInput2))
	// log.Printf("sample: %d", Solve(SampleInput3))
	// log.Printf("sample: %d", Solve(SampleInput4))
	// log.Printf("full: %d", Solve(FullInput))
}
