// https://adventofcode.com/2022/day/12

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/wafer-bw/adventofcode/tools/alphanum"
	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	"github.com/wafer-bw/adventofcode/tools/stack"
	"github.com/wafer-bw/adventofcode/tools/vector"
	"golang.org/x/exp/slices"
)

const (
	puzzleID string = "2022-12"
	maxClimb int    = 1
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
	return pos.X < 0 || pos.Y < 0 || pos.X >= len(h[0]) || pos.Y >= len(h)
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
	log.Println(solve(reader.Read(pather.Path(puzzleID, true, false))))
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

func brute(start, end vector.V2, heightMap HeightMap) int {
	pos := start
	maxMoves := 32
	moves, visited := []vector.V2{}, []vector.V2{start}

	log.Println("map:\n", heightMap)
	log.Printf("%v to %v", start, end)

	for pos != end {
		potentialMoves := getPotentialMoves(pos, heightMap, visited)
		selectedMove, err := decideMove(potentialMoves)
		if err != nil || len(moves) >= maxMoves {
			log.Printf("hit dead end at %s", pos)
			for n := len(moves) - 1; n >= 0; n-- {
				oldPos := pos
				pos = pos.Sub(moves[n])
				log.Println()
				log.Printf("backtrack move #%d: %s from %s to %s", len(moves), moves[n].Neg().ToDir(), oldPos, pos)
				pot := getPotentialMoves(pos, heightMap, visited)
				stack.Pop(&visited)
				stack.Pop(&moves)
				if len(pot) == 0 {
					continue
				} else {
					selectedMove, _ = decideMove(pot)
					break
				}
			}
		}

		oldPos := pos
		pos = pos.Add(selectedMove)
		visited = append(visited, pos)
		moves = append(moves, selectedMove)

		log.Printf("move #%d: %s from %s to %s", len(moves), selectedMove.ToDir(), oldPos, pos)
		log.Println(visited)
		log.Println()

		moveDirs := []string{}
		for _, m := range moves {
			moveDirs = append(moveDirs, m.ToDir())
		}
		log.Println(moveDirs)
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("")
		reader.ReadString('\n')
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
