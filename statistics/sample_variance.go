package statistics

import (
	"github.com/theriault/maths"
)

// SampleVariance returns the expectation of the squared deviation of a random
// variable from its mean after applying Bessel's correction to account for bias
// when the population mean is unknown.
//
// https://en.wikipedia.org/wiki/Variance#Sample_variance
//
// https://en.wikipedia.org/wiki/Bessel%27s_correction
func SampleVariance[A maths.Integer | maths.Float](X ...A) float64 {
	n := float64(len(X))
	return CentralMoment(X, 2) * (n / (n - 1))
}
