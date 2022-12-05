package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"github.com/wafer-bw/adventofcode/tools/stack"
)

const (
	puzzleID string = "2022-05"
)

func solve(lines []string) string {
	numStacks := (len(lines[0]) + 1) / 4
	stacks, instructions := parse(lines, numStacks)
	workedStacks := work(stacks, instructions)
	for i, stk := range workedStacks {
		log.Printf("Stack %d: %v", i, stk)
	}
	return getResult(workedStacks)
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func getResult(stacks [][]string) string {
	result := ""
	for _, stk := range stacks {
		result += stk[len(stk)-1]
	}
	return result
}

func work(stacks [][]string, instructions [][]string) [][]string {
	for _, instruction := range instructions {
		numMoves, _ := strconv.Atoi(instruction[0])
		from, _ := strconv.Atoi(instruction[1])
		to, _ := strconv.Atoi(instruction[2])
		for i := 0; i < numMoves; i++ {
			stack.Push(&stacks[to-1], stack.Pop(&stacks[from-1]))
		}
	}

	return stacks
}

func parse(lines []string, numStacks int) (stacks [][]string, instructions [][]string) {
	// get stacks
	instructionStart := 0
	stacks = make([][]string, numStacks)
	for i, ln := range lines {
		stackId := 0
		chars := strings.Split(ln, "")
		for j := 0; j < len(chars); j++ {
			if j == 0 {
				continue
			}

			r := chars[j]
			if j == 1 || (j-1)%4 == 0 {
				if r != " " {
					stacks[stackId] = append(stacks[stackId], r)
				}
				stackId++
			}
		}

		if ln == "" {
			instructionStart = i + 1
			break
		}
	}

	// invert stacks
	for stk := range stacks {
		stacks[stk] = reverse(stacks[stk])
	}

	// get instructions
	instructions = [][]string{}
	for i, ln := range lines {
		if i < instructionStart {
			continue
		}

		parts := strings.Split(ln, " ")
		instruction := []string{parts[1], parts[3], parts[5]}
		instructions = append(instructions, instruction)
	}

	return
}

// reverse a slice
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
