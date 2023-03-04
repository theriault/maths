package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// Mean returns the average (arithmetic mean) of one or more integers or floats.
//
// Nullary invocation returns math.NaN()
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Mean#Arithmetic_mean_(AM)
func Mean[A maths.Integer | maths.Float](X ...A) float64 {
	if len(X) == 0 {
		return math.NaN()
	}
	return float64(Sum(X...)) / float64(len(X))
}
