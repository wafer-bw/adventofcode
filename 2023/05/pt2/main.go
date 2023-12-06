package main

import (
	_ "embed"
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
		end := start + always.Accept(strconv.Atoi(parts[1])) - 1
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

	log.Println(names[Nature(0)])
	for _, s := range seeds {
		log.Printf("%d - %d", s.min, s.max)
	}

	for nat, c := range table {
		log.Println(names[Nature(nat+1)])
		for _, r := range c {
			log.Printf("\t%d - %d", r.min, r.max)
		}
	}

	translations := seeds
	v := math.MaxInt64
	for nat, c := range table {
		log.Println(names[Nature(nat+1)])
		for _, r := range c {
			newTranslations := []Range{}
			for _, t := range translations {
				if t.max < r.min || t.min > r.max {
					// fully outside range
					newTranslations = append(newTranslations,
						Range{min: t.min, max: t.max},
					)
				} else if t.min >= r.min && t.max <= r.max {
					// fully inside range
					newTranslations = append(newTranslations, Range{
						min: t.min + r.mut,
						max: t.max + r.mut,
					})
				} else if t.min < r.min && t.max <= r.max {
					// partially inside range (left)
					newTranslations = append(newTranslations,
						Range{min: t.min, max: r.min - 1},             // outside
						Range{min: r.min + r.mut, max: t.max + r.mut}, // inside
					)
				} else {
					// partially inside range (right)
					newTranslations = append(newTranslations,
						Range{min: t.min + r.mut, max: r.max + r.mut}, // inside
						Range{min: r.max + 1, max: t.max},             // outside
					)
				}
			}
			translations = newTranslations
		}
		for _, s := range translations {
			log.Printf("%d - %d", s.min, s.max)
		}
	}

	return v
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	// log.Printf("full: %d", Solve(FullInput))
}
