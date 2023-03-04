package numbertheory

import (
	"math"

	"github.com/theriault/maths"
)

// LCM returns the smallest positive integer that is divisible by both a and b.
//
// Time complexity: O(log n) - where n is whichever number is smaller
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Least_common_multiple
func LCM[A maths.Integer](a, b A) A {
	if a == 0 && b == 0 {
		return 0
	}
	return A(math.Abs(float64(a)) * (math.Abs(float64(b)) / float64(GCD(a, b))))
}
