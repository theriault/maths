package numbertheory

import "github.com/theriault/maths"

// AliquotSum returns the sum of all proper divisors of n (all divisors except n itself).
//
// https://en.wikipedia.org/wiki/Aliquot_sum
//
// https://oeis.org/A001065
func AliquotSum[A maths.Integer](n A) uint64 {
	return SumOfDivisors(n) - uint64(n)
}
