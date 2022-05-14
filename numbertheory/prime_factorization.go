package numbertheory

import (
	"math"
)

// PrimeFactorization returns the prime factorization of n.
//
// https://en.wikipedia.org/wiki/Integer_factorization
func PrimeFactorization[A Integer](n A) []uint64 {
	factors := make([]uint64, 0)
	if n <= 1 {
		return factors
	}

	remainder := uint64(n)

	for _, divisor := range []uint64{2, 3} {
		for remainder%divisor == 0 {
			factors = append(factors, divisor)
			remainder /= divisor
		}
	}

	divisor := uint64(5)
	sqrtn := math.Sqrt(float64(remainder))
	for float64(divisor) <= sqrtn {
		for remainder%divisor == 0 {
			factors = append(factors, divisor)
			remainder /= divisor
			sqrtn = math.Sqrt(float64(remainder))
		}
		divisor += 2
		for remainder%divisor == 0 {
			factors = append(factors, divisor)
			remainder /= divisor
			sqrtn = math.Sqrt(float64(remainder))
		}
		divisor += 4
	}

	if remainder > 1 {
		factors = append(factors, remainder)
	}

	return factors
}
