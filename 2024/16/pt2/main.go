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

	// there is likely a better way to do this block.
	facingRight := t.Dir
	facingLeft := t.Dir
	for i := 1; i <= 2; i++ {
		facingRight = facingRight.Rot(vector.RotateDirRight)
		if facingRight == approach {
			return t.ExtraCost(toT.Pos, facingRight, float64(c+(i*RotateCost)))
		}
		facingLeft = facingLeft.Rot(vector.RotateDirLeft)
		if facingLeft == approach {
			return t.ExtraCost(toT.Pos, facingLeft, float64(c+(i*RotateCost)))
		}
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

func (m *Map) AltpathPeek(p1 astar.Pather, step, score int, visited map[vector.V2]struct{}, chain map[string]struct{}) (map[vector.V2]struct{}, bool) {
	foundPath := false
	for _, p2 := range p1.Neighbors() {
		for _, dir := range vector.OrthoAdjacent2 {
			t1 := p1.(Tile)
			t2 := p2.(Tile)
			t2.Dir = dir

			// anti-loop
			if _, ok := chain[t2.InverseKey()]; ok {
				continue
			}
			chain[t1.Key()] = struct{}{}

			cost := int(t1.NeighborCost(t2))
			// fmt.Println(step, t1.Pos, t1.Dir.ToDirSymbol(), t2.Pos, t2.Dir.ToDirSymbol(), cost, score, len(visited))
			// fmt.Scanln()

			if t2.Pos == m.Start /*&& t2.Dir == StartDir.Neg()*/ {
				fmt.Println("found end", len(visited), len(chain))
				return visited, true
			} else if score <= 0 {
				return nil, false
			} else if score > 0 {
				nap := map[vector.V2]struct{}{}
				for k, v := range visited {
					nap[k] = v
				}
				nap[t1.Pos] = struct{}{}
				nap[t2.Pos] = struct{}{}
				ap, ok := m.AltpathPeek(t2, step+1, score-cost, nap, chain)
				if ok {
					for k, v := range ap {
						visited[k] = v
					}
					foundPath = true
				}
			}
		}
	}

	if foundPath {
		return visited, true
	}
	return nil, false
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
	// mscores := map[astar.Pather]float64{}
	mdistance := math.MaxInt64
	for _, dir := range vector.OrthoAdjacent2 {
		if path, distance, _, found := astar.Path(m.Region[start][vector.Cardinal2East], m.Region[end][dir]); found {
			mdistance = min(mdistance, int(distance))
			mpath = path
			// mscores = scores
		}
	}
	fmt.Println("cost", mdistance)

	// // view scores by step
	// why are these scores higher than distance?
	// for _, p := range mpath {
	// 	fmt.Println(mscores[p])
	// }

	t1 := mpath[0].(Tile)
	t1.Dir = EndDir.Neg()
	visited, ok := m.AltpathPeek(t1, 0, int(mdistance), map[vector.V2]struct{}{}, map[string]struct{}{})
	if !ok {
		panic("no path found")
	}

	for pos := range visited {
		mp := m.Region[pos][vector.Cardinal2East]
		mp.Type = TileTypeGoodSeat
		m.Region[pos][vector.Cardinal2East] = mp
	}
	fmt.Println(m.String())

	return len(visited)
}

// full input answer is 497 < x < 604

func main() {
	filter := []int{1, 2}
	for i, c := range Cases {
		if slices.Contains(filter, i) {
			continue
		}
		log.Printf("case=%d expect=%d got=%d", i, c.Answer, Solve(c.Input))
	}
}
