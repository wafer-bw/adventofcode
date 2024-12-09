package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 2858, Solve(SampleInput))
	})
	t.Run("solve correctly using full input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 6363268339304, Solve(FullInput))
	})
}
