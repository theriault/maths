package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// SampleKurtosis
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Kurtosis
func SampleKurtosis[A maths.Integer | maths.Float](X ...A) float64 {
	n := float64(len(X))
	mean := Mean(X...)
	return PowerSumAround(X, mean, 4) / math.Pow(PowerSumAround(X, mean, 2), 2) * ((n * (n + 1) * (n - 1)) / ((n - 2) * (n - 3)))
}
