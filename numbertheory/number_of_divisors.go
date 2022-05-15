package numbertheory

import "github.com/theriault/maths"

// NumberOfDivisors returns the number of divisors of n.
//
// https://en.wikipedia.org/wiki/Divisor_function
//
// https://oeis.org/A000005
func NumberOfDivisors[A maths.Integer](n A) uint64 {
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
