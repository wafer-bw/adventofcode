package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input 1", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 1751, Solve(SampleInput1)) // TODO: unsure if answer is correct
	})

	// t.Run("solve correctly using sample input 2", func(t *testing.T) {
	// 	t.Parallel()
	// 	require.Equal(t, 9021, Solve(SampleInput2))
	// })

	// t.Run("solve correctly using full input", func(t *testing.T) {
	// 	t.Parallel()
	// 	require.Equal(t, 1516281, Solve(FullInput))
	// })
}
