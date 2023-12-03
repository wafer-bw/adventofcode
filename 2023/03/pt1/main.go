package main

import (
	_ "embed"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/wafer-bw/always"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

var (
	numberPattern = regexp.MustCompile(`\d+`)
	symbolPattern = regexp.MustCompile(`[^\d.a-zA-Z]`)
)

type number struct {
	value  int
	row    int
	start  int
	finish int
}

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
			numbers = append(numbers, number{
				value:  always.Accept(strconv.Atoi(line[index[0]:index[1]])),
				row:    row,
				start:  index[0],
				finish: index[1],
			})
		}
	}

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

				if symbolPattern.MatchString(chars[row][col]) {
					s += n.value
					break Number
				}
			}
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
