package main

import (
	_ "embed"
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
	Region   map[vector.V2]map[vector.V2]Tile // Pos->Rotation->Tile
	Maximums vector.V2
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
	return t.Type != TileTypeWall
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

func (t Tile) NeighborCost(to astar.Pather) float64 {
	toT := to.(Tile)

	c := 1
	facing := t.Dir
	approach := toT.Pos.Sub(t.Pos)

	if facing == approach {
		return float64(c)
	}

	facing = t.Dir
	for i := range 2 {
		facing = facing.Rot(vector.RotateDirRight)
		if facing == approach {
			return float64(c + ((i + 1) * RotateCost))
		}
	}

	facing = t.Dir
	for i := range 2 {
		facing := facing.Rot(vector.RotateDirLeft)
		if facing == approach {
			return float64(c + ((i + 1) * RotateCost))
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
	TileTypeEmpty TileType = "."
	TileTypeWall  TileType = "#"
	TileTypeEnd   TileType = "E"
	TileTypeStart TileType = "S"
	TileTypePath  TileType = "*"
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	var start, end vector.V2
	m := Map{Region: map[vector.V2]map[vector.V2]Tile{}}
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
			}
			for _, d := range vector.Adjacent2 {
				if _, ok := m.Region[p]; !ok {
					m.Region[p] = map[vector.V2]Tile{}
				}
				m.Region[p][d] = Tile{Type: tt, Map: &m, Pos: p, Dir: d}
			}
		}
	}

	mdistance := math.MaxInt64
	for _, dir := range vector.OrthoAdjacent2 {
		path, distance, found := astar.Path(m.Region[start][vector.Cardinal2East], m.Region[end][dir])
		if found {
			mdistance = min(mdistance, int(distance))
		}

		// Visualization stuff:
		for _, p := range path {
			t := p.(Tile)
			mp := m.Region[t.Pos][vector.Cardinal2East]
			mp.Type = TileTypePath
			mp.Dir = t.Dir
			m.Region[t.Pos][vector.Cardinal2East] = mp
		}
		// fmt.Println(m.String())
	}

	return mdistance
}

func main() {
	filter := []int{}
	for i, c := range Cases {
		if slices.Contains(filter, i) {
			continue
		}
		log.Printf("case=%d expect=%d got=%d", i, c.Answer, Solve(c.Input))
	}
}
