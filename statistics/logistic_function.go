package statistics

import (
	"math"

	"github.com/theriault/maths"
)

// LogisticFunction
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Logistic_function
func LogisticFunction[A maths.Integer | maths.Float](max A, midpoint A, growthRate A) func(x A) float64 {
	L := float64(max)
	x0 := float64(midpoint)
	k := float64(growthRate)
	return func(x A) float64 {
		return L / (1 + math.Exp(-k*(float64(x)-x0)))
	}
}
