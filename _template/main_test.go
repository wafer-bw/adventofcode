package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

func TestSolve(t *testing.T) {
	require.Equal(t, -1, solve(reader.Read("../inputs/XXXX-XX-sample.txt")))
}
