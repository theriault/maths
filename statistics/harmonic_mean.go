package statistics

import (
	"github.com/theriault/maths"
)

// HarmonicMean returns the reciprocal of the arithmetic mean of the reciprocals of the given set of observations.
//
// Returns math.NaN() if:
// - nullary invocation
// - any value <= 0
//
// https://en.wikipedia.org/wiki/Harmonic_mean
func HarmonicMean[A maths.Integer | maths.Float](X ...A) float64 {
	return PowerMean(X, -1)
}
