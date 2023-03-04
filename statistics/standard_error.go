package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// StandardError returns a measure of the dispersion of sample means around the
// population mean.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Standard_error
func StandardError[A maths.Integer | maths.Float](Y ...A) float64 {
	return StandardDeviation(Y...) / math.Sqrt(float64(len(Y)))
}
