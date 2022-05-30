package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// Skewness returns the measure of the asymmetry of the probability distribution
// of a real-valued random variable about its mean. It is the 2nd central moment
// divided by the standard deviation raised to the third power.
//
// Zero skew means the values on both sides of the mean balance out overall.
// This can be because the distribution is symmetric, but can also be true for
// an asymmetric distribution where one tail is long and thin, and the other is
// short but fat.
//
// A negative skew means the distribution has a weighter tail on the left side.
// A positive skew means the distribution has a weighter tail on the right side.
//
// See also SampleSkewness.
//
// https://en.wikipedia.org/wiki/Skewness
func Skewness[A maths.Integer | maths.Float](X ...A) float64 {
	n := float64(len(X))
	mean := Mean(X...)
	return PowerSumAround(X, mean, 3) / n / math.Pow(math.Sqrt(PowerSumAround(X, mean, 2)/n), 3)
	// this is cleaner, but is more computationally expensive:
	//return CentralMoment(X, 3) / math.Pow(StandardDeviation(X...), 3)
}
