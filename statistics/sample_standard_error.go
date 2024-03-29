package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// SampleStandardError
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Standard_error
func SampleStandardError[A maths.Integer | maths.Float](Y ...A) float64 {
	return SampleStandardDeviation(Y...) / math.Sqrt(float64(len(Y)))
}
