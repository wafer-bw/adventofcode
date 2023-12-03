package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

var numberPattern = regexp.MustCompile(`\d+`)

type number struct {
	value  int
	row    int
	start  int
	finish int
}

type gear []int

func Solve(input string) int {
	s := 0
	lines := strings.Split(input, "\n")

	chars := make([][]string, len(lines))
	for i, line := range lines {
		chars[i] = make([]string, len(line))
		for j, char := range line {
			chars[i][j] = string(char)
		}
	}

	numbers := []number{}
	for row, line := range lines {
		indices := numberPattern.FindAllStringIndex(line, -1)
		for _, index := range indices {
			n, _ := strconv.Atoi(line[index[0]:index[1]])
			numbers = append(numbers, number{
				value:  n,
				row:    row,
				start:  index[0],
				finish: index[1],
			})
		}
	}

	gears := map[string]gear{}
	for _, n := range numbers {
	Number:
		for row := n.row - 1; row <= n.row+1; row++ {
			if row < 0 || row > len(lines)-1 {
				continue
			}

			for col := n.start - 1; col <= n.finish; col++ {
				if col < 0 || col > len(chars[row])-1 {
					continue
				}

				if chars[row][col] == "*" {
					gears[fmt.Sprintf("%d-%d", row, col)] = append(gears[fmt.Sprintf("%d-%d", row, col)], n.value)
					continue Number
				}
			}
		}
	}

	for _, g := range gears {
		if len(g) == 2 {
			s += g[0] * g[1]
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
