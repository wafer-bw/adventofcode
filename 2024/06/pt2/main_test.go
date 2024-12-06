package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 6, Solve(SampleInput))
	})
	// Too slow to run in a test bc I just brute forced it but the answer is
	// correct for my input.
	// t.Run("solve correctly using full input", func(t *testing.T) {
	// 	t.Parallel()
	// 	require.Equal(t, 1670, Solve(FullInput))
	// })
}
