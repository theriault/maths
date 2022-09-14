package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// NumberOfDivisors returns the number of divisors of n.
//
// https://en.wikipedia.org/wiki/Divisor_function
//
// https://oeis.org/A000005
func NumberOfDivisors[A maths.Integer](n A) uint64 {
	return primefactorization.NewPrimeFactorization(n).NumberOfDivisors()
}
