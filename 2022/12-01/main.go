// https://adventofcode.com/2022/day/12

package main

import (
	"fmt"
	"log"

	"github.com/wafer-bw/adventofcode/tools/alphanum"
	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
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
	return pos.Y < 0 || pos.X < 0 || pos.Y >= len(h) || pos.X >= len(h[0])
}

func (h HeightMap) TooHighToClimb(from, to vector.V2, max int) bool {
	return h.HeightAt(to)-h.HeightAt(from) > max
}

func (h HeightMap) String() string {
	msg := ""
	for _, y := range h {
		for _, x := range y {
			msg += string(alphanum.ToChar(x))
		}
		msg += "\n"
	}
	return msg
}

func solve(lines []string) int {
	start, end, heightMap := setup(lines)
	return wander(start, end, heightMap)
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

func wander(start, end vector.V2, heightMap HeightMap) int {
	step, lowest := 0, -1
	paths := [][]vector.V2{{start}}

	fmt.Printf("map:\n%s\n", heightMap)
	fmt.Printf("%v to %v\n\n", start, end)

	for len(paths) > 0 {
		step++
		c := len(paths)

		newPaths := [][]vector.V2{}
		nexts := []vector.V2{}
		for id := 0; id < c; id++ {
			path := paths[id]
			last := path[len(path)-1]

			if last == end && (lowest == -1 || len(path) < lowest) {
				lowest = len(path)
				continue
			}
			potentialMoves := getPotentialMoves(last, heightMap, path)

			for _, move := range potentialMoves {
				next := last.Add(move)
				newPath := append([]vector.V2{}, append(path, next)...)
				lookAhead := getPotentialMoves(next, heightMap, newPath)
				if len(lookAhead) == 0 && next != end {
					continue
				} else if slices.Contains(nexts, next) {
					continue
				}
				nexts = append(nexts, next)
				newPaths = append(newPaths, newPath)
			}
		}
		paths = newPaths
		log.Printf("%d: currently have %d paths", step, len(paths))
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Println()
		// reader.ReadString('\n')
	}

	return lowest - 1
}

func getPotentialMoves(pos vector.V2, h HeightMap, visited []vector.V2) []vector.V2 {
	potentialMoves := []vector.V2{}
	for _, dir := range directions {
		target := pos.Add(dir)
		if h.OutOfBounds(target) {
			continue
		} else if h.TooHighToClimb(pos, target, maxClimb) {
			continue
		} else if alreadyVisited(target, visited) {
			continue
		}

		potentialMoves = append(potentialMoves, dir)
	}

	return potentialMoves
}

func alreadyVisited(pos vector.V2, visited []vector.V2) bool {
	return slices.Contains(visited, pos)
}
