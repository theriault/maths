package primefactorization

import (
	"math"
)

// Totient (Euler's totient) returns the number of positive integers up to a given integer n that are relatively prime
// to n.
//
// Time complexity: O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Euler's_totient_function
//
// https://oeis.org/A000010
func (pf *PrimeFactorization) Totient() uint64 {
	return pf.TotientK(1)
}

// TotientK (Jordan's totient) returns the number of k-tuples of positive integers that are less than or equal to n and
// that together with n form a coprime set of k + 1 integers.
//
// Time complexity: O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Jordan's_totient_function
//
// https://oeis.org/A007434 (k=2)
func (pf *PrimeFactorization) TotientK(k int64) uint64 {
	J := math.Pow(float64(pf.n), float64(k))
	last := uint64(0)
	for _, p := range pf.primes {
		if p != last {
			J *= 1 - 1/math.Pow(float64(p), float64(k))
		}
		last = p
	}

	return uint64(J)
}
