package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

const (
	ACost int = 3
	BCost int = 1
	Add   int = 10000000000000
)

type Machine struct {
	ButtonA vector.V2
	ButtonB vector.V2
	Prize   vector.V2
}

// Find intercepts using Cramer's Rule.
// https://stackoverflow.com/a/39395330
func (m Machine) Intercepts() (int, int, bool) {
	aX, aY := m.ButtonA.X, m.ButtonA.Y
	bX, bY := m.ButtonB.X, m.ButtonB.Y
	pX, pY := m.Prize.X, m.Prize.Y

	d, dX, dY := aX*bY-bX*aY, pX*bY-bX*pY, aX*pY-pX*aY
	if d == 0 || dX != (dX/d)*d || dY != (dY/d)*d {
		return 0, 0, false
	}

	return dX / d, dY / d, true
}

func Solve(input string) int {
	s := 0

	machines := []Machine{}
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if i%4 == 0 {
			machines = append(machines, Machine{})
			line = strings.TrimPrefix(line, "Button A: X+")
			parts := strings.Split(line, ", ")
			parts[1] = strings.TrimPrefix(parts[1], "Y+")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			machines[len(machines)-1].ButtonA = vector.V2{X: x, Y: y}
		} else if i%4 == 1 {
			line = strings.TrimPrefix(line, "Button B: X+")
			parts := strings.Split(line, ", ")
			parts[1] = strings.TrimPrefix(parts[1], "Y+")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			machines[len(machines)-1].ButtonB = vector.V2{X: x, Y: y}
		} else if i%4 == 2 {
			line = strings.TrimPrefix(line, "Prize: X=")
			parts := strings.Split(line, ", ")
			parts[1] = strings.TrimPrefix(parts[1], "Y=")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			machines[len(machines)-1].Prize = vector.V2{X: Add + x, Y: Add + y}
		}
	}

	for _, m := range machines {
		x, y, ok := m.Intercepts()
		if ok {
			s += x*ACost + y*BCost
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
