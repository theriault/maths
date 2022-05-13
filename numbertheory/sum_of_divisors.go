package numbertheory

// Sum-of-divisors function sums the divisors of an integer.
//
// https://en.wikipedia.org/wiki/Divisor_function
// https://oeis.org/A000203
func SumOfDivisors(n uint64) uint64 {
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
