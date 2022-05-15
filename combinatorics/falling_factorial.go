package combinatorics

import (
	"fmt"

	"github.com/theriault/maths"
)

// FallingFactorial returns the number of ways of choosing an ordered list of length n consisting of distinct elements
// drawn from a collection of size x. It is possible that the result may overflow.
//
// When x = n, this is the same as Factorial(x, n).
//
// https://en.wikipedia.org/wiki/Falling_and_rising_factorials
func FallingFactorial[A maths.Integer, B maths.Integer](x A, n B) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be >= 0")
	}
	if n == 0 {
		return 1, nil
	}
	if float64(n) > float64(x) {
		return 0, nil
	}
	product := uint64(x)
	for i, n := 1, int(n); i < n; i++ {
		product *= (uint64(x) - uint64(i))
	}
	return product, nil
}
