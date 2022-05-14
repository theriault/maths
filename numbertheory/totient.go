package numbertheory

import "math"

// Totient (Euler's totient) returns the number of positive integers up to a given integer n that are relatively prime
// to n.
//
// https://en.wikipedia.org/wiki/Euler's_totient_function
// https://oeis.org/A000010
func Totient[A Integer](n A) uint64 {
	return TotientK(n, 1)
}

// TotientK (Jordan's totient) returns the number of k-tuples of positive integers that are less than or equal to n and
// that together with n form a coprime set of k + 1 integers.
//
// https://en.wikipedia.org/wiki/Jordan's_totient_function
//
// https://oeis.org/A007434 (k=2)
func TotientK[A Integer, B Integer](n A, k B) uint64 {
	J := math.Pow(float64(n), float64(k))
	primes := PrimeFactorization(n)

	last := uint64(0)
	for _, p := range primes {
		if p != last {
			J *= 1 - 1/math.Pow(float64(p), float64(k))
		}
		last = p
	}

	return uint64(J)
}
