package numbertheory

// Möbius function μ(n) returns 0 if n is not square-free, -1 if n is square-free and has an odd number of prime
// factors, or 1 if n is square-free and has an even number of prime factors.
//
// https://en.wikipedia.org/wiki/Mobius_function
// https://oeis.org/A008683
func Mobius(n uint64) int8 {
	factors := PrimeFactorization(n)
	last := uint64(1)
	for _, p := range factors {
		if p == last {
			return 0
		}
		last = p
	}
	if len(factors)&1 == 1 {
		return -1
	} else {
		return 1
	}
}
