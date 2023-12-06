package main

import (
	_ "embed"
	"fmt"
	"log"
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

type race struct {
	time          int
	distance      int
	winningStrats []raceStrategy
}

func (r race) String() string {
	return fmt.Sprintf("t: %d,  d: %d", r.time, r.distance)
}

type raceStrategy struct {
	race           *race
	buttonDuration int
	runDistance    int
}

func (rs raceStrategy) String() string {
	return fmt.Sprintf("rt: %d, rd: %d, bt: %d,  dd: %d", rs.race.time, rs.race.distance, rs.buttonDuration, rs.runDistance)
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	timeParts := strings.Split(lines[0], " ")
	raceTimes := []int{}
	for i, part := range timeParts {
		if i == 0 {
			continue
		}
		if strings.TrimSpace(part) == "" {
			continue
		}
		raceTimes = append(raceTimes, always.Accept(strconv.Atoi(part)))
	}

	distParts := strings.Split(lines[1], " ")
	raceDistances := []int{}
	for i, part := range distParts {
		if i == 0 {
			continue
		}
		if strings.TrimSpace(part) == "" {
			continue
		}
		raceDistances = append(raceDistances, always.Accept(strconv.Atoi(part)))
	}

	races := make([]race, len(raceTimes))
	for i := 0; i < len(races); i++ {
		races[i] = race{
			time:     raceTimes[i],
			distance: raceDistances[i],
		}
	}

	// 7s, 9m
	// --
	// t -> t-bt -> (t-bt)*t -> v
	// 0 -> 7-0 -> 0 -> 0
	// 1 -> 7-1 -> 6*1 -> 6
	// 2 -> 7-2 -> 5*2 -> 10
	// 3 -> 7-3 -> 4*3 -> 12
	// 4 -> 7-4 -> 3*4 -> 12
	// 5 -> 7-5 -> 2*5 -> 10
	// 6 -> 7-6 -> 1*6 -> 6
	// 7 -> 7-7 -> 0 -> 0

	for i, race := range races {
		for buttonTime := 1; buttonTime < race.time; buttonTime++ {
			outcome := (race.time - buttonTime) * buttonTime
			if outcome <= race.distance {
				continue
			}
			strat := raceStrategy{
				race:           &race,
				buttonDuration: buttonTime,
				runDistance:    outcome,
			}
			races[i].winningStrats = append(races[i].winningStrats, strat)
		}
	}

	v := 1
	for _, race := range races {
		fmt.Println(len(race.winningStrats))
		v *= len(race.winningStrats)
	}

	return v
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
