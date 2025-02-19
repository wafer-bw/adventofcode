package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/wafer-bw/adventofcode/2024/16/astar"
	"github.com/wafer-bw/adventofcode/tools/stack"
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

var (
	StartDir = vector.Cardinal2East
	EndDir   = vector.Cardinal2North // this is a simplification due to my inputs.
)

type Map struct {
	Region      map[vector.V2]map[vector.V2]Tile // Pos->Rotation->Tile
	Maximums    vector.V2
	Start       vector.V2
	End         vector.V2
	VisitedMode bool
	Visited     map[vector.V2]map[vector.V2]int
	Seats       map[vector.V2]struct{}
	Comparisons map[astar.Pather]map[astar.Pather]struct{}
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

func (t Tile) Key() string {
	return t.Pos.String() + t.Dir.String()
}

func (t Tile) InverseKey() string {
	return t.Pos.String() + t.Dir.Neg().String()
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

func (t Tile) NeighborCost(to astar.Pather) float64 {
	toT := to.(Tile)
	c := 1
	facing := t.Dir
	approach := toT.Pos.Sub(t.Pos)

	if facing == approach {
		return float64(c)
	}

	facingRight := t.Dir
	facingRight = facingRight.Rot(vector.RotateDirRight)
	if facingRight == approach {
		return float64(c + (RotateCost))
	}

	facingLeft := t.Dir
	facingLeft = facingLeft.Rot(vector.RotateDirLeft)
	if facingLeft == approach {
		return float64(c + (RotateCost))
	}

	return math.MaxFloat32
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
	TileTypeGoodSeat TileType = "O"
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

func (m Map) Altpaths(p1 astar.Pather, scores map[astar.Pather]float64) map[vector.V2]struct{} {
	step := 0
	stk := &[]Tile{p1.(Tile)}
	seats := map[vector.V2]struct{}{p1.(Tile).Pos: {}}
	visited := map[string]struct{}{}

	for stack.Len(stk) > 0 {
		step++
		t1 := stack.Pop(stk)

		if _, ok := visited[t1.Key()]; ok {
			continue
		}
		visited[t1.Key()] = struct{}{}

		faceLeft := t1.Dir.Rot(vector.RotateDirLeft)
		faceRight := t1.Dir.Rot(vector.RotateDirRight)

		// these values may look weird because we are going in reverse.
		for _, t2 := range []Tile{
			m.Region[t1.Pos.Add(t1.Dir.Neg())][t1.Dir],      // forward
			m.Region[t1.Pos.Add(faceLeft.Neg())][faceRight], // left
			m.Region[t1.Pos.Add(faceRight.Neg())][faceLeft], // right
			m.Region[t1.Pos.Add(t1.Dir.Neg())][faceRight],   // forwardLeft
			m.Region[t1.Pos.Add(t1.Dir.Neg())][faceLeft],    // forwardRight
		} {
			cost := t2.NeighborCost(t1)
			t1Score := scores[m.Region[t1.Pos][t1.Dir]]
			t2Score := scores[m.Region[t2.Pos][t2.Dir]]
			if t1Score-cost == t2Score {
				stack.Push(stk, t2)
				seats[t2.Pos] = struct{}{}
			}
		}
	}

	return seats
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	var start, end vector.V2
	m := Map{Region: map[vector.V2]map[vector.V2]Tile{}, Visited: map[vector.V2]map[vector.V2]int{}, Comparisons: map[astar.Pather]map[astar.Pather]struct{}{}, Seats: map[vector.V2]struct{}{}}
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
				m.Start = p
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
	mpath := []astar.Pather{}
	mscores := map[astar.Pather]float64{}
	mdistance := math.MaxInt64
	if path, distance, scores, found := astar.Path(m.Region[start][StartDir], m.Region[end][EndDir]); found {
		mdistance = min(mdistance, int(distance))
		mpath = path
		mscores = scores
	}
	fmt.Println("cost", mdistance)

	t1 := mpath[0].(Tile)
	t1.Dir = EndDir
	visited := m.Altpaths(t1, mscores)

	for pos := range visited {
		mp := m.Region[pos][vector.Cardinal2East]
		mp.Type = TileTypeGoodSeat
		m.Region[pos][vector.Cardinal2East] = mp
	}
	fmt.Println(m.String())

	return len(visited)
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
