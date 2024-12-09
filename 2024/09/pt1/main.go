package main

import (
	_ "embed"
	"log"
	"strconv"

	"github.com/wafer-bw/adventofcode/tools/slics"
)

var (
	//go:embed input-sample.txt
	SampleInput string
	//go:embed input.txt
	FullInput string
)

type File struct {
	ID   int
	Size int
}

type Block struct {
	File *File
}

type Drive []Block

func (d Drive) String() string {
	s := ""
	for _, b := range d {
		if b.File == nil {
			s += "."
		} else {
			// s += "[" + strconv.Itoa(b.File.ID) + "]"
			s += strconv.Itoa(b.File.ID)
		}
	}
	return s
}

func (d Drive) FirstEmptyBlock() (int, bool) {
	for i, b := range d {
		if b.File == nil {
			return i, true
		}
	}
	return 0, false
}

func (d Drive) LastFilledBlock() (int, bool) {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i].File != nil {
			return i, true
		}
	}
	return 0, false
}

func Solve(input string) int {
	s := 0

	drive := Drive{}
	fileMap := map[int]*File{}
	for i := 0; i < len(input); i += 2 {
		id := i / 2
		size, _ := strconv.Atoi(string(input[i]))
		emptySpace := 0
		if i < len(input)-1 {
			emptySpace, _ = strconv.Atoi(string(input[i+1]))
		}

		file := &File{ID: id, Size: size}
		fileMap[id] = file

		for j := 0; j < size; j++ {
			drive = append(drive, Block{File: file})
		}
		for j := 0; j < emptySpace; j++ {
			drive = append(drive, Block{})
		}
	}

	for {
		emptyBlockIdx, eOk := drive.FirstEmptyBlock()
		if !eOk {
			break
		}
		filledBlockIdx, fOk := drive.LastFilledBlock()
		if !fOk {
			break
		}

		if emptyBlockIdx > filledBlockIdx {
			break
		}

		slics.Swap(drive, emptyBlockIdx, filledBlockIdx)
	}

	for i, b := range drive {
		if b.File == nil {
			break
		}
		s += i * b.File.ID
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
