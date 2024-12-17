package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/astar"
	"github.com/wafer-bw/adventofcode/tools/test"
	"github.com/wafer-bw/adventofcode/tools/vector"
)

//go:embed inputs.txt
var Inputs string

var Cases test.Cases = test.GetCases(Inputs)

const (
	RotateCost int = 1000
	MoveCost   int = 1
)

type Map struct {
	Region      map[vector.V2]map[vector.V2]Tile // Pos->Rotation->Tile
	Maximums    vector.V2
	End         vector.V2
	VisitedMode bool
	Visited     map[vector.V2]map[vector.V2]struct{}
}

func (m Map) String() string {
	s := "\n"
	for y := 0; y <= m.Maximums.Y; y++ {
		for x := 0; x <= m.Maximums.X; x++ {
			if t, ok := m.Region[vector.V2{X: x, Y: y}][vector.Cardinal2East]; ok {
				s += t.String()
			}
		}
		s += "\n"
	}
	return s
}

type Tile struct {
	Type TileType
	Map  *Map
	Pos  vector.V2 // position
	Dir  vector.V2 // direction // N,S,E,W.
}

func (t Tile) String() string {
	if t.Type == TileTypePath {
		return t.Dir.ToDirSymbol()
	}
	return string(t.Type)
}

func (t Tile) Occupiable() bool {
	return t.Type != TileTypeWall && t.Type != TileTypeTempWall
}

func (t Tile) Neighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	for _, dir := range vector.OrthoAdjacent2 {
		if neighbor, ok := t.Map.Region[t.Pos.Add(dir)][dir]; ok && neighbor.Occupiable() {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (t Tile) ExtraCost(to, facing vector.V2, c float64) float64 {
	if !t.Map.VisitedMode {
		return c
	}

	_, ok1 := t.Map.Visited[to]
	_, ok2 := t.Map.Visited[to][facing]
	if ok1 && ok2 {
		return c + 1
	}

	return c
}

func (t Tile) NeighborCost(to astar.Pather) float64 {
	toT := to.(Tile)

	c := 1

	facing := t.Dir
	approach := toT.Pos.Sub(t.Pos)

	if facing == approach {
		return t.ExtraCost(toT.Pos, facing, float64(c))
	}

	facing = t.Dir
	for i := range 2 {
		facing := facing.Rot(vector.RotateDirLeft)
		if facing == approach {
			return t.ExtraCost(toT.Pos, facing, float64(c+((i+1)*RotateCost)))
		}
	}

	facing = t.Dir
	for i := range 2 {
		facing = facing.Rot(vector.RotateDirRight)
		if facing == approach {
			return t.ExtraCost(toT.Pos, facing, float64(c+((i+1)*RotateCost)))
		}
	}

	panic("no path neighbor cost found")
}

func (t Tile) EstimatedCost(to astar.Pather) float64 {
	toT := to.(Tile)
	return t.Pos.ManhattanDistance(toT.Pos)
}

type TileType string

const (
	TileTypeEmpty    TileType = "."
	TileTypeWall     TileType = "#"
	TileTypeEnd      TileType = "E"
	TileTypeStart    TileType = "S"
	TileTypePath     TileType = "*"
	TileTypeTempWall TileType = "&"
)

type Paths map[string]struct{}

// return new version of map and a bool if the path that was added is new
func (p Paths) Add(path []astar.Pather) (map[string]struct{}, bool) {
	s := ""
	for _, p := range path {
		t := p.(Tile)
		s += t.Pos.String()
	}
	if _, ok := p[s]; ok {
		return p, false
	}
	p[s] = struct{}{}
	return p, true
}

type Seats map[string]vector.V2

func (s Seats) Add(path []astar.Pather) map[string]vector.V2 {
	for _, p := range path {
		t := p.(Tile)
		s[t.Pos.String()] = t.Pos
	}
	return s
}

func CostToRotate(t, f vector.V2) int {
	facing := t
	approach := f
	for i := range 2 {
		facing := facing.Rot(vector.RotateDirLeft)
		if facing == approach {
			return (i + 1) * RotateCost
		}
	}

	facing = t
	for i := range 2 {
		facing = facing.Rot(vector.RotateDirRight)
		if facing == approach {
			return (i + 1) * RotateCost
		}
	}

	panic("no cost found")
}

func ScorePath(path []astar.Pather) int {
	s := -1
	for i, p := range path {
		s++
		if i == len(path)-1 {
			break
		}
		t := p.(Tile)
		next := path[i+1].(Tile)
		if t.Dir != next.Dir {
			s += CostToRotate(t.Dir, next.Dir)
		}
	}
	return s
}

func PathString(path []astar.Pather) string {
	s := ""
	for _, p := range path {
		t := p.(Tile)
		s += t.Pos.String()
	}
	return s
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	var start, end vector.V2
	m := Map{Region: map[vector.V2]map[vector.V2]Tile{}, Visited: map[vector.V2]map[vector.V2]struct{}{}}
	for y, line := range lines {
		if y == 0 {
			m.Maximums.X = (len(line) - 1)
			m.Maximums.Y = len(lines) - 1
		}
		for x, cell := range strings.Split(line, "") {
			p := vector.V2{X: x, Y: y}
			tt := TileType(cell)
			if tt == TileTypeStart {
				start = p
				tt = TileTypeEmpty
			} else if tt == TileTypeEnd {
				end = p
				tt = TileTypeEmpty
				m.End = p
			}
			for _, d := range vector.Adjacent2 {
				if _, ok := m.Region[p]; !ok {
					m.Region[p] = map[vector.V2]Tile{}
				}
				m.Region[p][d] = Tile{Type: tt, Map: &m, Pos: p, Dir: d}
			}
		}
	}

	// find best path

	mdistance := math.MaxInt64
	for _, dir := range vector.OrthoAdjacent2 {
		_, distance, found := astar.Path(m.Region[start][vector.Cardinal2East], m.Region[end][dir])
		if found {
			mdistance = min(mdistance, int(distance))
		}
	}

	// collect possible other alternate paths
	paths := Paths{}
	seats := Seats{}
	for i := 0; i < 10; i++ {
		for _, dir := range vector.OrthoAdjacent2 {
			m.VisitedMode = true
			path, distance, found := astar.Path(m.Region[start][vector.Cardinal2East], m.Region[end][dir])
			_ = distance
			// fmt.Println(mdistance, distance, ScorePath(path))
			// TODO: rescore path manually to see if its distance is correct.
			if found && ScorePath(path) == mdistance {
				seats = seats.Add(path)
				paths, _ = paths.Add(path)
				for _, p := range path {
					t := p.(Tile)
					if _, ok := m.Visited[t.Pos]; !ok {
						m.Visited[t.Pos] = map[vector.V2]struct{}{}
					}
					if _, ok := m.Visited[t.Pos][t.Dir]; !ok {
						m.Visited[t.Pos][t.Dir] = struct{}{}
					}
				}
			}
		}
	}

	for _, s := range seats {
		mp := m.Region[s][vector.Cardinal2East]
		mp.Type = TileType("O")
		m.Region[s][vector.Cardinal2East] = mp
	}
	fmt.Println(m.String())

	return len(seats)
}

// full input answer is 497 < x < 604

func main() {
	filter := []int{}
	for i, c := range Cases {
		if slices.Contains(filter, i) {
			continue
		}
		log.Printf("case=%d expect=%d got=%d", i, c.Answer, Solve(c.Input))
	}
}
