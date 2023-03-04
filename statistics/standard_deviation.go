package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// StandardDeviation returns a measure of the amount of variation or dispersion
// of a set of values from its mean/expected value.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Standard_deviation
func StandardDeviation[A maths.Integer | maths.Float](Y ...A) float64 {
	return math.Sqrt(Variance(Y...))
}
