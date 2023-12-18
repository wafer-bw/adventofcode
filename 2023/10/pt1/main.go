package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
	//go:embed input-sample1.txt
	SampleInput1 string
	//go:embed input-sample2.txt
	SampleInput2 string
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

	// fmt.Println(startVec)
	// fmt.Println(pipeSketch)
	// fmt.Println(startConnections)

	steps := 0
	at := vector.V2{X: startVec.X, Y: startVec.Y}
	heading := startConnections[0]
	for {
		steps++
		from := vector.V2{X: at.X, Y: at.Y}
		at = at.Add(heading)
		// fmt.Println(from, pipeSketch[from.Y][from.X], "-", heading, "->", pipeSketch[at.Y][at.X], at)
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

	return steps / 2
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
	log.Printf("full: %d", Solve(FullInput))
}
