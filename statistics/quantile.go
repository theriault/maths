package statistics

import (
	"fmt"
	"math"
	"sort"

	"github.com/theriault/maths"
)

// Quantile returns cut points dividing the range of a probability distribution into continuous intervals with equal
// probabilities. The first argument contains the data. The second argument q indicates the number of q-iles to return.
//
// https://en.wikipedia.org/wiki/Quantile
func Quantile[A maths.Integer | maths.Float](numbers []A, q int) ([]float64, error) {
	l := len(numbers)
	if q <= 0 {
		return nil, fmt.Errorf("q must be >= 1: %v", q)
	}
	if l < 1 {
		return nil, fmt.Errorf("insufficient data to produce %v-iles - need at least 1 number", q)
	}

	sortedNumbers := make([]float64, l)
	for k, v := range numbers {
		sortedNumbers[k] = float64(v)
	}
	sort.Float64s(sortedNumbers)

	quantiles := make([]float64, q-1)
	for k := range quantiles {
		i, f := math.Modf((float64(k) + 1) / float64(q) * (float64(l) - 1))
		w := int(i)
		if f == 0 {
			quantiles[k] = sortedNumbers[w]
		} else {
			quantiles[k] = sortedNumbers[w] + (sortedNumbers[w+1]-sortedNumbers[w])*f
		}
	}
	return quantiles, nil
}
