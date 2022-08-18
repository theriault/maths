package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// Radical returns the product of the distinct primes dividing n.
//
// https://en.wikipedia.org/wiki/Radical_of_an_integer
//
// https://oeis.org/A007947
func Radical[A maths.Integer](n A) uint64 {
	return primefactorization.NewPrimeFactorization(n).Radical()
}
