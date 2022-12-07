package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID          string = "2022-07"
	spaceNeeded       int    = 30000000
	filesystemMaxSize int    = 70000000
)

type fs struct {
	name string
	dir  bool
	abv  *fs
	blw  []*fs
	size *int
}

func (f *fs) getDirSize() int {
	size := 0
	for _, f := range f.blw {
		if f.dir {
			size += f.getDirSize()
		} else {
			size += *f.size
		}
	}
	f.size = &size
	return size
}

func solve(lines []string) int {
	var filesystem *fs = &fs{
		name: "root",
		dir:  true,
		blw:  []*fs{},
		abv:  nil,
	}

	cmd, trg := "", ""
	at := filesystem
	for _, ln := range lines {
		// log.Println(at.name, ln)

		if strings.HasPrefix(ln, "$ ") {
			cmd, trg = getCommand(ln)
			switch cmd {
			case "ls":
				continue
			case "cd":
				if trg == ".." {
					at = at.abv
					continue
				} else {
					at.blw = append(at.blw, &fs{
						name: trg,
						dir:  true,
						abv:  at,
					})
					at = at.blw[len(at.blw)-1]
				}
				// cmd, trg = "", ""
				continue
			}
		} else if cmd != "" {
			if strings.HasPrefix(ln, "dir") {
				at.blw = append(at.blw, &fs{
					name: strings.Split(ln, " ")[1],
					dir:  true,
				})
				continue
			} else {
				fileParts := strings.Split(ln, " ")
				fileSize, _ := strconv.Atoi(fileParts[0])
				fileName := fileParts[1]
				at.blw = append(at.blw, &fs{
					name: fileName,
					dir:  false,
					size: &fileSize,
					abv:  at,
				})
			}
		}
	}

	filesystem.getDirSize()
	walk(filesystem, "")
	unused := filesystemMaxSize - *filesystem.size
	need := spaceNeeded - unused
	return getDirToDelete(filesystem, need)
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, true, false))))
}

func getDirToDelete(fs *fs, need int) int {
	size := filesystemMaxSize
	for _, f := range fs.blw {
		if f.dir && *f.size >= need && *f.size < size {
			size = *f.size
		}
		ns := getDirToDelete(f, need)
		if ns < size && ns >= need {
			size = ns
		}
	}

	return size
}

func walk(fs *fs, idt string) {
	if !fs.dir {
		return
	} else {
		for _, f := range fs.blw {
			walk(f, idt+"\t-")
		}
	}
}

func getCommand(ln string) (string, string) {
	cmd, trg := "", ""

	cmdFull := strings.TrimPrefix(ln, "$ ")
	cmdParts := strings.Split(cmdFull, " ")

	cmd = cmdParts[0]
	if len(cmdParts) > 1 {
		trg = cmdParts[1]
	}

	return cmd, trg
}
