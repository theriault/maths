package statistics

import (
	"github.com/theriault/maths"
)

// SimpleMovingAverage returns the mean of the previous k datapoints.
//
// If k < 1, returns an empty result.
//
// If k > n, returns a 1-element slice.
//
// https://en.wikipedia.org/wiki/Moving_average#Simple_moving_average
func SimpleMovingAverage[Integer maths.Integer, Number maths.Integer | maths.Float](k Integer, numbers ...Number) []float64 {
	ki := int(k)
	if k < 1 {
		return make([]float64, 0)
	}
	if ki > len(numbers) {
		ki = len(numbers)
	}
	avgs := make([]float64, len(numbers)-ki+1)
	avgs[0] = Mean(numbers[0:ki]...)
	kr := 1 / float64(ki)
	for i := 1; i < len(avgs); i++ {
		avgs[i] = avgs[i-1] + kr*(float64(numbers[ki])-float64(numbers[i-1]))
		ki++
	}

	return avgs
}
