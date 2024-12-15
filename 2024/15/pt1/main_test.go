package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input 1", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 2028, Solve(SampleInput1))
	})

	t.Run("solve correctly using sample input 2", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 10092, Solve(SampleInput2))
	})

	t.Run("solve correctly using full input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 1516281, Solve(FullInput))
	})
}
