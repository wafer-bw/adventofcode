package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
	// "github.com/stretchr/testify/require"
	// "github.com/wafer-bw/adventofcode/tools/pather"
	// "github.com/wafer-bw/adventofcode/tools/reader"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 29, solve(reader.Read(pather.Path(puzzleID, true, true))))
	})
	// 402 is the correct answer for my input but it takes 20 minutes to run LOL
	// When I have more time to try to learn A* I'll come back to this.
	// t.Run("solve correctly using full input", func(t *testing.T) {
	// 	t.Parallel()
	// 	require.Equal(t, 402, solve(reader.Read(pather.Path(puzzleID, false, true))))
	// })
}
