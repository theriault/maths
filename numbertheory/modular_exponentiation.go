package numbertheory

import "github.com/theriault/maths"

// ModExp is exponentiation performed over a modulus.
//
// Time complexity: O(log n) - where n is the exponent
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Modular_exponentiation
func ModExp[A maths.Integer](base, exponent, modulus A) int64 {
	result := int64(1)
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * int64(base)) % int64(modulus)
		}
		exponent = exponent >> 1
		base = (base * base) % modulus
	}
	return result
}
