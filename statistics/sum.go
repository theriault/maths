package statistics

import (
	"github.com/theriault/maths"
)

// Sum returns the result of adding together all the arguments. The result
// will be the same type as the arguments.
//
// Nullary invocation returns the additive identity (0).
//
// The underlying type may overflow.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Summation
func Sum[A maths.Integer | maths.Float](X ...A) A {
	s := A(0)
	for _, x := range X {
		s += x
	}
	return s
}
