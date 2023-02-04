package alphanum

func ToNum(r rune) int {
	return int(r) - 96
}

func ToChar(n int) rune {
	return rune('a' - 1 + n)
}
