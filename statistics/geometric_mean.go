package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// GeometricMean returns the central tendency or typical value of a set of integers/floats by using the product of their
// values.
//
// Returns math.NaN() if:
// - nullary invocation
// - the product of the values is < 0
//
// https://en.wikipedia.org/wiki/Geometric_mean
func GeometricMean[A maths.Integer | maths.Float](numbers ...A) float64 {
	if len(numbers) == 0 {
		return math.NaN()
	}
	product := float64(1)
	for _, n := range numbers {
		if n == 0 {
			return 0
		}
		product *= float64(n)
	}
	if product < 0 {
		return math.NaN()
	}
	return math.Pow(product, 1/float64(len(numbers)))
}
