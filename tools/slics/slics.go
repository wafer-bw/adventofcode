package slics

import "fmt"

func Swap[T any](s []T, a, b int) {
	// panics if indices out of bounds
	s[a], s[b] = s[b], s[a]
}

func SwapChunk[T any](s []T, a, b int, chunkSize int) error {
	// panics if indices out of bounds
	if a+chunkSize > len(s) || b+chunkSize > len(s) {
		return fmt.Errorf("chunks out of bounds")
	} else if a+chunkSize > b {
		return fmt.Errorf("chunks overlap")
	}

	for k := 0; k < chunkSize; k++ {
		s[a+k], s[b+k] = s[b+k], s[a+k]
	}

	return nil
}
