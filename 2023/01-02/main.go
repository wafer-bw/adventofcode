// https://adventofcode.com/2022/day/1

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const puzzleID string = "2023-01"

var words = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func solve(lines []string) int {
	sum := 0

	for _, line := range lines {
		inc, _ := strconv.Atoi(fmt.Sprintf("%d%d", getDigit(line), getRevDigit(line)))
		sum += inc

	}

	return sum
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func getDigit(s string) int {
	build := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if n, err := strconv.Atoi(string(ch)); err == nil {
			return n
		}
		build += string(ch)

		for word, digit := range words {
			if strings.Contains(build, word) {
				return digit
			}
		}
	}
	return 0
}

func getRevDigit(s string) int {
	build := ""
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if n, err := strconv.Atoi(string(ch)); err == nil {
			return n
		}
		build = string(ch) + build

		for word, digit := range words {
			if strings.Contains(build, word) {
				return digit
			}
		}
	}
	return 0
}
