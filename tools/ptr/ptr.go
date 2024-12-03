package ptr

// To returns a pointer to the passed value.
func To[T any](t T) *T {
	return &t
}
