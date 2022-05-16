package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// HarmonicMean returns the reciprocal of the arithmetic mean of the reciprocals of the given set of observations.
//
// Returns math.NaN() if:
// - nullary invocation
// - any value <= 0
//
// https://en.wikipedia.org/wiki/Harmonic_mean
func HarmonicMean[A maths.Integer | maths.Float](numbers ...A) float64 {
	if len(numbers) == 0 {
		return math.NaN()
	}
	sum := float64(0)
	for _, n := range numbers {
		if n <= 0 {
			return math.NaN()
		}
		sum += 1 / float64(n)
	}
	return float64(len(numbers)) / sum
}
