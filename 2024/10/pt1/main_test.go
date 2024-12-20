package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input 1", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 1, Solve(SampleInput1))
	})

	t.Run("solve correctly using sample input 2", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 2, Solve(SampleInput2))
	})

	t.Run("solve correctly using sample input 3", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 4, Solve(SampleInput3))
	})

	t.Run("solve correctly using sample input 4", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 3, Solve(SampleInput4))
	})

	t.Run("solve correctly using sample input 5", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 36, Solve(SampleInput5))
	})

	t.Run("solve correctly using full input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 472, Solve(FullInput))
	})
}
