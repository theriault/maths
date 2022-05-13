package numbertheory

import "math"

// Euler's totient function counts the positive integers up to a given integer n that are relatively prime to n.
//
// https://en.wikipedia.org/wiki/Euler's_totient_function
// https://oeis.org/A000010
func Totient(n uint64) uint64 {
	return TotientK(n, 1)
}

// Jordan's totient function equals the number of k-tuples of positive integers that are less than or equal to n and
// that together with n form a coprime set of k + 1 integers.
//
// https://en.wikipedia.org/wiki/Jordan's_totient_function
// https://oeis.org/A007434 (k=2)
func TotientK(n uint64, k uint64) uint64 {
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
