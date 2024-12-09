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

func (d Drive) FirstEmptyBlockSpace(size int) (int, bool) {
	for i, b := range d {
		if b.File == nil {
			space := 0
			for j := i; j < len(d); j++ {
				if d[j].File != nil {
					break
				}
				space++
			}
			if space >= size {
				return i, true
			}
		}
	}
	return 0, false
}

func (d Drive) Files() []*File {
	seen := map[*File]struct{}{}
	files := []*File{}
	for _, b := range d {
		if b.File == nil {
			continue
		}
		if _, ok := seen[b.File]; ok {
			continue
		}
		seen[b.File] = struct{}{}
		files = append(files, b.File)
	}
	return files
}

func (d Drive) FileIndex(file *File) (int, bool) {
	for i, b := range d {
		if b.File == file {
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

	filesList := drive.Files()
	for i := len(filesList) - 1; i >= 0; i-- {
		file := filesList[i]
		fileIdx, ok := drive.FileIndex(file)
		if !ok {
			panic("file not found")
		}

		emptyBlockIdx, eOk := drive.FirstEmptyBlockSpace(file.Size)
		if !eOk {
			continue
		}

		if emptyBlockIdx > fileIdx {
			continue
		}

		if err := slics.SwapChunk(drive, emptyBlockIdx, fileIdx, file.Size); err != nil {
			panic(err)
		}
	}

	for i, b := range drive {
		if b.File == nil {
			continue
		}
		s += i * b.File.ID
	}

	return s
}

func main() {
	log.Printf("sample: %d", Solve(SampleInput))
	log.Printf("full: %d", Solve(FullInput))
}
