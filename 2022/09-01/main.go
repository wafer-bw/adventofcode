// https://adventofcode.com/2022/day/9

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"github.com/wafer-bw/adventofcode/tools/vector"
)

const (
	puzzleID      string = "2022-09"
	tailMovespeed int    = 1
)

func solve(lines []string) int {
	m := map[int]map[int]struct{}{}
	h := vector.V2{0, 0}
	t := vector.V2{0, 0}

	for _, ln := range lines {
		parts := strings.Split(ln, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		for i := 0; i < steps; i++ {
			if _, ok := m[t.X]; !ok {
				m[t.X] = map[int]struct{}{}
			}
			if _, ok := m[t.X][t.Y]; !ok {
				m[t.X][t.Y] = struct{}{}
			}

			h = move(h, dir)
			separation := t.OrthoDistance(h)
			distance := t.Distance(h)
			if distance >= 2 {
				t = translate(t, vector.V2{-separation.X, -separation.Y}, tailMovespeed)
			}
		}
	}

	positions := map[string]struct{}{}
	for x, X := range m {
		for y := range X {
			positions[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
		}
	}

	return len(positions)
}

func draw(m map[int]map[int]struct{}, h vector.V2, t vector.V2) {
	for y := 6; y >= 0; y-- {
		for x := 0; x < 6; x++ {
			if x == h.X && y == h.Y {
				fmt.Print("H")
				continue
			}
			if x == t.X && y == t.Y {
				fmt.Print("T")
				continue
			}

			if _, ok := m[x]; ok {
				if _, ok := m[x][y]; ok {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func translate(pos, dir vector.V2, max int) vector.V2 {
	if dir.X > max {
		dir.X = max
	} else if dir.X < -max {
		dir.X = -max
	}

	if dir.Y > max {
		dir.Y = max
	} else if dir.Y < -max {
		dir.Y = -max
	}

	return vector.V2{pos.X + dir.X, pos.Y + dir.Y}
}

func move(pos vector.V2, dir string) vector.V2 {
	switch dir {
	case "U":
		pos.Y++
	case "R":
		pos.X++
	case "D":
		pos.Y--
	case "L":
		pos.X--
	}

	return pos
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
