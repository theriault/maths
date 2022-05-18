package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// Median returns the value separating the higher half from the lower half of a data sample, a population, or a
// probability distribution.
//
// Nullary invocation returns math.NaN()
//
// https://en.wikipedia.org/wiki/Median
func Median[A maths.Integer | maths.Float](numbers ...A) float64 {
	m, err := Quantile(numbers, 2)
	if err != nil {
		return math.NaN()
	}
	return m[0]
}
