package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

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

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}

func Solve(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		inc, _ := strconv.Atoi(fmt.Sprintf("%d%d", getDigit(line), getRevDigit(line)))
		sum += inc

	}

	return sum
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
