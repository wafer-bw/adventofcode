package stack_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/adventofcode/tools/stack"
)

func TestPush(t *testing.T) {
	s := []int{1, 2, 3}
	stack.Push(&s, 4)
	require.Equal(t, 4, s[3])
	require.Len(t, s, 4)
}

func TestPop(t *testing.T) {
	s := []int{1, 2, 3}
	res := stack.Pop(&s)
	require.Equal(t, 3, res)
	require.Len(t, s, 2)
}
