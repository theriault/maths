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
	l := float64(len(numbers))
	if l == 0 {
		return math.NaN()
	}
	sum := float64(0)
	for _, n := range numbers {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return math.NaN()
		}
		sum += math.Log2(float64(n))
	}
	return math.Exp2(1 / l * sum)
}
