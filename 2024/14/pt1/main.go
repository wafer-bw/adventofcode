package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

const seconds int = 100

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type Map struct {
	MaxX        int
	MaxY        int
	RobotCounts map[vector.V2]int
	Robots      []Robot
	Steps       int
}

func (m *Map) Step() {
	for i := range m.Robots {
		from := m.Robots[i].P
		to := m.Robots[i].P.Add(m.Robots[i].V)
		to.X = (to.X + m.MaxX) % m.MaxX
		to.Y = (to.Y + m.MaxY) % m.MaxY

		m.RobotCounts[from] -= 1
		m.Robots[i].P = to
		m.RobotCounts[to] += 1
	}
	m.Steps += 1
}

func (m Map) QuadCount() (int, int, int, int) {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for p, c := range m.RobotCounts {
		if c == 0 {
			continue
		}
		middleX := m.MaxX / 2
		middleY := m.MaxY / 2

		if p.X < middleX && p.Y < middleY {
			q1 += c
		} else if p.X > middleX && p.Y < middleY {
			q2 += c
		} else if p.X < middleX && p.Y > middleY {
			q3 += c
		} else if p.X > middleX && p.Y > middleY {
			q4 += c
		}
	}
	return q1, q2, q3, q4
}

func (m *Map) String() string {
	s := ""
	for y := 0; y < m.MaxY; y++ {
		for x := 0; x < m.MaxX; x++ {
			if r, ok := m.RobotCounts[vector.V2{X: x, Y: y}]; ok && r > 0 {
				s += fmt.Sprintf("%d", r)
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

type Robot struct {
	P vector.V2 // position
	V vector.V2 // velocity
}

func Solve(input string) int {
	m := Map{MaxX: 0, MaxY: 0}
	lines := strings.Split(input, "\n")
	robots := make([]Robot, len(lines)-1)
	for i, line := range lines {
		if i == 0 {
			_, _ = fmt.Sscanf(line, "%d,%d", &m.MaxX, &m.MaxY)
			continue
		}
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robots[i-1].P.X, &robots[i-1].P.Y, &robots[i-1].V.X, &robots[i-1].V.Y)
	}

	m.Robots = robots
	m.RobotCounts = make(map[vector.V2]int)
	for i := range m.Robots {
		m.RobotCounts[robots[i].P] += 1
	}

	for range seconds {
		m.Step()
	}

	q1, q2, q3, q4 := m.QuadCount()
	return q1 * q2 * q3 * q4
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
