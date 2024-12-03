package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/mth"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type Report struct {
	Levels    []int
	decreases []int
	increases []int
	jumps     []int
}

func (report *Report) Evaluate() {
	for i, level := range report.Levels {
		if i == 0 {
			continue
		}

		if mth.IntAbs(level-report.Levels[i-1]) < 1 || mth.IntAbs(level-report.Levels[i-1]) > 3 {
			report.jumps = append(report.jumps, i)
		}

		if level < report.Levels[i-1] {
			report.decreases = append(report.decreases, i)
		} else if level > report.Levels[i-1] {
			report.increases = append(report.increases, i)
		}
	}
}

func (report Report) Dampen(idx int) Report {
	newLevels := make([]int, 0, len(report.Levels)-1)
	for i, level := range report.Levels {
		if i != idx {
			newLevels = append(newLevels, level)
		}
	}
	return Report{Levels: newLevels}
}

func (report Report) Safe() bool {
	if len(report.jumps) > 0 {
		return false
	} else if len(report.decreases) > 0 && len(report.increases) > 0 {
		return false
	}

	return true
}

func Solve(input string) int {
	s := 0

	lines := strings.Split(input, "\n")
	reports := make([]Report, len(lines))
	for i, line := range lines {
		levelStrs := strings.Split(line, " ")
		for _, levelStr := range levelStrs {
			level, _ := strconv.Atoi(levelStr)
			reports[i].Levels = append(reports[i].Levels, level)
		}
	}

	for _, report := range reports {
		report.Evaluate()
		if report.Safe() {
			s++
			continue
		}
		for i := 0; i < len(report.Levels); i++ {
			r := report.Dampen(i)
			r.Evaluate()
			if r.Safe() {
				s++
				break
			}
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	// log.Printf("full: %d", Solve(FullInput))
}
