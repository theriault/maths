package numbertheory

// Number-of-divisors function counts the number of divisors of an integer.
//
// https://en.wikipedia.org/wiki/Divisor_function
// https://oeis.org/A000005
func NumberOfDivisors(n uint64) uint64 {
	factors := PrimeFactorization(n)

	product := uint64(1)

	last := uint64(0)
	exponent := uint64(0)
	for _, p := range factors {
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
