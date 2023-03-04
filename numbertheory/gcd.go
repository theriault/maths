package numbertheory

import "github.com/theriault/maths"

// GCD returns the largest positive integer that divides a and b.
//
// Time complexity: O(log n) - where n is whichever number is smaller
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Greatest_common_divisor
func GCD[A maths.Integer](a, b A) A {
	if a == 0 && b == 0 {
		return 1
	}
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		b, a = a%b, b
	}
	return a
}
