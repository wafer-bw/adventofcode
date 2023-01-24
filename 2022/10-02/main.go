// https://adventofcode.com/2022/day/10#part2

package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID string = "2022-10"
)

type cmd string

const (
	noopCmd  cmd  = "noop"
	addxCmd  cmd  = "addx"
	litPixel rune = '#'
)

type operation struct {
	index           int
	cyclesRemaining int
	cmd             cmd
	val             int
}

func solve(lines []string) string {
	var (
		op       *operation = nil
		register int        = 1
		cycle    int        = 0
		line     int        = 0
		screen   []string   = []string{
			"........................................",
			"........................................",
			"........................................",
			"........................................",
			"........................................",
			"........................................",
		}
		y int = 0
	)

	for i := 0; i < len(lines)*2; i++ {
		cycle++
		x := (cycle - 1) % 40

		for j := -1; j < 2; j++ {
			idx := register + j
			if x == idx && idx >= 0 && idx <= 39 {
				screen[y] = replaceStringAtIndex(screen[y], litPixel, idx)
				// log.Printf("cycle: %d, register: %d, line %d, x: %d, y: %d", cycle%40, register, line+1, idx, y)
				// fmt.Printf(strings.Join(screen, "\n") + "\n\n")
				break
			}
		}

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

		if cycle%40 == 0 {
			y++
			if y > 5 {
				return strings.Join(screen, "\n")
				// y = 0
			}
		}
	}

	return strings.Join(screen, "\n")
}

func replaceStringAtIndex(s string, r rune, i int) string {
	runes := []rune(s)
	runes[i] = r
	return string(runes)
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
