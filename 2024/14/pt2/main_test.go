package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed input-solution.txt
var expect string

func TestSolve(t *testing.T) {
	t.Parallel()

	t.Run("solve correctly using full input", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, expect, Solve(FullInput, 7371))
	})
}
