package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// PowerSumAround returns the result of summing the deviations of the elements
// of X from y raised to the power of p.
//
// If X has no elements, returns the additive identity (0).
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Sums_of_powers
func PowerSumAround[A maths.Integer | maths.Float](X []A, y float64, p int) float64 {
	s := float64(0)
	for _, x := range X {
		s += math.Pow(float64(x)-y, float64(p))
	}
	return s
}
