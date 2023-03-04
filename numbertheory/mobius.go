package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// Mobius (Möbius function μ(n)) returns 0 if n is not square-free, -1 if n is square-free and has an odd number of
// prime factors, or 1 if n is square-free and has an even number of prime factors.
//
// Time complexity: O(sqrt n)
// Space complexity: O(log n)
//
// https://en.wikipedia.org/wiki/Mobius_function
//
// https://oeis.org/A008683
func Mobius[A maths.Integer](n A) int8 {
	return primefactorization.NewPrimeFactorization(n).Mobius()
}
