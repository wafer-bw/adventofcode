package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, fmt.Sprint(uint64(46)), fmt.Sprint(Solve(SampleInput)))
	})

	// This test is commented out because it takes ~5m to complete.
	// t.Run("solve correctly using full input", func(t *testing.T) {
	// 	t.Parallel()
	// 	require.Equal(t, 84206669, Solve(FullInput))
	// })
}
