package numbertheory

import (
	"math"
)

// Prime factorization
func PrimeFactorization(n uint64) []uint64 {
	factors := make([]uint64, 0)
	if n < 1 {
		return factors
	}

	remainder := n

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

// Euler Totient
// https://oeis.org/A000010
func Totient(n uint64) uint64 {
	return TotientK(n, 1)
}

// Jordan function
// k=2 - https://oeis.org/A007434
func TotientK(n uint64, k uint64) uint64 {
	J := math.Pow(float64(n), float64(k))
	primes := PrimeFactorization(n)

	last := uint64(0)
	for _, p := range primes {
		if p != last {
			J *= 1 - 1/math.Pow(float64(p), float64(k))
		}
		last = p
	}

	return uint64(J)
}

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
