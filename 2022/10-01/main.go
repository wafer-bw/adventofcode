// https://adventofcode.com/2022/day/10

package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"golang.org/x/exp/slices"
)

const (
	puzzleID string = "2022-10"
)

type cmd string

const (
	noopCmd cmd = "noop"
	addxCmd cmd = "addx"
)

type operation struct {
	index           int
	cyclesRemaining int
	cmd             cmd
	val             int
}

func solve(lines []string) int {
	var (
		op       *operation = nil
		register int        = 1
		cycle    int        = 0
		line     int        = 0
		sum      int        = 0
		poi      []int      = []int{20, 60, 100, 140, 180, 220}
	)

	for i := 0; i < len(lines)*2; i++ {
		cycle++

		signal := cycle * register
		if slices.Contains(poi, cycle) {
			sum += signal
		}
		log.Printf("cycle: %d, register: %d, signal: %d, sum: %d, line %d", cycle, register, signal, sum, line+1)

		if line >= len(lines) {
			break
		} else if op != nil {
			op.cyclesRemaining--
			if op.cyclesRemaining == 0 {
				register += op.val
				op = nil
				line++
			}
		} else {
			cmd, val := getCmdVal(lines[line])
			switch cmd {
			case noopCmd:
				line++
			case addxCmd:
				op = &operation{
					index:           line,
					cyclesRemaining: 1,
					cmd:             cmd,
					val:             val,
				}
			default:
				log.Fatal("unknown command", cmd)
			}
		}
	}

	return sum
}

func getCmdVal(line string) (cmd, int) {
	parts := strings.Split(line, " ")
	cmd := cmd(parts[0])
	val := 0
	if len(parts) > 1 {
		val, _ = strconv.Atoi(parts[1])
	}
	return cmd, val
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
