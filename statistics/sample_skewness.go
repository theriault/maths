package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// SampleSkewness
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Skewness
func SampleSkewness[A maths.Integer | maths.Float](X ...A) float64 {
	n := float64(len(X))
	mean := Mean(X...)
	stddev := SampleStandardDeviation(X...)
	return PowerSumAround(X, mean, 3) / math.Pow(stddev, 3) * (n / ((n - 1) * (n - 2)))
}
