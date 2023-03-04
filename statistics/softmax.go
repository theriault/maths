package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// GeneralizedSoftmax returns a softmax of a given vector of float64 adjusted by some temperature.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
//
// https://en.wikipedia.org/wiki/Softmax_function
func GeneralizedSoftmax[Number maths.Integer | maths.Float](X []Number, temperature float64) []float64 {
	return MutableGeneralizedSoftmax(maths.Coerce(X, 0.0), temperature)
}

// Softmax returns the softmax vector of a given vector of numbers, which are coerced to float64.
// Softmax applies the exponential function to each element of X and then normalizes the resulting values by dividing
// by the sum of all exponentials. This normalization guarantees that the sum of the components of the output vector is
// always equal to 1.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
//
// https://en.wikipedia.org/wiki/Softmax_function
func Softmax[Number maths.Integer | maths.Float](X []Number) []float64 {
	return MutableSoftmax(maths.Coerce(X, 0.0))
}

// MutableSoftmax modifies a vector of float64s to be their softmax values.
//
// Time Complexity: O(n) - n is the size of X
// Space Complexity: O(1) - no extra space required - modifies X in place
//
// https://en.wikipedia.org/wiki/Softmax_function
func MutableSoftmax(X []float64) []float64 {
	return MutableGeneralizedSoftmax(X, 1.0)
}

// MutableGeneralizedSoftmax modifies a vector of float64s to be their softmax values using the provided temperature.
// The temperature parameter adjusts the smoothness of the softmax function. A higher temperature results in a smoother,
// more uniform distribution of probabilities across the vector. A lower temperature results in a more peaky
// distribution.
//
// Time Complexity: O(n) - n is the size of X
// Space Complexity: O(1) - no extra space required - modifies X in place
//
// https://en.wikipedia.org/wiki/Softmax_function
func MutableGeneralizedSoftmax(X []float64, temperature float64) []float64 {
	// Calculate base-e exponential of X and get its summation
	sum := 0.0
	for i, x := range X {
		X[i] = math.Exp(x / temperature)
		sum += X[i]
	}

	// Divide the base-e exponential of X by the summation across X
	for i := range X {
		X[i] /= sum
	}

	return X
}
