package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using sample input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 875318608908, Solve(SampleInput))
	})

	t.Run("solve correctly using full input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 83551068361379, Solve(FullInput))
	})
}
