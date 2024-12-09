package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type Map [][]Tile

func (m Map) String() string {
	str := ""
	for _, row := range m {
		strs := make([]string, len(row))
		for i, tile := range row {
			strs[i] = string(tile)
		}
		str += strings.Join(strs, "") + "\n"
	}
	return str
}

func (m Map) AddAntinode(antinode vector.V2) bool {
	if antinode.X < 0 || antinode.Y < 0 || antinode.X >= len(m[0]) || antinode.Y >= len(m) {
		return false
	}
	m[antinode.Y][antinode.X] = TileAntinode
	return true
}

func (m Map) CountAntinodes() int {
	count := 0
	for _, row := range m {
		for _, tile := range row {
			if tile == TileAntinode {
				count++
			}
		}
	}
	return count
}

type Tile string

const (
	TileEmpty    Tile = "."
	TileAntinode Tile = "#"
)

type Antenna struct {
	Freq string
	Pos  vector.V2
}

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	m := make(Map, len(lines))
	antennas := []Antenna{}
	for y, line := range lines {
		tiles := strings.Split(line, "")
		m[y] = make([]Tile, len(tiles))
		for x, tile := range tiles {
			m[y][x] = Tile(tile)
			if m[y][x] != TileEmpty {
				antennas = append(antennas, Antenna{Freq: tile, Pos: vector.V2{X: x, Y: y}})
			}
		}
	}

	for _, antennaA := range antennas {
		for _, antennaB := range antennas {
			if antennaA.Freq != antennaB.Freq {
				continue
			} else if antennaA.Pos == antennaB.Pos {
				continue
			}

			odist := antennaA.Pos.OrthoDistance(antennaB.Pos)
			antinode := antennaA.Pos.Add(odist)
			_ = m.AddAntinode(antinode)
			slope := antennaA.Pos.OrthoSlope(antennaB.Pos)
			negSlope := slope.Neg()
			an := antinode
			for {
				an = an.Add(slope)
				if ok := m.AddAntinode(an); !ok {
					break
				}
			}

			antinode = antennaB.Pos.Sub(odist)
			_ = m.AddAntinode(antinode)
			an = antinode
			for {
				an = an.Add(negSlope)
				if ok := m.AddAntinode(an); !ok {
					break
				}
			}
		}
	}
	return m.CountAntinodes()
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
