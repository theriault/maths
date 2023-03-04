package primefactorization

// Radical returns the product of the distinct primes dividing n.
//
// Time complexity: O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Radical_of_an_integer
//
// https://oeis.org/A007947
func (pf *PrimeFactorization) Radical() uint64 {
	radical := uint64(1)
	last := uint64(1)
	for _, p := range pf.primes {
		if p == last {
			continue
		}
		radical *= p
		last = p
	}
	return radical
}
