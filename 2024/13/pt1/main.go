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

// y = (8400 - 94x) / 22
func (m Machine) CalcL1Y(x int) float64 {
	return (float64(m.Prize.X) - float64(m.ButtonA.X)*float64(x)) / float64(m.ButtonB.X)
}

// y = (5400 - 34x) / 67
func (m Machine) CalcL2Y(x int) float64 {
	return (float64(m.Prize.Y) - float64(m.ButtonA.Y)*float64(x)) / float64(m.ButtonB.Y)
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
		for x := 0; x <= MaxPlays; x++ {
			y1 := machine.CalcL1Y(x)
			y2 := machine.CalcL2Y(x)
			if y1 == y2 {
				s += int(x)*ACost + int(y1)*BCost
			}
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
