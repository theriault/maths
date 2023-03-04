package numbertheory

import (
	"github.com/theriault/maths"
)

// Coprime returns whether two integers share no common divisors other than 1.
//
// Time complexity: O(log n) - where n is whichever number is smaller
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Coprime_integers
func Coprime[A maths.Integer](a, b A) bool {
	return GCD(a, b) == 1
}
