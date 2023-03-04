package numbertheory

import (
	"github.com/theriault/maths"
	"github.com/theriault/maths/numbertheory/primefactorization"
)

// Totient (Euler's totient) returns the number of positive integers up to a given integer n that are relatively prime
// to n.
//
// Time complexity: O(sqrt n)
// Space complexity: O(log n)
//
// https://en.wikipedia.org/wiki/Euler's_totient_function
//
// https://oeis.org/A000010
func Totient[A maths.Integer](n A) uint64 {
	return primefactorization.NewPrimeFactorization(n).Totient()
}

// TotientK (Jordan's totient) returns the number of k-tuples of positive integers that are less than or equal to n and
// that together with n form a coprime set of k + 1 integers.
//
// Time complexity: O(sqrt n)
// Space complexity: O(log n)
//
// https://en.wikipedia.org/wiki/Jordan's_totient_function
//
// https://oeis.org/A007434 (k=2)
func TotientK[A maths.Integer, B maths.Integer](n A, k B) uint64 {
	return primefactorization.NewPrimeFactorization(n).TotientK(int64(k))
}
