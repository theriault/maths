package statistics

import (
	"github.com/theriault/maths"
)

// Sum returns the result of adding together all the arguments.
// Nullary invocation returns the additive identity (0).
//
// https://en.wikipedia.org/wiki/Summation
func Sum[A maths.Integer | maths.Float](numbers ...A) float64 {
	sum := float64(0)
	for _, v := range numbers {
		sum += float64(v)
	}
	return sum
}
