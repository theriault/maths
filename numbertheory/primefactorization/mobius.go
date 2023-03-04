package primefactorization

// Mobius (Möbius function μ(n)) returns 0 if n is not square-free, -1 if n is square-free and has an odd number of
// prime factors, or 1 if n is square-free and has an even number of prime factors.
//
// Time complexity: O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Mobius_function
//
// https://oeis.org/A008683
func (pf *PrimeFactorization) Mobius() int8 {
	last := uint64(1)
	for _, p := range pf.primes {
		if p == last {
			return 0
		}
		last = p
	}
	if len(pf.primes)&1 == 1 {
		return -1
	} else {
		return 1
	}
}
