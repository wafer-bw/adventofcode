package main

import (
	_ "embed"
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

var multPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func Solve(input string) int {
	s := 0

	data := strings.Join(strings.Split(input, "\n"), " ")
	matches := multPattern.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		if len(match) == 3 {
			a := match[1]
			b := match[2]
			aInt, _ := strconv.Atoi(a)
			bInt, _ := strconv.Atoi(b)
			s += aInt * bInt
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
