package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/adventofcode/tools/pather"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 4, solve(reader.Read(pather.Path(puzzleID, true, true))))
	})
	t.Run("solve correctly using full input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 926, solve(reader.Read(pather.Path(puzzleID, false, true))))
	})
}
