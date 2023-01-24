// https://adventofcode.com/2022/day/9

package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID      string = "2022-09"
	tailMovespeed int    = 1
)

type v2 struct{ x, y int }

func (a v2) orthoDistance(b v2) v2 {
	dx := a.x - b.x
	dy := a.y - b.y
	return v2{dx, dy}
}

func (a v2) distance(b v2) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}

func solve(lines []string) int {
	m := map[int]map[int]struct{}{}
	h := v2{0, 0}
	t := v2{0, 0}

	for _, ln := range lines {
		parts := strings.Split(ln, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		for i := 0; i < steps; i++ {
			if _, ok := m[t.x]; !ok {
				m[t.x] = map[int]struct{}{}
			}
			if _, ok := m[t.x][t.y]; !ok {
				m[t.x][t.y] = struct{}{}
			}

			h = move(h, dir)
			separation := t.orthoDistance(h)
			distance := t.distance(h)
			if distance >= 2 {
				t = translate(t, v2{-separation.x, -separation.y}, tailMovespeed)
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

func draw(m map[int]map[int]struct{}, h v2, t v2) {
	for y := 6; y >= 0; y-- {
		for x := 0; x < 6; x++ {
			if x == h.x && y == h.y {
				fmt.Print("H")
				continue
			}
			if x == t.x && y == t.y {
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

func translate(pos, dir v2, max int) v2 {
	if dir.x > max {
		dir.x = max
	} else if dir.x < -max {
		dir.x = -max
	}

	if dir.y > max {
		dir.y = max
	} else if dir.y < -max {
		dir.y = -max
	}

	return v2{pos.x + dir.x, pos.y + dir.y}
}

func move(pos v2, dir string) v2 {
	switch dir {
	case "U":
		pos.y++
	case "R":
		pos.x++
	case "D":
		pos.y--
	case "L":
		pos.x--
	}

	return pos
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
