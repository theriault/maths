package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// Kurtosis returns the "tailedness" of the probability distribution of a
// real-valued random variable. It is the 4th central moment divided by the
// standard deviation raised to the 4th power.
//
// See also SampleKurtosis and ExcessSampleKurtosis.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Kurtosis
func Kurtosis[A maths.Integer | maths.Float](X ...A) float64 {
	n := float64(len(X))
	mean := Mean(X...)
	return PowerSumAround(X, mean, 4) / n / math.Pow(math.Sqrt(PowerSumAround(X, mean, 2)/n), 4)
	// this is cleaner, but is more computationally expensive:
	// return CentralMoment(X, 4) / math.Pow(StandardDeviation(X...), 4)
}
