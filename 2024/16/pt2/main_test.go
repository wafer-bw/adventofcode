package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	filter := []int{}
	for i, c := range Cases {
		if slices.Contains(filter, i) {
			continue
		}

		t.Run(fmt.Sprintf("solve correctly using input %d", i), func(t *testing.T) {
			require.Equal(t, c.Answer, Solve(c.Input))
		})
	}
}
