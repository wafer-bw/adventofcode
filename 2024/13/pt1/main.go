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
	ACost    int = 3
	BCost    int = 1
	MaxPlays int = 100
)

type Machine struct {
	ButtonA vector.V2
	ButtonB vector.V2
	Prize   vector.V2
}

func (m Machine) CalcL1X(y int) float64 {
	return (float64(m.Prize.X) - float64(m.ButtonB.X)*float64(y)) / float64(m.ButtonA.X)
}
func (m Machine) CalcL2X(y int) float64 {
	return (float64(m.Prize.Y) - float64(m.ButtonB.Y)*float64(y)) / float64(m.ButtonA.Y)
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
			machines[len(machines)-1].Prize = vector.V2{X: x, Y: y}
		}
	}

	for _, machine := range machines {
		for y := 0; y <= MaxPlays; y++ {
			x1 := machine.CalcL1X(y)
			x2 := machine.CalcL2X(y)
			if x1 == x2 {
				s += int(x1)*ACost + y*BCost
			}
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
