package numbertheory

import (
	"math"

	"github.com/theriault/maths"
)

var smallPrimes = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}
var step = []uint64{2, 4, 2, 4, 2, 4, 6}

func init() {
	m := 100
	s := make([]bool, m)
	p := 0
	for k, v := range s {
		if v {
			continue
		}
		p++
		l := k + 2
		for z := l*l - 2; z < m; z += l {
			s[z] = true
		}
	}
	smallPrimes = make([]uint64, p)
	p = 0
	for k, v := range s {
		if !v {
			smallPrimes[p] = uint64(k + 2)
			p++
		}
	}
}

// PrimeFactorization returns the prime factorization of n.
//
// https://en.wikipedia.org/wiki/Integer_factorization
func PrimeFactorization[A maths.Integer](n A) []uint64 {
	factors := make([]uint64, 0)
	if n <= 1 {
		return factors
	}

	remainder := uint64(n)

	for _, p := range smallPrimes {
		for remainder%p == 0 {
			factors = append(factors, p)
			remainder /= p
		}
	}
	divisor := uint64(41)
	f := 0
	l := 0
	sqrtn := uint64(math.Sqrt(float64(remainder)))

	for divisor <= sqrtn {
		for _, s := range step {
			for remainder%divisor == 0 {
				f++
				factors = append(factors, divisor)
				remainder /= divisor
			}
			divisor += s
		}
		if f != l {
			l, sqrtn = f, uint64(math.Sqrt(float64(remainder)))
		}
	}

	if remainder > 1 {
		factors = append(factors, remainder)
	}

	return factors
}
