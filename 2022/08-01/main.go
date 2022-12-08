package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID string = "2022-08"
)

type tree struct {
	height   int
	visibleN bool
	visibleE bool
	visibleS bool
	visibleW bool
}

func (t tree) IsVisible() bool {
	return t.visibleN || t.visibleE || t.visibleS || t.visibleW
}

func solve(lines []string) int {
	grid, trees := buildTreesGrid(lines)
	observeForest(grid)
	drawForest(grid)

	visibleTrees := 0
	for _, tree := range trees {
		if tree.IsVisible() {
			visibleTrees++
		}
	}
	return visibleTrees
}

func main() {
	log.Println(solve(reader.Read(pather.Path(puzzleID, false, false))))
}

func buildTreesGrid(lines []string) ([][]*tree, []*tree) {
	trees := []*tree{}
	grid := make([][]*tree, len(lines))

	for i, ln := range lines {
		for j, c := range ln {
			h, _ := strconv.Atoi(string(c))
			tree := &tree{height: h}

			if i == 0 {
				tree.visibleN = true
			} else if j == 0 {
				tree.visibleW = true
			} else if i == len(lines)-1 {
				tree.visibleS = true
			} else if j == len(ln)-1 {
				tree.visibleE = true
			}

			grid[i] = append(grid[i], tree)
			trees = append(trees, tree)
		}
	}

	return grid, trees
}

func drawForest(grid [][]*tree) {
	for i := range grid {
		row := ""
		for j := range grid[i] {
			tree := grid[i][j]
			if tree.IsVisible() {
				row += fmt.Sprintf("%d", tree.height) // fmt.Sprintf("%dV", tree.height)
			} else {
				row += "_" // fmt.Sprintf("%dX", tree.height)
			}
		}
		log.Println(row)
	}
}

func observeForest(grid [][]*tree) {
	tallestN, tallestW := map[int]int{}, map[int]int{}
	for i := range grid {
		for j := range grid[i] {
			tree := grid[i][j]

			if _, ok := tallestN[i]; !ok {
				tallestN[i] = tree.height
			}
			if _, ok := tallestW[j]; !ok {
				tallestW[j] = tree.height
			}

			if tree.height > tallestN[i] {
				tallestN[i] = tree.height
				grid[i][j].visibleN = true
			}

			if tree.height > tallestW[j] {
				tallestW[j] = tree.height
				grid[i][j].visibleW = true
			}
		}
	}

	tallestE, tallestS := map[int]int{}, map[int]int{}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			tree := grid[i][j]

			if _, ok := tallestE[j]; !ok {
				tallestE[j] = tree.height
			}
			if _, ok := tallestS[i]; !ok {
				tallestS[i] = tree.height
			}

			if tree.height > tallestE[j] {
				tallestE[j] = tree.height
				grid[i][j].visibleE = true
			}

			if tree.height > tallestS[i] {
				tallestS[i] = tree.height
				grid[i][j].visibleS = true
			}
		}
	}
}
