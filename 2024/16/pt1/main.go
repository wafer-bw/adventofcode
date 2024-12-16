package main

import (
	_ "embed"
	"log"
	"slices"
	"strings"

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
	Region   map[vector.V2]Tile
	Maximums vector.V2
}

func (m Map) String() string {
	s := "\n"
	for y := 0; y <= m.Maximums.Y; y++ {
		for x := 0; x <= m.Maximums.X; x++ {
			if t, ok := m.Region[vector.V2{X: x, Y: y}]; ok {
				s += string(t.Type)
			}
		}
		s += "\n"
	}
	return s
}

type Tile struct {
	Type TileType
}

func (t Tile) Occupiable() bool {
	return t.Type != TileTypeWall
}

type TileType string

const (
	TileTypeEmpty TileType = "."
	TileTypeWall  TileType = "#"
	TileTypeEnd   TileType = "E"
	TileTypeStart TileType = "S"
)

type Player struct {
	Fac vector.V2 // facing
	Pos vector.V2 // position
	Spd int       // speed

	Cost int
}

// velocity
func (p Player) Vel() vector.V2 {
	return p.Fac.Mul(p.Spd)
}

// rotate
func (p *Player) Rot(dir vector.RotateDir) {
	p.Cost += RotateCost
	p.Fac = p.Fac.Rot(dir)
}

// move
func (p *Player) Mov() {
	p.Cost += MoveCost
	p.Pos = p.Pos.Add(p.Vel())
}

func Solve(input string) int {
	s := 0

	lines := strings.Split(input, "\n")
	m := Map{Region: make(map[vector.V2]Tile, len(lines)*len(lines[0]))}
	for y, line := range lines {
		if y == 0 {
			m.Maximums.X = (len(line) - 1)
			m.Maximums.Y = len(lines) - 1
		}
		for x, row := range strings.Split(line, "") {
			m.Region[vector.V2{X: x, Y: y}] = Tile{Type: TileType(row)}
		}
	}

	log.Println(m.String())

	return s
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
