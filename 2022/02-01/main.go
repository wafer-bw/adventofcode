// https://adventofcode.com/2022/day/2

package main

import (
	"fmt"
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
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	shapeScores = map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissors
	}
)

func solve(lines []string) int {
	score := 0
	for _, ln := range lines {
		parts := strings.Split(ln, " ")
		p1 := parts[0]
		p2 := parts[1]

		score += shapeScores[p2]
		w := winner(p1, p2)
		if w == 0 {
			score += 3
		} else if w == 2 {
			score += 6
		}
	}

	return score
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func winner(p1, p2 string) int {
	v := fmt.Sprintf("%s vs %s", shapeNames[p1], shapeNames[p2])

	if shapeNames[p1] == shapeNames[p2] {
		log.Printf("draw: %s", v)
		return 0
	}

	if p1 == "A" && p2 == "Y" || p1 == "B" && p2 == "Z" || p1 == "C" && p2 == "X" {
		log.Printf("p2 wins: %s", v)
		return 2
	}

	log.Printf("p1 wins: %s", v)
	return 1
}
