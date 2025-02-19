package stack

func Len[T any](slice *[]T) int {
	return len(*slice)
}

func Push[T any](slice *[]T, i T) {
	*slice = append(*slice, i)
}

func Pop[T any](slice *[]T) T {
	s := *slice
	res := s[len(s)-1]
	*slice = s[:len(s)-1]
	return res
}

func PopBack[T any](slice *[]T) T {
	s := *slice
	res := s[0]
	*slice = s[1:]
	return res
}

func PushN[T any](slice *[]T, i ...T) {
	*slice = append(*slice, i...)
}

func PopN[T any](slice *[]T, n int) []T {
	s := *slice
	res := s[len(s)-n:]
	*slice = s[:len(s)-n]
	return res
}
