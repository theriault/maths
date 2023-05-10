package numbertheory

import "github.com/theriault/maths"

// ExtendedGCD computes the GCD and the coefficients of Bézout's identity for integers a and b, which are the integers
// x and y such that ax + by = gcd(a,b).
//
// Time complexity:  O(log n)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm
func ExtendedGCD[Integer maths.Integer](a, b Integer) (Integer, Integer, Integer) {
	var quotient Integer
	x, y, u, v := Integer(0), Integer(1), Integer(1), Integer(0)
	for a != 0 {
		quotient, b, a = b/a, a, b%a
		x, y, u, v = u, v, x-u*quotient, y-v*quotient
	}
	return Integer(b), x, y
}
