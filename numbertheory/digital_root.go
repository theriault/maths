package numbertheory

import (
	"github.com/theriault/maths"
)

// DigitalRoot returns a single digit after iteratively summing the digits of n in a given radix/base.
//
// https://en.wikipedia.org/wiki/Digital_root
func DigitalRoot[A maths.Integer](n, b A) A {
	// any b<1 will return 0. any n<1 will return 0. caller can verify it was an error by checking if n>0
	if b < 1 || n < 1 {
		return 0
	}
	// base 1 will be treated as the unary numeral system. because we already verified n >= 1, return 1 for unary/binary
	if b <= 2 {
		return 1
	}

	x := n % (b - 1)
	if x == 0 {
		return b - 1
	}
	return x
}
