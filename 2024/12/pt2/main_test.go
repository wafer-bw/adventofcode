package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input 1", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 80, Solve(SampleInput1))
	})

	t.Run("solve correctly using sample input 2", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 436, Solve(SampleInput2))
	})

	t.Run("solve correctly using sample input 3", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 368, Solve(SampleInput3))
	})

	t.Run("solve correctly using sample input 4", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 236, Solve(SampleInput4))
	})

	t.Run("solve correctly using full input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 897062, Solve(FullInput))
	})
}
