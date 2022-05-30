package statistics

import (
	"github.com/theriault/maths"
)

// Variance returns the expectation of the squared deviation of a random
// variable from its mean. Variance is a measure of dispersion, meaning it is a
// measure of how far a set of numbers is spread out from their average value.
//
// https://en.wikipedia.org/wiki/Variance#Population_variance
func Variance[A maths.Integer | maths.Float](X ...A) float64 {
	return CentralMoment(X, 2)
}
