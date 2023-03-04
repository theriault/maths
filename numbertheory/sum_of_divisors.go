package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// SumOfDivisors returns the sum of the divisors of n.
//
// Time complexity: O(sqrt n)
// Space complexity: O(log n)
//
// https://en.wikipedia.org/wiki/Divisor_function
//
// https://oeis.org/A000203
func SumOfDivisors[A maths.Integer](n A) uint64 {
	if n == 0 {
		return 0
	}
	return primefactorization.NewPrimeFactorization(n).SumOfDivisors()
}
