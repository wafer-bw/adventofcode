package main

import (
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	input       string = "inputs/2022-03.txt"
	upperOffset int    = 38
	lowerOffset int    = 96
)

func priority(c rune) int {
	if strings.ToLower(string(c)) == string(c) {
		return int(c) - lowerOffset
	}

	return int(c) - upperOffset
}

func shared(a, b string) []rune {
	shared := []rune{}

	for _, c := range a {
		if strings.Contains(b, string(c)) {
			shared = append(shared, c)
		}
	}

	return shared
}

func solve(lines []string) int {
	sum := 0
	for _, ln := range lines {
		p1, p2 := ln[:len(ln)/2], ln[len(ln)/2:]
		sh := shared(p1, p2)
		sum += priority(sh[0])
	}

	return sum
}

func main() {
	log.Println(solve(reader.Read(input)))
}
