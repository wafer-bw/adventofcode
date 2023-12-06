package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/wafer-bw/always"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

var seedPattern = regexp.MustCompile(`\d+ \d+`)

type Range struct {
	min int
	max int
	mut int
}

type Ranges []Range

type Table []Ranges

type Nature int

const (
	NatureSeed Nature = iota
	NatureSoil
	NatureFertilizer
	NatureWater
	NatureLight
	NatureTemp
	NatureHumidity
	NatureLocation
)

var names = map[Nature]string{
	NatureSeed:       "seed",
	NatureSoil:       "soil",
	NatureFertilizer: "fertilizer",
	NatureWater:      "water",
	NatureLight:      "light",
	NatureTemp:       "temp",
	NatureHumidity:   "humidity",
	NatureLocation:   "location",
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	seeds := []Range{}
	for _, match := range seedPattern.FindAllString(lines[0], -1) {
		parts := strings.Split(match, " ")
		start := always.Accept(strconv.Atoi(parts[0]))
		end := start + always.Accept(strconv.Atoi(parts[1]))
		seeds = append(seeds, Range{min: start, max: end})
	}

	table := Table{}
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		} else if strings.Contains(line, ":") {
			table = append(table, Ranges{})
			continue
		}

		parts := strings.Split(line, " ")
		dst := always.Accept(strconv.Atoi(parts[0]))
		src := always.Accept(strconv.Atoi(parts[1]))
		length := always.Accept(strconv.Atoi(parts[2]))

		table[len(table)-1] = append(table[len(table)-1], Range{
			min: src,
			max: src + length - 1,
			mut: (src - dst) * -1,
		})
	}

	translation := seeds
	v := math.MaxInt64
	for _, c := range table {
		for _, r := range c {
			removeTranslation := []int{}
			addTranslation := []Range{}
			for ti, t := range translation { // need to make the order of loop nesting correct.
				log.Println(names[Nature(ti)])
				for _, s := range translation {
					log.Printf("%d - %d", s.min, s.max)
				}
				fmt.Println()
				if t.max < r.min || t.min > r.max {
					// fully outside range
					continue
				} else if t.min >= r.min && t.max <= r.max {
					// fully inside range
					translation[ti].min += r.mut
					translation[ti].max += r.mut
				} else if t.min < r.min && t.max <= r.max {
					// partially inside range (left)
					removeTranslation = append(removeTranslation, ti)
					addTranslation = append(addTranslation,
						Range{min: t.min, max: r.min - 1},             // outside
						Range{min: r.min + r.mut, max: t.max + r.mut}, // inside
					)
				} else if t.min >= r.min && t.max > r.max {
					// partially inside range (right)
					removeTranslation = append(removeTranslation, ti)
					addTranslation = append(addTranslation,
						Range{min: t.min + r.mut, max: r.max + r.mut}, // inside
						Range{min: r.max + 1, max: t.max},             // outside
					)
				}
			}
			translation = remove(translation, removeTranslation...)
			translation = append(translation, addTranslation...)
			for _, s := range translation {
				log.Printf("%d - %d", s.min, s.max)
			}
		}
	}

	for _, s := range translation {
		log.Printf("%d - %d", s.min, s.max)
	}

	return v
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	// log.Printf("full: %d", Solve(FullInput))
}

func remove[T any](slice []T, ids ...int) []T {
	for i, idx := range ids {
		slice = append(slice[:idx-i], slice[idx-i+1:]...)
	}
	return slice
}
