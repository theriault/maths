package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// Mean returns the average (arithmetic mean) of one or more integers or floats.
//
// Nullary invocation returns math.NaN()
//
// https://en.wikipedia.org/wiki/Mean#Arithmetic_mean_(AM)
func Mean[A maths.Integer | maths.Float](numbers ...A) float64 {
	if len(numbers) == 0 {
		return math.NaN()
	}
	return Sum(numbers...) / float64(len(numbers))
}
