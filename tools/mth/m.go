// Package mth provides common math algorithms.
package mth

import "math"

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

// Sum returns the sum of all [Numeric] elements in s.
func Sum[T Numeric](s []T) T {
	var sum T
	if len(s) == 0 {
		return sum
	}

	for _, v := range s {
		sum += v
	}

	return sum
}

// Numeric constraints a type to be a number which can be used for specific
// math operations.
type Numeric interface {
	int | int32 | int64 | int16 | int8 | float32 | float64
}

func IntAbs(x int) int {
	return int(math.Abs(float64(x)))
}

// PMod returns non-negative solution to x % d.
// -26%7 = 2
// https://www.wolframalpha.com/input?i=-26%257
func PMod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	if d < 0 {
		return x - d
	}
	return x + d
}
