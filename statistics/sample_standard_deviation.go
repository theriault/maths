package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// SampleStandardDeviation returns a measure of the amount of variation or
// dispersion of a set of values from its mean/expected value after applying
// Bessel's correction to account for bias when the population mean is unknown.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Standard_deviation
func SampleStandardDeviation[A maths.Integer | maths.Float](Y ...A) float64 {
	return math.Sqrt(SampleVariance(Y...))
}
