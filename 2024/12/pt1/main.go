package main

import (
	_ "embed"
	"log"
	"math"
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
	//go:embed input.txt
	FullInput string
)

type Map [][]Tile

func (m Map) OutOfBounds(pos vector.V2) bool {
	return pos.X < 0 || pos.X >= len(m[0]) || pos.Y < 0 || pos.Y >= len(m)
}

func (m Map) String() string {
	str := ""
	for _, row := range m {
		strs := make([]string, len(row))
		for i, tile := range row {
			strs[i] = string(tile.Label)
		}
		str += strings.Join(strs, "") + "\n"
	}
	return str
}

type Tile struct {
	Label  string
	Region *Region
	Pos    vector.V2
}

func (t Tile) NonRegionDiagonalAdjacents(m Map) map[vector.V2]struct{} {
	nonRegionAdjacents := map[vector.V2]struct{}{}
	diagonalAdjacents := []vector.V2{
		t.Pos.Add(vector.V2{X: 1, Y: 1}),
		t.Pos.Add(vector.V2{X: -1, Y: 1}),
		t.Pos.Add(vector.V2{X: -1, Y: -1}),
		t.Pos.Add(vector.V2{X: 1, Y: -1}),
	}

	for _, adjacent := range diagonalAdjacents {
		if m.OutOfBounds(adjacent) {
			nonRegionAdjacents[adjacent] = struct{}{}
		} else if m[adjacent.Y][adjacent.X].Region != t.Region {
			nonRegionAdjacents[adjacent] = struct{}{}
		}
	}

	return nonRegionAdjacents
}

func (t Tile) NonRegionAdjacents(m Map) map[vector.V2]struct{} {
	nonRegionAdjacents := map[vector.V2]struct{}{}
	adjacents := []vector.V2{
		t.Pos.Add(vector.V2{X: 1}),
		t.Pos.Add(vector.V2{X: -1}),
		t.Pos.Add(vector.V2{Y: 1}),
		t.Pos.Add(vector.V2{Y: -1}),
	}

	for _, adjacent := range adjacents {
		if m.OutOfBounds(adjacent) {
			nonRegionAdjacents[adjacent] = struct{}{}
		} else if m[adjacent.Y][adjacent.X].Region != t.Region {
			nonRegionAdjacents[adjacent] = struct{}{}
		}
	}

	return nonRegionAdjacents
}

type Region struct {
	ID    RegionID
	Label string
	Tiles []*Tile
}

func (r Region) Area() int {
	return len(r.Tiles)
}

func (r Region) GetMinMaxCorners() (min vector.V2, max vector.V2) {
	min = vector.V2{X: math.MaxInt, Y: math.MaxInt}
	max = vector.V2{X: 0, Y: 0}
	for _, tile := range r.Tiles {
		if tile.Pos.X < min.X {
			min.X = tile.Pos.X
		}
		if tile.Pos.Y < min.Y {
			min.Y = tile.Pos.Y
		}
		if tile.Pos.X > max.X {
			max.X = tile.Pos.X
		}
		if tile.Pos.Y > max.Y {
			max.Y = tile.Pos.Y
		}
	}
	return
}

// P: 26
// OC: 5
// IC: 1
// NCE: 12

func (r *Region) CornersAndEdges(m Map) (int, int, int) {
	min, max := r.GetMinMaxCorners()
	nces := map[vector.V2]struct{}{}
	ics := map[vector.V2]struct{}{}
	ocs := map[vector.V2]struct{}{}

	// non-corner edges
	ncesc := 0
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if m[y][x].Region != r {
				continue
			}
			nras := m[y][x].NonRegionAdjacents(m)
			for nra := range nras {
				ncesc++
				nces[nra] = struct{}{}
			}
		}
	}

	// inner corners
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			pos := vector.V2{X: x, Y: y}
			if m[y][x].Region != r {
				continue
			} else if _, ok := nces[pos]; ok {
				continue
			}
			nrdas := m[y][x].NonRegionDiagonalAdjacents(m)
			if len(nrdas) != 1 {
				continue
			}
			for nrda := range nrdas {
				if _, ok := nces[nrda]; ok {
					continue
				}
				ics[nrda] = struct{}{}
			}
		}
	}

	// outer corners
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			pos := vector.V2{X: x, Y: y}
			if m[y][x].Region != r {
				continue
			} else if _, ok := nces[pos]; ok {
				continue
			} else if _, ok := ics[pos]; ok {
				continue
			}
			nrdas := m[y][x].NonRegionDiagonalAdjacents(m)
			if len(nrdas) < 2 {
				continue
			}
			for nrda := range nrdas {
				if _, ok := nces[nrda]; ok {
					continue
				} else if _, ok := ics[nrda]; ok {
					continue
				}
				ocs[nrda] = struct{}{}
			}
		}
	}

	return len(ocs), len(ics), ncesc
}

func (r *Region) Perimiter(m Map) int {
	oc, ic, ed := r.CornersAndEdges(m)
	_ = oc
	_ = ic
	perimiter := 0
	// perimiter += oc
	// perimiter += ic
	perimiter += ed

	return perimiter
}

type RegionID int

func Solve(input string) int {
	s := 0

	lines := strings.Split(input, "\n")
	m := make(Map, len(lines))
	for y, line := range lines {
		tiles := strings.Split(line, "")
		m[y] = make([]Tile, len(tiles))
		for x, tile := range tiles {
			m[y][x] = Tile{
				Label: tile,
				Pos:   vector.V2{X: x, Y: y},
			}
		}
	}

	regions := map[RegionID]*Region{}
	for y, row := range m {
		for x, tile := range row {
			if m[y][x].Region == nil {
				region := &Region{
					ID:    RegionID(len(regions) + 1),
					Label: tile.Label,
					Tiles: []*Tile{},
				}
				floodFillIdentify(m, vector.V2{X: x, Y: y}, region)
				regions[region.ID] = region
			}
		}
	}

	// fmt.Println(len(regions))
	for _, region := range regions {
		// min, max := region.GetMinMaxCorners()
		// oc, ic, nce := region.CornersAndEdges(m)
		// fmt.Println(region.Label, min, max, oc, ic, nce, region.Perimiter(m))
		s += region.Perimiter(m) * region.Area()
	}

	return s
}

func main() {
	log.Printf("sample1: %d", Solve(SampleInput1))
	// log.Printf("sample2: %d", Solve(SampleInput2))
	// log.Printf("sample3: %d", Solve(SampleInput3))
	// log.Printf("full: %d", Solve(FullInput))
}

func floodFillIdentify(m Map, pos vector.V2, region *Region) {
	if m.OutOfBounds(pos) {
		return
	} else if m[pos.Y][pos.X].Region != nil {
		return
	} else if m[pos.Y][pos.X].Label != region.Label {
		return
	}

	region.Tiles = append(region.Tiles, &m[pos.Y][pos.X])
	m[pos.Y][pos.X].Region = region
	floodFillIdentify(m, pos.Add(vector.V2{X: 1}), region)
	floodFillIdentify(m, pos.Add(vector.V2{X: -1}), region)
	floodFillIdentify(m, pos.Add(vector.V2{Y: 1}), region)
	floodFillIdentify(m, pos.Add(vector.V2{Y: -1}), region)
}
