package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/mth"
	"github.com/wafer-bw/always"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

func Solve(input string) int {
	lines := strings.Split(input, "\n")

	resolvedSequences := [][][]int{}
	for _, ln := range lines {
		seq := lineToSeq(ln)
		res := resolveSeq(seq)
		resolvedSequences = append(resolvedSequences, res)
	}

	s := 0
	for _, rs := range resolvedSequences {
		s += extropolateSeq(rs)
	}

	return s
}

func resolveSeq(seq []int) [][]int {
	sequences := [][]int{seq}
	for {
		nseq := []int{}
		for i := 0; i < len(seq)-1; i++ {
			nseq = append(nseq, seq[i+1]-seq[i])
		}
		sequences = append(sequences, nseq)
		seq = nseq
		if mth.Sum(seq) == 0 {
			break
		}
	}

	return sequences
}

func extropolateSeq(seq [][]int) int {
	for i := len(seq) - 1; i >= 0; i-- {
		if i == len(seq)-1 {
			seq[i] = append([]int{0}, seq[i]...)
			continue
		}
		left := seq[i][len(seq[i])-1]
		below := seq[i+1][len(seq[i+1])-1]
		seq[i] = append(seq[i], left+below)
	}

	return seq[0][len(seq[0])-1]
}

func lineToSeq(ln string) []int {
	seq := []int{}
	for _, ch := range strings.Split(ln, " ") {
		seq = append(seq, always.Accept(strconv.Atoi(ch)))
	}
	return seq
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
