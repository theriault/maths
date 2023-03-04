package statistics

import (
	"github.com/theriault/maths"
)

// RootMeanSquare
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Root_mean_square
func RootMeanSquare[A maths.Integer | maths.Float](X ...A) float64 {
	return GeneralizedMean(X, 2)
}
