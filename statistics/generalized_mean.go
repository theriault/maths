package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// GeneralizedMean
//
// https://en.wikipedia.org/wiki/Generalized_mean
func GeneralizedMean[A maths.Integer | maths.Float](X []A, p float64) float64 {
	// as p approaches 0, the general mean equals the geometric mean
	if p == 0 {
		return GeometricMean(X...)
	}
	// as p approaches negative infinity, the general mean equals the minimum
	if p == math.Inf(-1) {
		return float64(Min(X...))
	}
	// as p approaches positive infinity, the general mean equals the maximum
	if p == math.Inf(+1) {
		return float64(Max(X...))
	}
	if len(X) == 0 {
		return math.NaN()
	}
	sum := PowerSum(X, p)
	if sum == 0 {
		return math.NaN()
	}
	return math.Pow(sum/float64(len(X)), float64(1)/p)
}

// PowerMean is an alias for GeneralizedMean
func PowerMean[A maths.Integer | maths.Float](X []A, p float64) float64 {
	return GeneralizedMean(X, p)
}
