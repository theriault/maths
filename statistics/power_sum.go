package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// PowerSum returns the result of adding together all the arguments raised to
// the pth power.
//
// Nullary invocation returns the additive identity (0).
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Sums_of_powers
func PowerSum[A maths.Integer | maths.Float](X []A, p float64) float64 {
	s := float64(0)
	for _, x := range X {
		s += math.Pow(float64(x), p)
	}
	return s
}
