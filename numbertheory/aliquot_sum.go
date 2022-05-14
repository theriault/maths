package numbertheory

// AliquotSum returns the sum of all proper divisors of n (all divisors except n itself).
//
// https://en.wikipedia.org/wiki/Aliquot_sum
//
// https://oeis.org/A001065
func AliquotSum[A Integer](n A) uint64 {
	if n == 1 {
		return 1
	}
	return SumOfDivisors(n) - uint64(n)
}
