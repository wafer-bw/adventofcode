package main

import (
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/wafer-bw/always"
)

type Nodes []*Node

type Node struct {
	nature Nature
	number int
	src    *Node
	dst    *Node
}

type Nature int

const (
	NatureSeed Nature = iota
	NatureSoil
	NatureFertilizer
	NatureWater
	NatureLight
	NatureTemp
	NatureHumidity
	NatureLocation
)

var _ = map[Nature]string{
	NatureSeed:       "seed",
	NatureSoil:       "soil",
	NatureFertilizer: "fertilizer",
	NatureWater:      "water",
	NatureLight:      "light",
	NatureTemp:       "temp",
	NatureHumidity:   "humidity",
	NatureLocation:   "location",
}

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")
	seeds := strings.Split(lines[0], " ")[1:]

	nodes := Nodes{}
	for _, seed := range seeds {
		nodes = append(nodes, &Node{
			nature: NatureSeed,
			number: always.Accept(strconv.Atoi(seed)),
		})
	}

	mode := NatureSeed
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		} else if strings.Contains(line, ":") {
			nodes = populateMissing(nodes, mode)
			mode++
			continue
		}

		parts := strings.Split(line, " ")
		dst := always.Accept(strconv.Atoi(parts[0]))
		src := always.Accept(strconv.Atoi(parts[1]))
		length := always.Accept(strconv.Atoi(parts[2]))

		nodes = populateLinked(src, dst, length, nodes, mode)
	}

	nodes = populateMissing(nodes, mode)

	v := math.MaxInt64
	for _, n := range nodes {
		if n.nature == NatureLocation {
			if n.number < v {
				v = n.number
			}
		}
	}

	return v
}

func populateMissing(nodes Nodes, mode Nature) Nodes {
	for _, n := range nodes {
		if n.nature != mode-1 {
			continue
		}
		if n.dst == nil {
			n.dst = &Node{
				nature: mode,
				number: n.number,
				src:    n,
			}
			nodes = append(nodes, n.dst)
		}
	}

	return nodes
}

func populateLinked(src, dst, length int, nodes Nodes, mode Nature) Nodes {
	for _, n := range nodes {
		if n.nature != mode-1 {
			continue
		}

		if n.number >= src && n.number <= src+length-1 {
			n.dst = &Node{
				src:    n,
				nature: mode,
				number: n.number + dst - src,
			}
			nodes = append(nodes, n.dst)
		}

	}

	return nodes
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
