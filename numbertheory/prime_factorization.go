package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// PrimeFactorization returns the prime factorization of n.
//
// https://en.wikipedia.org/wiki/Integer_factorization
func PrimeFactorization[A maths.Integer](n A) []uint64 {
	return primefactorization.NewPrimeFactorization(n).Primes()
}
