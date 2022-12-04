package main

import (
	"log"
	"strconv"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const puzzleID string = "2022-01"

func solve(lines []string) int {
	most := 0
	current := 0

	for _, ln := range lines {
		if ln != "" {
			c, err := strconv.Atoi(ln)
			if err != nil {
				log.Fatal(err)
			}
			current += c
		}

		if current > most {
			most = current
		}

		if ln == "" {
			current = 0
		}
	}

	return most
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}
