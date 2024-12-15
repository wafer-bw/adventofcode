package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/vector"
)

var (
	//go:embed input-sample1.txt
	SampleInput1 string
	//go:embed input-sample2.txt
	SampleInput2 string
	//go:embed input.txt
	FullInput string
)

type Map struct {
	MaxX    int
	MaxY    int
	Region  map[vector.V2]Tile
	Robot   vector.V2
	Steps   int
	Moveset MoveSet
}

type MoveSet []Move

func (m MoveSet) String() string {
	s := ""
	for _, move := range m {
		s += move.S
	}
	return s
}

type Move struct {
	S string
	V vector.V2
}

func (m *Map) Step() bool {
	if len(m.Moveset) == 0 {
		return false
	}
	m.Steps++
	inst := m.Moveset[0]
	m.Moveset = m.Moveset[1:]

	nextEmpty, ok := m.EmptyTowards(m.Robot, inst.V)
	if ok {
		m.PushTowards(m.Robot, inst.V, nextEmpty)
	}
	return true
}

func (m Map) EmptyTowards(p, d vector.V2) (vector.V2, bool) {
	dst := p.Add(d)
	t, ok := m.Region[dst]
	if !ok {
		panic("out of bounds")
	}

	if t.Type == TileTypeWall {
		return vector.V2{}, false
	} else if t.Type == TileTypeBox {
		return m.EmptyTowards(dst, d)
	} else if t.Type == TileTypeOpen {
		return dst, true
	}

	panic("unexpected terminal")
}

func (m *Map) PushTowards(p, d, trm vector.V2) {
	if m.Region[trm].Type != TileTypeOpen {
		panic("trm is not open")
	}

	p1 := trm
	for {
		np := d.Neg()
		p2 := p1.Add(np)
		m.Region[p1] = m.Region[p2]
		if m.Region[p1].Type == TileTypeRobot {
			m.Robot = p1
			m.Region[p2] = Tile{Type: TileTypeOpen}
			return
		} else if m.Region[p1].Type == TileTypeWall {
			panic("wall")
		}
		p1 = p2
	}
}

func (m *Map) Score() int {
	s := 0
	for p, t := range m.Region {
		if t.Type == TileTypeBox {
			s += 100*p.Y + p.X
		}
	}
	return s
}

var moveMap map[string]vector.V2 = map[string]vector.V2{
	"^": vector.Cardinal2North,
	">": vector.Cardinal2East,
	"v": vector.Cardinal2South,
	"<": vector.Cardinal2West,
}

func (m *Map) String() string {
	s := ""
	for y := 0; y <= m.MaxY; y++ {
		for x := 0; x <= m.MaxX; x++ {
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

type TileType string

const (
	TileTypeRobot TileType = "@"
	TileTypeWall  TileType = "#"
	TileTypeOpen  TileType = "."
	TileTypeBox   TileType = "O"
)

func Solve(input string, manual ...bool) int {
	m := Map{Region: map[vector.V2]Tile{}, Moveset: MoveSet{}}
	lines := strings.Split(input, "\n")
	mapScan := true
	for i, line := range lines {
		if i == 0 {
			m.MaxX = len(line) - 1
		} else if line == "" {
			mapScan = false
			m.MaxY = i
			continue
		}

		if mapScan {
			for j, c := range strings.Split(line, "") {
				m.Region[vector.V2{X: j, Y: i}] = Tile{Type: TileType(c)}
				if c == string(TileTypeRobot) {
					m.Robot = vector.V2{X: j, Y: i}
				}
			}
		} else {
			for _, c := range strings.Split(line, "") {
				m.Moveset = append(m.Moveset, Move{S: c, V: moveMap[c]})
			}
		}
	}

	if len(manual) > 0 && manual[0] {
		fmt.Println(m.String())
		fmt.Println(m.Score())
		fmt.Println(m.Moveset)
		_, _ = fmt.Scanln()
	}

	for m.Step() {
		if len(manual) > 0 && manual[0] {
			fmt.Println(m.String())
			fmt.Println(m.Moveset)
			_, _ = fmt.Scanln()
		}
	}

	if len(manual) > 0 && manual[0] {
		fmt.Println(m.Score())
	}

	return m.Score()
}

func main() {
	log.Printf("sample1: %d", Solve(SampleInput1))
	log.Printf("sample2: %d", Solve(SampleInput2))
	log.Printf("full: %d", Solve(FullInput))
}
