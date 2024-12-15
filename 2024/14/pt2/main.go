package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/mth"
	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
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

func (m *Map) Step(s int) {
	for i := range m.Robots {
		from := m.Robots[i].P
		vel := m.Robots[i].V.Mul(s)
		to := m.Robots[i].P.Add(vel)
		to = vector.V2{X: mth.PMod(to.X, m.MaxX), Y: mth.PMod(to.Y, m.MaxY)}
		m.RobotCounts[from] -= 1
		m.Robots[i].P = to
		m.RobotCounts[to] += 1
	}
	m.Steps += s
}

func (m Map) QuadCheck() (int, int, int, int) {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for p, c := range m.RobotCounts {
		if c == 0 {
			continue
		}
		middleX := m.MaxX / 2
		middleY := m.MaxY / 2

		if p.X < middleX && p.Y < middleY {
			q1 += 1
		} else if p.X > middleX && p.Y < middleY {
			q2 += 1
		} else if p.X < middleX && p.Y > middleY {
			q3 += 1
		} else if p.X > middleX && p.Y > middleY {
			q4 += 1
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

func Solve(input string, step ...int) string {
	grids := map[uint64]struct{}{}
	_ = grids
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

	if len(step) > 0 {
		m.Step(step[0])
		return m.String()
	}

	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	q1, q2, q3, q4 := 0, 0, 0, 0
	for {
		if m.Steps%1000 == 0 {
			fmt.Println(m.Steps)
		}
		m.Step(1)

		flag := false
		qq1, qq2, qq3, qq4 := m.QuadCheck()
		if qq1 > q1 {
			flag = true
			q1 = qq1
		} else if qq2 > q2 {
			q2 = qq2
			flag = true
		} else if qq3 > q3 {
			q3 = qq3
			flag = true
		} else if qq4 > q4 {
			q4 = qq4
			flag = true
		}

		if flag {
			if _, err = f.Write([]byte(fmt.Sprintf("%d\n", m.Steps))); err != nil {
				panic(err)
			}
			if _, err := f.Write([]byte(m.String())); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	Solve(FullInput)
}
