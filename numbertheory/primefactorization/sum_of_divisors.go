package primefactorization

// SumOfDivisors returns the sum of the divisors of n.
//
// Time complexity: O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Divisor_function
//
// https://oeis.org/A000203
func (pf *PrimeFactorization) SumOfDivisors() uint64 {
	last := uint64(0)
	product := uint64(1)
	primePower := uint64(1)
	primePowerSum := uint64(1)
	for _, p := range pf.primes {
		if p == last {
			primePower *= p
			primePowerSum += primePower
		} else {
			product *= primePowerSum
			last = p
			primePower = p
			primePowerSum = 1 + p
		}
	}
	product *= primePowerSum

	return product
}
