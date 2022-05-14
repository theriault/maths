package numbertheory

// SumOfDivisors returns the sum of the divisors of n.
//
// https://en.wikipedia.org/wiki/Divisor_function
//
// https://oeis.org/A000203
func SumOfDivisors[A Integer](n A) uint64 {
	if n == 0 {
		return 0
	}
	factors := PrimeFactorization(n)
	last := uint64(0)
	product := uint64(1)
	primePower := uint64(1)
	primePowerSum := uint64(1)
	for _, p := range factors {
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
