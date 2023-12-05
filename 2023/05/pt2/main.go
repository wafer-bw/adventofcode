package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/wafer-bw/always"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

var seedPattern = regexp.MustCompile(`\d+ \d+`)

type Seeds struct {
	start int
	end   int
}

type Range struct {
	min    int
	max    int
	mutate int
}

type Table []Range

type Tables []Table

func (t Table) translate(v int) int {
	for _, r := range t {
		if v >= r.min && v <= r.max {
			return v + r.mutate
		}
	}
	return v
}

func (ts Tables) translate(v int) int {
	for _, t := range ts {
		v = t.translate(v)
	}
	return v
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	seeds := []Seeds{}
	for _, match := range seedPattern.FindAllString(lines[0], -1) {
		parts := strings.Split(match, " ")
		start := always.Accept(strconv.Atoi(parts[0]))
		end := start + always.Accept(strconv.Atoi(parts[1]))
		seeds = append(seeds, Seeds{start: start, end: end})
	}

	tables := Tables{}
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		} else if strings.Contains(line, ":") {
			tables = append(tables, Table{})
			continue
		}

		parts := strings.Split(line, " ")
		dst := always.Accept(strconv.Atoi(parts[0]))
		src := always.Accept(strconv.Atoi(parts[1]))
		length := always.Accept(strconv.Atoi(parts[2]))

		tables[len(tables)-1] = append(tables[len(tables)-1], Range{
			min:    src,
			max:    src + length - 1,
			mutate: (src - dst) * -1,
		})
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	v := math.MaxInt64
	batch := 10000000

	for _, s := range seeds {
		for i := s.start; i <= s.end; i += batch {
			end := i + batch
			if end > s.end {
				end = s.end
			}

			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				localMin := math.MaxInt64
				for j := start; j <= end; j++ {
					if n := tables.translate(j); n < localMin {
						localMin = n
					}
				}

				mu.Lock()
				if localMin < v {
					fmt.Println(localMin)
					v = localMin
				}
				mu.Unlock()
			}(i, end)
		}
	}

	wg.Wait()

	return v
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
