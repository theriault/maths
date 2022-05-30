package statistics

import (
	"github.com/theriault/maths"
)

// Sum returns the result of adding together all the arguments.
// Nullary invocation returns the additive identity (0).
//
// https://en.wikipedia.org/wiki/Summation
func Sum[A maths.Integer | maths.Float](X ...A) float64 {
	s := float64(0)
	for _, x := range X {
		s += float64(x)
	}
	return s
}
