package statistics

import (
	"github.com/theriault/maths"
)

// CentralMoment returns the expected value of a specified integer power of the
// deviation of the elements of X from the mean.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Central_moment
func CentralMoment[A maths.Integer | maths.Float](X []A, p int) float64 {
	return PowerSumAround(X, Mean(X...), p) / float64(len(X))
}
