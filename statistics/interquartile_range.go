package statistics

import (
	"github.com/theriault/maths"
)

// InterquartileRange returns the difference between the 75th and 25th percentiles.
//
// https://en.wikipedia.org/wiki/Interquartile_range
func InterquartileRange[A maths.Integer | maths.Float](X ...A) (float64, error) {
	Q, err := Quantile(X, 4)
	if err != nil {
		return 0, err
	}

	return Q[2] - Q[0], nil
}
