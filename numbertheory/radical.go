package numbertheory

// Radical is the product of the distinct primes dividing n.
//
// https://en.wikipedia.org/wiki/Radical_of_an_integer
// https://oeis.org/A007947
func Radical(n uint64) uint64 {
	factors := PrimeFactorization(n)
	radical := uint64(1)
	last := uint64(1)
	for _, p := range factors {
		if p == last {
			continue
		}
		radical *= p
		last = p
	}
	return radical
}
