package statistics

import (
	"github.com/theriault/maths"
)

// Range returns the the difference between the max and min.
// This function returns the same underlying type.
// Nullary invocation returns the default value of the type.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Range_(statistics)
func Range[A maths.Integer | maths.Float](a ...A) A {
	return Max(a...) - Min(a...)
}
