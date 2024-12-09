package main

import (
	_ "embed"
	"fmt"
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

	antinodes := map[vector.V2]struct{}{}
	for _, antennaA := range antennas {
		for _, antennaB := range antennas {
			if antennaA.Freq != antennaB.Freq {
				continue
			} else if antennaA.Pos == antennaB.Pos {
				continue
			}

			odist := antennaA.Pos.OrthoDistance(antennaB.Pos)
			fmt.Println(antennaA, antennaB, odist)
			antinode := antennaA.Pos.Add(odist)
			if ok := m.AddAntinode(antinode); ok {
				antinodes[antinode] = struct{}{}
			}
			antinode = antennaB.Pos.Sub(odist)
			if ok := m.AddAntinode(antinode); ok {
				antinodes[antinode] = struct{}{}
			}
		}
	}

	return len(antinodes)
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
