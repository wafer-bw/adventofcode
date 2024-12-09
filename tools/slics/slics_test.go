package slics_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/adventofcode/tools/slics"
)

func TestSwap(t *testing.T) {
	t.Parallel()

	t.Run("a", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		slics.Swap(s, 0, 4)
		require.Equal(t, []int{5, 2, 3, 4, 1}, s)
	})

	t.Run("b", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		slics.Swap(s, 1, 3)
		require.Equal(t, []int{1, 4, 3, 2, 5}, s)
	})
}

func TestSwapChunk(t *testing.T) {
	t.Parallel()

	t.Run("swap chunks with no overlap", func(t *testing.T) {
		t.Parallel()
		s := []int{1, 2, 3, 4, 5, 6, 7, 8}
		err := slics.SwapChunk(s, 0, 4, 4)
		require.NoError(t, err)
		require.Equal(t, []int{5, 6, 7, 8, 1, 2, 3, 4}, s)
	})

	t.Run("return error if chunks overlap", func(t *testing.T) {
		t.Parallel()
		s := []int{1, 2, 3, 4, 5, 6, 7, 8}
		err := slics.SwapChunk(s, 0, 2, 4)
		require.Error(t, err)
		require.Equal(t, "chunks overlap", err.Error())
	})

	t.Run("return error if chunks out of bounds", func(t *testing.T) {
		t.Parallel()
		s := []int{1, 2, 3, 4, 5, 6, 7, 8}
		err := slics.SwapChunk(s, 0, 6, 4)
		require.Error(t, err)
		require.Equal(t, "chunks out of bounds", err.Error())
	})

}
