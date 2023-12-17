// Package m provides common math algorithms.
package m

// GreatestCommonDivisor returns the greatest common divisor of a and b.
func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LeastCommonMultiple returns the least common multiple of a and b.
func LeastCommonMultiple(a, b int) int {
	return a * b / GreatestCommonDivisor(a, b)
}
