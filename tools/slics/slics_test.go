package slics_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/adventofcode/tools/slics"
)

func TestSwap(t *testing.T) {
	t.Parallel()
	s := []int{1, 2, 3, 4, 5}
	slics.Swap(s, 0, 4)
	require.Equal(t, []int{5, 2, 3, 4, 1}, s)

	slics.Swap(s, 1, 3)
	require.Equal(t, []int{5, 4, 3, 2, 1}, s)
}
