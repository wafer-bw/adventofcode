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
	timeParts = timeParts[1:]
	numStr := ""
	for _, part := range timeParts {
		part = strings.TrimSpace(part)
		numStr += part
	}
	raceTimes = append(raceTimes, always.Accept(strconv.Atoi(numStr)))

	distParts := strings.Split(lines[1], " ")
	raceDistances := []int{}
	distParts = distParts[1:]
	distStr := ""
	for _, part := range distParts {
		part = strings.TrimSpace(part)
		distStr += part
	}
	raceDistances = append(raceDistances, always.Accept(strconv.Atoi(distStr)))

	races := make([]race, len(raceTimes))
	for i := 0; i < len(races); i++ {
		races[i] = race{
			time:     raceTimes[i],
			distance: raceDistances[i],
		}
	}

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
