package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// Max returns the largest value.
// This function returns the value as the same underlying type.
// Nullary invocation returns the default value of the type.
// Any NaN returns NaN.
//
// https://en.wikipedia.org/wiki/Sample_maximum_and_minimum
func Max[A maths.Integer | maths.Float](a ...A) A {
	l := len(a)
	if l == 0 {
		return 0
	}

	max := a[0]
	for _, v := range a {
		if math.IsNaN(float64(v)) {
			return A(v)
		}
		if v > max {
			max = v
		}
	}
	return max
}
