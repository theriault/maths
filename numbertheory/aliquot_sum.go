package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// AliquotSum returns the sum of all proper divisors of n (all divisors except n itself).
//
// Time complexity: O(sqrt n)
// Space complexity: O(log n)
//
// https://en.wikipedia.org/wiki/Aliquot_sum
//
// https://oeis.org/A001065
func AliquotSum[A maths.Integer](n A) uint64 {
	return primefactorization.NewPrimeFactorization(n).AliquotSum()
}
