// https://adventofcode.com/2022/day/1

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const puzzleID string = "2023-01"

func solve(lines []string) int {
	s := 0

	for _, line := range lines {
		numbers := []int{}
		for _, ch := range line {
			if n, err := strconv.Atoi(string(ch)); err == nil {
				numbers = append(numbers, n)
			}
		}
		value := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
		numericalValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		s += numericalValue
	}

	return s
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
