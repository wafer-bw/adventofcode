package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
	//go:embed input-sample1.txt
	SampleInput1 string
	//go:embed input-sample2.txt
	SampleInput2 string
	//go:embed input-sample3.txt
	SampleInput3 string
	//go:embed input-sample4.txt
	SampleInput4 string
	//go:embed input-sample5.txt
	SampleInput5 string
	//go:embed input.txt
	FullInput string
)

type Map struct {
	Region     [][]*Tile
	RegionMap  map[vector.V2]*Tile
	Trailheads []*Tile
	Peaks      []*Tile
}

func (m Map) String() string {
	str := ""
	for _, row := range m.Region {
		strs := make([]string, len(row))
		for i, tile := range row {
			if tile == nil {
				strs[i] = "."
				continue
			}
			strs[i] = strconv.Itoa(tile.Height)
		}
		str += strings.Join(strs, "") + "\n"
	}
	return str
}

type Tile struct {
	Position   vector.V2
	Height     int
	Peak       bool
	Trailhead  bool
	Trailheads map[*Tile]struct{}
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	m := Map{Region: make([][]*Tile, len(lines))}
	m.RegionMap = map[vector.V2]*Tile{}
	for y, row := range lines {
		m.Region[y] = make([]*Tile, len(row))
		for x, tile := range row {
			h, err := strconv.Atoi(string(tile))
			if err != nil {
				continue
			}
			pos := vector.V2{X: x, Y: y}
			tile := &Tile{
				Position:   pos,
				Height:     h,
				Peak:       h == 9,
				Trailhead:  h == 0,
				Trailheads: map[*Tile]struct{}{},
			}
			m.Region[y][x] = tile
			m.RegionMap[pos] = tile
			if tile.Trailhead {
				m.Trailheads = append(m.Trailheads, tile)
			} else if tile.Peak {
				m.Peaks = append(m.Peaks, tile)
			}
		}
	}

	for _, trailhead := range m.Trailheads {
		m.Traverse(trailhead, trailhead, make(map[vector.V2]struct{}))
	}

	s := 0
	for _, peak := range m.Peaks {
		s += len(peak.Trailheads)
	}
	return s
}

func (m *Map) Traverse(startTile *Tile, tile *Tile, visited map[vector.V2]struct{}) {
	visited[tile.Position] = struct{}{}
	for _, adj := range vector.OrthoAdjacent2 {
		adjPos := tile.Position.Add(adj)
		if _, ok := visited[adjPos]; ok {
			continue
		}
		adjTile, ok := m.RegionMap[adjPos]
		if !ok || adjTile == nil || adjTile.Height != tile.Height+1 {
			continue
		}

		if adjTile.Peak {
			adjTile.Trailheads[startTile] = struct{}{}
		} else {
			m.Traverse(startTile, adjTile, visited)
		}
	}
}

func main() {
	log.Printf("sample1: %d", Solve(SampleInput1))
	log.Printf("sample2: %d", Solve(SampleInput2))
	log.Printf("sample3: %d", Solve(SampleInput3))
	log.Printf("sample4: %d", Solve(SampleInput4))
	log.Printf("sample5: %d", Solve(SampleInput5))
	log.Printf("full: %d", Solve(FullInput))
}
