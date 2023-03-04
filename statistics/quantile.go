package statistics

import (
	"fmt"
	"math"
	"sort"

	"github.com/theriault/maths"
)

// Quantile returns cut points dividing the range of a probability distribution into continuous intervals with equal
// probabilities. X contains a slice of any primitive int or float type. q indicates the number of q-iles to return.
//
// Time complexity: best O(n + m) if X is already sorted - worst O(n log n) if X needs to be sorted
// Space complexity: O(n + m) - where m is the number of quantiles
//
// https://en.wikipedia.org/wiki/Quantile
func Quantile[A maths.Integer | maths.Float](X []A, q int) ([]float64, error) {
	// this will incur a cost of O(n log n) if X wasn't already sorted
	sortedX := maths.Coerce(X, 0.0)
	if !sort.Float64sAreSorted(sortedX) {
		sort.Float64s(sortedX)
	}
	return QuantilePresorted(sortedX, q)
}

// QuantilePresorted returns cut points dividing the range of a probability distribution into continuous intervals with
// equal probabilities. q indicates the number of q-iles to return.
//
// WARNING: This function assumes X is sorted already. If X is not sorted, you will get back an incorrect result.
//
// Time complexity: O(m) - where m is the number of quantiles
// Space complexity: O(m) - where m is the number of quantiles
//
// https://en.wikipedia.org/wiki/Quantile
func QuantilePresorted(X []float64, q int) ([]float64, error) {
	l := len(X)
	if q <= 0 {
		return nil, fmt.Errorf("q must be >= 1: %v", q)
	}
	if l < 1 {
		return nil, fmt.Errorf("insufficient data to produce %v-iles - need at least 1 number", q)
	}

	quantiles := make([]float64, q-1)
	for k := range quantiles {
		i, f := math.Modf((float64(k) + 1) / float64(q) * (float64(l) - 1))
		w := int(i)
		if f == 0 {
			quantiles[k] = X[w]
		} else {
			quantiles[k] = X[w] + (X[w+1]-X[w])*f
		}
	}
	return quantiles, nil
}

// Tertile is an alias for Quantile(X, 3)
func Tertile[A maths.Integer | maths.Float](X []A) ([]float64, error) {
	return Quantile(X, 3)
}

// Quartile is an alias for Quantile(X, 4)
func Quartile[A maths.Integer | maths.Float](X []A) ([]float64, error) {
	return Quantile(X, 4)
}

// Percentile is an alias for Quantile(X, 100)
func Percentile[A maths.Integer | maths.Float](X []A) ([]float64, error) {
	return Quantile(X, 100)
}
