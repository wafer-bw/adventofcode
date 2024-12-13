package main

import (
	_ "embed"
	"fmt"
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
	//go:embed input-sample4.txt
	SampleInput4 string
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

func (r *Region) Perimiter(m Map) int {
	sides := 0
	min, max := r.GetMinMaxCorners()
	for y := min.Y - 1; y <= max.Y+1; y++ {
		for x := min.X - 1; x <= max.X+1; x++ {
			sides += r.CountCorners(m, vector.V2{X: x, Y: y})
		}
	}

	return sides
}

func (r *Region) CountCorners(m Map, pos vector.V2) int {
	inRegion := 0
	walks := []vector.V2{
		{Y: -1},
		{X: 1},
		{Y: 1},
		{X: -1},
	}

	inRegions := map[int]struct{}{}
	for i, walk := range walks {
		pos = pos.Add(walk)
		if m.OutOfBounds(pos) {
			continue
		}
		if m[pos.Y][pos.X].Region.ID == r.ID {
			inRegions[i] = struct{}{}
			inRegion++
		}
	}

	_, ok0 := inRegions[0]
	_, ok1 := inRegions[1]
	_, ok2 := inRegions[2]
	_, ok3 := inRegions[3]
	if (ok0 && !ok1 && ok2 && !ok3) || (!ok0 && ok1 && !ok2 && ok3) {
		return 2
	}

	if inRegion == 3 || inRegion == 1 {
		return 1
	}
	return 0
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

	for _, region := range regions {
		p := region.Perimiter(m)
		a := region.Area()
		s += p * a
		fmt.Println(region.Label, p)
	}

	return s
}

func main() {
	// log.Printf("sample1: %d", Solve(SampleInput1))
	// log.Printf("sample2: %d", Solve(SampleInput2))
	log.Printf("sample3: %d", Solve(SampleInput3))
	// log.Printf("sample4: %d", Solve(SampleInput4))
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
