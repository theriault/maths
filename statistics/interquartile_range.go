package statistics

import (
	"github.com/theriault/maths"
)

// InterquartileRange returns the difference between the 75th and 25th percentiles.
//
// Time Complexity: best O(n) if X is sorted. worst O(n log n) if X is not sorted.
// Space complexity: O(n)
//
// https://en.wikipedia.org/wiki/Interquartile_range
func InterquartileRange[A maths.Integer | maths.Float](X ...A) (float64, error) {
	Q, err := Quantile(X, 4)
	if err != nil {
		return 0, err
	}

	return Q[2] - Q[0], nil
}
