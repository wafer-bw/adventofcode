// https://adventofcode.com/2022/day/8#part2

package main

import (
	"log"
	"strconv"

	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

const (
	puzzleID string = "2022-08"
)

type tree struct {
	height       int
	x            int
	y            int
	visibleN     bool
	visibleE     bool
	visibleS     bool
	visibleW     bool
	scenicScoreN int
	scenicScoreE int
	scenicScoreS int
	scenicScoreW int
}

func (t tree) IsVisible() bool {
	return t.visibleN || t.visibleE || t.visibleS || t.visibleW
}

func (t tree) ScenicScore() int {
	return t.scenicScoreN * t.scenicScoreE * t.scenicScoreS * t.scenicScoreW
}

func solve(lines []string) int {
	grid, trees := buildTreesGrid(lines)
	getScenicScores(trees, grid)

	highestScenicScore := 0
	for _, tree := range trees {
		if tree.ScenicScore() > highestScenicScore {
			highestScenicScore = tree.ScenicScore()
		}
	}
	return highestScenicScore
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
			tree := &tree{height: h, x: j, y: i}

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

func getScenicScores(trees []*tree, grid [][]*tree) {
	for idx, tree := range trees {
		x, y := tree.x, tree.y

		// Norhtward View
		for i := x + 1; i <= len(grid[y])-1; i++ {
			if grid[y][i].height < tree.height {
				trees[idx].scenicScoreN++
			} else {
				trees[idx].scenicScoreN++
				break
			}
		}

		// Southward View
		for i := x - 1; i >= 0; i-- {
			if grid[y][i].height < tree.height {
				trees[idx].scenicScoreS++
			} else {
				trees[idx].scenicScoreS++
				break
			}
		}

		// Eastward View
		for i := y + 1; i <= len(grid)-1; i++ {
			if grid[i][x].height < tree.height {
				trees[idx].scenicScoreE++
			} else {
				trees[idx].scenicScoreE++
				break
			}
		}

		// Westward View
		for i := y - 1; i >= 0; i-- {
			if grid[i][x].height < tree.height {
				trees[idx].scenicScoreW++
			} else {
				trees[idx].scenicScoreW++
				break
			}
		}
	}
}
