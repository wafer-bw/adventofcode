// https://adventofcode.com/2022/day/2#part2

package main

import (
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const puzzleID string = "2022-02"

var (
	shapeNames = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
	}

	shapeScores = map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors
	}

	directives = map[string]int{
		"X": 1, // lose
		"Y": 0, // draw
		"Z": 2, // win
	}
)

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func solve(lines []string) int {
	score := 0
	for _, ln := range lines {
		parts := strings.Split(ln, " ")
		p1 := parts[0]
		directive := parts[1]

		for p2 := range shapeNames {
			o, w := winner(p1, p2)
			if w == directives[directive] {
				log.Printf("%s %s vs %s", o, shapeNames[p1], shapeNames[p2])
				score += shapeScores[p2] + outcomeScore(w)
				break
			}
		}
	}

	return score
}

func winner(p1, p2 string) (string, int) {
	if shapeNames[p1] == shapeNames[p2] {
		return "draw", 0
	} else if p1 == "A" && p2 == "B" || p1 == "B" && p2 == "C" || p1 == "C" && p2 == "A" {
		return "win", 2
	}

	return "lose", 1
}

func outcomeScore(w int) int {
	if w == 0 {
		return 3
	} else if w == 2 {
		return 6
	}

	return 0
}
