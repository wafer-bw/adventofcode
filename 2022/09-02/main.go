// https://adventofcode.com/2022/day/9#part2

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
	knotMovespeed int    = 1
)

type knot struct {
	pos     v2
	follows *knot
}

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
	knots := []*knot{
		{pos: v2{0, 0}}, {pos: v2{0, 0}}, {pos: v2{0, 0}}, {pos: v2{0, 0}}, {pos: v2{0, 0}},
		{pos: v2{0, 0}}, {pos: v2{0, 0}}, {pos: v2{0, 0}}, {pos: v2{0, 0}}, {pos: v2{0, 0}},
	}
	for i, k := range knots {
		if i == 0 {
			continue
		}
		k.follows = knots[i-1]
	}
	h := knots[0]
	t := knots[len(knots)-1]

	for _, ln := range lines {
		parts := strings.Split(ln, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		h = knots[0]

		for i := 0; i < steps; i++ {
			if _, ok := m[t.pos.x]; !ok {
				m[t.pos.x] = map[int]struct{}{}
			}
			if _, ok := m[t.pos.x][t.pos.y]; !ok {
				m[t.pos.x][t.pos.y] = struct{}{}
			}

			h.pos = move(h.pos, dir)
			for _, k := range knots {
				if k.follows == nil {
					continue
				}
				separation := k.pos.orthoDistance(k.follows.pos)
				distance := k.pos.distance(k.follows.pos)
				if distance >= 2 {
					k.pos = translate(k.pos, v2{-separation.x, -separation.y}, knotMovespeed)
				}
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
