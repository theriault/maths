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
// Time complexity: best O(n) if X is sorted. worst O(n log n) if X needs to be sorted. If X is guaranteed to be sorted,
// 					it is recommended that you just call QuantilePresorted(X, 2)
// Space complexity: O(n)
//
// https://en.wikipedia.org/wiki/Median
func Median[Number maths.Integer | maths.Float](X ...Number) float64 {
	m, err := Quantile(X, 2)
	if err != nil {
		return math.NaN()
	}
	return m[0]
}
