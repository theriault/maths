package statistics

import (
	"github.com/theriault/maths"
)

// ExcessSampleKurtosis
//
// https://en.wikipedia.org/wiki/Kurtosis
func ExcessSampleKurtosis[A maths.Integer | maths.Float](X ...A) float64 {
	n := float64(len(X))
	return ((n+1)*(Kurtosis(X...)-3) + 6) * (n - 1) / ((n - 2) * (n - 3))
}
