package statistics

import (
	"errors"
	"math"

	"github.com/theriault/maths"
)

// WeightedGeneralizedMean
//
// https://en.wikipedia.org/wiki/Generalized_mean
func WeightedGeneralizedMean[A maths.Integer | maths.Float, B maths.Integer | maths.Float](X []A, W []B, p float64) (float64, error) {
	if len(X) == 0 {
		return 0, errors.New("x must have at least 1 element")
	}
	if len(X) != len(W) {
		return 0, errors.New("x and w are not the same lengths")
	}
	if p == 0 {
		sumWeights := float64(0)
		productPowers := float64(1)
		for k, x := range X {
			productPowers *= math.Pow(float64(x), float64(W[k]))
			sumWeights += float64(W[k])
		}
		return math.Pow(productPowers, 1/sumWeights), nil
	}
	// as p approaches negative infinity, the general mean equals the minimum
	if p == math.Inf(-1) {
		min := float64(X[0])
		minWeight := float64(W[0])
		for k, x := range X {
			if math.IsNaN(float64(x)) {
				return math.NaN(), nil
			}
			if float64(x) < min {
				min = float64(x)
				minWeight = float64(W[k])
			}
		}
		return min * minWeight, nil
	}
	// as p approaches positive infinity, the general mean equals the maximum
	if p == math.Inf(+1) {
		max := float64(X[0])
		maxWeight := float64(W[0])
		for k, x := range X {
			if math.IsNaN(float64(x)) {
				return math.NaN(), nil
			}
			if float64(x) > max {
				max = float64(x)
				maxWeight = float64(W[k])
			}
		}
		return max * maxWeight, nil
	}
	sumWeights := float64(0)
	sumPowers := float64(0)
	for k, x := range X {
		sumPowers += math.Pow(float64(W[k])*float64(x), p)
		sumWeights += float64(W[k])
	}
	return math.Pow(sumPowers/sumWeights, 1/p), nil
}

// WeightedPowerMean is an alias for GeneralizedWeightedMean
func WeightedPowerMean[A maths.Integer | maths.Float, B maths.Integer | maths.Float](X []A, W []B, p float64) (float64, error) {
	return WeightedGeneralizedMean(X, W, p)
}

// WeightedGeometricMean
//
// https://en.wikipedia.org/wiki/Weighted_geometric_mean
func WeightedGeometricMean[A maths.Integer | maths.Float, B maths.Integer | maths.Float](X []A, W []B) (float64, error) {
	return WeightedPowerMean(X, W, 0)
}

// WeightedMean
//
// https://en.wikipedia.org/wiki/Weighted_arithmetic_mean
func WeightedMean[A maths.Integer | maths.Float, B maths.Integer | maths.Float](X []A, W []B) (float64, error) {
	return WeightedPowerMean(X, W, 1)
}

// WeightedHarmonicMean
//
// https://en.wikipedia.org/wiki/Harmonic_mean#Weighted_harmonic_mean
func WeightedHarmonicMean[A maths.Integer | maths.Float, B maths.Integer | maths.Float](X []A, W []B) (float64, error) {
	return WeightedPowerMean(X, W, -1)
}
