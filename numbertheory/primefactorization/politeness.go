package primefactorization

// Politeness returns the number of ways n can be expressed as the sum of
// consecutive numbers.
//
// https://en.wikipedia.org/wiki/Polite_number
//
// https://oeis.org/A069283
func (pf *PrimeFactorization) Politeness() uint64 {
	last := uint64(1)
	exponent := uint64(1)
	product := uint64(1)
	for _, p := range pf.primes {
		if p == 2 {
			continue
		}
		if p == last {
			exponent++
		} else {
			product *= exponent
			last = p
			exponent = 2
		}
	}
	product *= exponent
	return product - 1
}
