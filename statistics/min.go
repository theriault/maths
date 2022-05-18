package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// Min returns the smallest value.
// This function returns the value as the same underlying type.
// Nullary invocation returns the default value of the type.
//
// https://en.wikipedia.org/wiki/Sample_maximum_and_minimum
func Min[A maths.Integer | maths.Float](a ...A) A {
	if len(a) == 0 {
		var a A
		return a
	}
	min := a[0]
	for _, v := range a {
		if math.IsNaN(float64(v)) {
			return A(v)
		}
		if v < min {
			min = v
		}
	}
	return min
}
