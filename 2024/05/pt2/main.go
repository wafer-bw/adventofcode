package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type Rules map[int]map[int]struct{}

func (rules Rules) Compare(a, b int) int {
	behinds := rules[b]
	if _, ok := behinds[a]; ok {
		return 0
	}
	return -1
}

func Solve(input string) int {
	s := 0

	parts := strings.Split(input, "\n\n")
	rulesLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	updatesLines := strings.Split(strings.TrimSpace(parts[1]), "\n")

	rulesMap := Rules{}
	for _, rule := range rulesLines {
		parts := strings.Split(strings.TrimSpace(rule), "|")
		subject, _ := strconv.Atoi(parts[0])
		before, _ := strconv.Atoi(parts[1])
		if _, ok := rulesMap[subject]; !ok {
			rulesMap[subject] = map[int]struct{}{}
		}
		rulesMap[subject][before] = struct{}{}
	}

	updates := [][]int{}
	for _, updateStr := range updatesLines {
		parts := strings.Split(strings.TrimSpace(updateStr), ",")
		update := []int{}
		for _, page := range parts {
			number, _ := strconv.Atoi(page)
			update = append(update, number)
		}
		updates = append(updates, update)
	}

	for _, update := range updates {
		orig := make([]int, len(update))
		copy(orig, update)
		slices.SortStableFunc(update, func(a, b int) int {
			return rulesMap.Compare(a, b)
		})
		if fmt.Sprint(orig) != fmt.Sprint(update) {
			s += update[len(update)/2]
		}
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
