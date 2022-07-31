package numbertheory

import (
	"math"

	"github.com/theriault/maths"
)

// DigitSum returns the sum of the individual digits of n within the given radix/base.
//
// If b=1, returns n (unary numeral system)
//
// If b<1 or n<=0, returns 0. Caller can verify it was an error by checking if n>0. A future version of this
// function may handle negative n.
//
// https://en.wikipedia.org/wiki/Digit_sum
func DigitSum[A maths.Integer](n, b A) A {
	// any b<1 will return 0. any n<1 will return 0. caller can verify it was an error by checking if n>0
	if b < 1 || n < 1 {
		return 0
	}
	// base 1 will be treated as the unary numeral system
	if b == 1 {
		return n
	}
	var bi int = 1
	var bip int = 1
	var nf int = int(n)
	var bf int = int(b)
	var s int
	var k int = int(math.Floor(math.Log2(float64(n)) / math.Log2(float64(b))))
	for i := 0; i <= k; i++ {
		bip *= bf
		s += (nf%bip - nf%bi) / bi
		bi = bip
	}
	return A(s)
}
