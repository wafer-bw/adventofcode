// https://adventofcode.com/2022/day/12#part2

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
	checked := 0
	starts := []vector.V2{}
	_, end, heightMap := setup(lines)

	fmt.Printf("map:\n%s\n", heightMap)

	for y := range heightMap {
		for x := range heightMap[y] {
			if heightMap[y][x] == 1 {
				starts = append(starts, vector.V2{X: x, Y: y})
			}
		}
	}

	least := -1
	for _, start := range starts {
		checked++
		n, err := wander(start, end, heightMap)
		if err == nil && (least == -1 || n < least) {
			least = n
		} else {
			log.Printf("%s for start %s", err, start)
			continue
		}

		log.Printf("current shortest path is %d", least)
		log.Printf("checked %d/%d paths (%.2f%%)", checked, len(starts), float64(checked)/float64(len(starts))*100)
	}

	return least
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

func wander(start, end vector.V2, heightMap HeightMap) (int, error) {
	step, lowest := 0, -1
	paths := [][]vector.V2{{start}}
	log.Printf("%v to %v\n", start, end)

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
	}

	if lowest-1 < 0 {
		return -1, fmt.Errorf("no path found")
	}
	return lowest - 1, nil
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
