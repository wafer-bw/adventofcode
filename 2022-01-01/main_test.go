package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/adventofcode/tools/reader"
)

func TestSolve(t *testing.T) {
	require.Equal(t, 24000, solve(reader.Read("../inputs/2022-01-sample.txt")))
}
