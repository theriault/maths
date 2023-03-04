package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// Politeness returns the number of ways n can be expressed as the sum of
// consecutive numbers.
//
// Time complexity: O(sqrt n)
// Space complexity: O(log n)
//
// https://en.wikipedia.org/wiki/Polite_number
//
// https://oeis.org/A069283
func Politeness[A maths.Integer](n A) uint64 {
	return primefactorization.NewPrimeFactorization(n).Politeness()
}
