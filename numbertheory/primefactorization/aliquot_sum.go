package primefactorization

// AliquotSum returns the sum of all proper divisors of n (all divisors except n itself).
//
// Time complexity: O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Aliquot_sum
//
// https://oeis.org/A001065
func (pf *PrimeFactorization) AliquotSum() uint64 {
	return pf.SumOfDivisors() - pf.n
}
