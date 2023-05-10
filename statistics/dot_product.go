package statistics

import "github.com/theriault/maths"

// DotProduct calculates the dot product of two slices of the same type.
//
// If X and Y are not the same length, it will return the dot product of the smaller sized array.
//
// Time complexity:  O(min(|X|, |Y|)) - we iterate over X and Y up to the length of the smaller array
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Dot_product
func DotProduct[T maths.Integer | maths.Float](X, Y []T) T {
	l := len(X)
	if len(Y) < l {
		l = len(Y)
	}
	sum := T(0)
	for i := 0; i < l; i++ {
		sum += X[i] * Y[i]
	}
	return sum
}
