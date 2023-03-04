package primefactorization

// NumberOfDivisors returns the number of divisors of n.
//
// Time complexity: O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Divisor_function
//
// https://oeis.org/A000005
func (pf *PrimeFactorization) NumberOfDivisors() uint64 {
	product := uint64(1)

	last := uint64(0)
	exponent := uint64(0)
	for _, p := range pf.primes {
		if p == last {
			exponent += 1
		} else {
			if last != 0 {
				product *= exponent + 1
			}
			exponent = 1
			last = p
		}
	}
	if last != 0 {
		product *= exponent + 1
	}

	return product
}
