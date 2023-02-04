// https://adventofcode.com/2022/day/12

package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/wafer-bw/adventofcode/tools/alphanum"
	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"github.com/wafer-bw/adventofcode/tools/vector"
	"golang.org/x/exp/slices"
)

const (
	puzzleID string = "2022-12"
	maxClimb        = 1
)

var (
	up         vector.V2   = vector.V2{X: 0, Y: -1}
	down       vector.V2   = vector.V2{X: 0, Y: 1}
	left       vector.V2   = vector.V2{X: -1, Y: 0}
	right      vector.V2   = vector.V2{X: 1, Y: 0}
	directions []vector.V2 = []vector.V2{up, down, left, right}
)

type HeightMap [][]int

func (h HeightMap) HeightAt(pos vector.V2) int {
	return h[pos.Y][pos.X]
}

func (h HeightMap) OutOfBounds(pos vector.V2) bool {
	return pos.X < 0 || pos.Y < 0 || pos.X >= len(h[0])-1 || pos.Y >= len(h)-1
}

func (h HeightMap) TooHighToClimb(from, to vector.V2, max int) bool {
	return h.HeightAt(to)-h.HeightAt(from) > max
}

func (h HeightMap) String() string {
	msg := ""
	for _, x := range h {
		for _, y := range x {
			msg += string(alphanum.ToChar(y))
		}
		msg += "\n"
	}
	return msg
}

func solve(lines []string) int {
	// log.SetOutput(io.Discard)
	start, end, heightMap := setup(lines)
	return brute(start, end, heightMap)
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func setup(lines []string) (start, end vector.V2, heightMap HeightMap) {
	for i, ln := range lines {
		heightMap = append(heightMap, []int{})
		for j, ch := range ln {
			r := ch
			if ch == 'S' {
				r = 'a'
				start = vector.V2{X: j, Y: i}
			} else if ch == 'E' {
				r = 'z'
				end = vector.V2{X: j, Y: i}
			}

			n := alphanum.ToNum(r)
			heightMap[i] = append(heightMap[i], n)
		}
	}

	return
}

func heightAt(pos vector.V2, heightMap HeightMap) int {
	return heightMap[pos.Y][pos.X]
}

func brute(start, end vector.V2, heightMap HeightMap) int {
	pos := start
	maxMoves := len(heightMap) * len(heightMap[0])
	moves, visited := []vector.V2{}, []vector.V2{start}

	for pos != end {
		potentialMoves := getPotentialMoves(pos, heightMap, visited)
		selectedMove, err := decideMove(potentialMoves)
		if err != nil {
			log.Printf("hit dead end at %s", pos)
			// TODO: backtrack
			return -1
		}

		oldPos := pos
		pos = pos.Add(selectedMove)
		visited = append(visited, pos)
		moves = append(moves, selectedMove)

		log.Printf("move #%d: %s from %s to %s", len(moves), selectedMove.ToDir(), oldPos, pos)
		log.Println(visited)
		moveDirs := []string{}
		for _, m := range moves {
			moveDirs = append(moveDirs, m.ToDir())
		}
		log.Println(moveDirs)
		log.Println()

	}

	if len(moves) > maxMoves {
		return -1
	}

	return len(moves)
}

func getPotentialMoves(pos vector.V2, h HeightMap, visited []vector.V2) []vector.V2 {
	potentialMoves := []vector.V2{}
	for _, dir := range directions {
		target := pos.Add(dir)
		if h.OutOfBounds(target) {
			log.Printf("%s is out of bounds", dir.ToDir())
			continue
		} else if h.TooHighToClimb(pos, target, maxClimb) {
			log.Printf("%s is too high (%d)", dir.ToDir(), h.HeightAt(pos)-h.HeightAt(target))
			continue
		} else if alreadyVisited(target, visited) {
			log.Printf("%s has already been visited", target)
			continue
		}

		potentialMoves = append(potentialMoves, dir)
	}

	return potentialMoves
}

func decideMove(potentialMoves []vector.V2) (vector.V2, error) {
	if len(potentialMoves) == 0 {
		return vector.V2{}, fmt.Errorf("no potential moves")
	} else if len(potentialMoves) == 1 {
		return potentialMoves[0], nil
	}
	return potentialMoves[rand.Intn(len(potentialMoves)-1)], nil
}

func alreadyVisited(pos vector.V2, visited []vector.V2) bool {
	return slices.Contains(visited, pos)
}
