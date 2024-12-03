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
	Unsafe bool
	Mode   Mode
	Levels []int
}

type Mode string

const (
	ModeUnknown    Mode = ""
	ModeIncreasing Mode = "increasing"
	ModeDecreasing Mode = "decreasing"
)

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

	for i, report := range reports {
		for j, level := range report.Levels {
			if j != 0 && (mth.IntAbs(level-report.Levels[j-1]) < 1 || mth.IntAbs(level-report.Levels[j-1]) > 3) {
				reports[i].Unsafe = true
			}

			var observedMode Mode
			if j != 0 {
				if level < report.Levels[j-1] {
					observedMode = ModeDecreasing
				} else if level > report.Levels[j-1] {
					observedMode = ModeIncreasing
				} else {
					reports[i].Unsafe = true
					break
				}
			}

			if j == 1 {
				reports[i].Mode = observedMode
			} else if j > 0 {
				if reports[i].Mode != observedMode {
					reports[i].Unsafe = true
					break
				}
			}
		}
	}

	for _, report := range reports {
		if !report.Unsafe {
			s++
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	// log.Printf("full: %d", Solve(FullInput))
}
