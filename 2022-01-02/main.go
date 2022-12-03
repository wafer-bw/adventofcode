package main

import (
	"log"
	"strconv"

	"github.com/wafer-bw/adventofcode/tools/reader"
	"golang.org/x/exp/slices"
)

const input string = "inputs/2022-01.txt"

func solve(lines []string) int {
	i := 1
	elves := []int{}

	for _, ln := range lines {
		c, err := strconv.Atoi(ln)
		if err != nil {
			i++
			continue
		}

		if len(elves) < i {
			elves = append(elves, 0)
		}
		elves[i-1] += c
	}

	slices.Sort(elves)
	topThree := elves[len(elves)-3:]

	sum := 0
	for _, c := range topThree {
		log.Println(c)
		sum += c
	}

	return sum
}

func main() {
	log.Println(solve(reader.Read(input)))
}
