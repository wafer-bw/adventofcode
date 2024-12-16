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

type MoveAction struct {
	From vector.V2
	To   vector.V2
}

func (ma MoveAction) String() string {
	return fmt.Sprintf("%s -> %s", ma.From, ma.To)
}

type Map struct {
	MaxX    int
	MaxY    int
	Region  map[vector.V2]Tile
	Robot   vector.V2
	Steps   int
	Moveset MoveSet
}

func (m *Map) Step() bool {
	if len(m.Moveset) == 0 {
		return false
	}
	m.Steps++
	inst := m.Moveset[0]
	m.Moveset = m.Moveset[1:]

	mas, ok := m.EmptyTowards(m.Robot, inst.V, []MoveAction{})
	if ok {
		m.PerformMoves(mas)
	}
	return true
}

func (m Map) EmptyTowards(p, d vector.V2, mas []MoveAction) ([]MoveAction, bool) {
	dst := p.Add(d)
	t, ok := m.Region[dst]
	if !ok {
		panic("out of bounds")
	}

	if t.Type == TileTypeWall {
		return nil, false
	} else if (d == vector.Cardinal2East || d == vector.Cardinal2West) && (t.Type == TileTypeBox1 || t.Type == TileTypeBox2) {
		mas = append(mas, MoveAction{From: p, To: dst})
		mas, ok := m.EmptyTowards(dst, d, mas)
		return mas, ok
	} else if (d == vector.Cardinal2North || d == vector.Cardinal2South) && (t.Type == TileTypeBox1 || t.Type == TileTypeBox2) {
		var dst2 vector.V2
		if t.Type == TileTypeBox1 {
			dst2 = dst.Add(vector.V2{X: 1})
		} else {
			dst2 = dst.Add(vector.V2{X: -1})
		}
		mas = append(mas, MoveAction{From: p, To: dst})
		mas, ok1 := m.EmptyTowards(dst, d, mas)
		mas, ok2 := m.EmptyTowards(dst2, d, mas)
		return mas, ok1 && ok2
	} else if t.Type == TileTypeOpen {
		mas = append(mas, MoveAction{From: p, To: dst})
		return mas, true
	}

	panic("unexpected terminal")
}

func (m *Map) PerformMoves(mas []MoveAction) {
	performed := map[string]struct{}{}
	for len(mas) != 0 {
		for i, ma := range mas {
			if _, ok := performed[ma.String()]; ok {
				mas = append(mas[:i], mas[i+1:]...)
				break
			}
			if m.Region[ma.To].Type == TileTypeOpen {
				m.Region[ma.To] = m.Region[ma.From]
				if m.Region[ma.To].Type == TileTypeRobot {
					m.Robot = ma.To
				}
				m.Region[ma.From] = Tile{Type: TileTypeOpen}
				mas = append(mas[:i], mas[i+1:]...)
				performed[ma.String()] = struct{}{}
				break
			}

		}
	}
}

func (m *Map) Score() int {
	s := 0
	for p, t := range m.Region {
		if t.Type == TileTypeBox1 {
			s += 100*p.Y + p.X
		}
	}
	return s
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

var moveMap map[string]vector.V2 = map[string]vector.V2{
	"^": vector.Cardinal2North,
	">": vector.Cardinal2East,
	"v": vector.Cardinal2South,
	"<": vector.Cardinal2West,
}

var expandMap map[TileType][]TileType = map[TileType][]TileType{
	TileTypeWall:  {TileTypeWall, TileTypeWall},
	TileTypeOpen:  {TileTypeOpen, TileTypeOpen},
	TileTypeBox:   {TileTypeBox1, TileTypeBox2},
	TileTypeRobot: {TileTypeRobot, TileTypeOpen},
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
	TileTypeBox1  TileType = "["
	TileTypeBox2  TileType = "]"
)

func Solve(input string, manual ...bool) int {
	m := Map{Region: map[vector.V2]Tile{}, Moveset: MoveSet{}}
	lines := strings.Split(input, "\n")
	mapScan := true
	for i, line := range lines {
		if i == 0 {
			m.MaxX = (len(line) - 1) * 2
		} else if line == "" {
			mapScan = false
			m.MaxY = i - 1
			continue
		}

		if mapScan {
			k := 0
			for _, c := range strings.Split(line, "") {
				exp := expandMap[TileType(c)]
				m.Region[vector.V2{X: k, Y: i}] = Tile{Type: exp[0]}
				if c == string(TileTypeRobot) {
					m.Robot = vector.V2{X: k, Y: i}
				}
				k++
				m.Region[vector.V2{X: k, Y: i}] = Tile{Type: exp[1]}
				k++
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

	// fmt.Println(m.String())

	return m.Score()
}

func main() {
	log.Printf("sample1: %d", Solve(SampleInput1))
	log.Printf("sample2: %d", Solve(SampleInput2))
	log.Printf("full: %d", Solve(FullInput))
}
