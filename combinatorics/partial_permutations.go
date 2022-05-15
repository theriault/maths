package combinatorics

import (
	"fmt"

	"github.com/theriault/maths"
)

// PartialPermutations returns the number of ordered arrangements in which no element occurs more than once, but without
// the requirement of using all the elements from a given set (using only k elements)
//
//        n!
//    ----------
//     (n - k)!
//
// https://en.wikipedia.org/wiki/Permutation
func PartialPermutations[A maths.Integer, B maths.Integer](n A, k B) (uint64, error) {
	if k < 0 {
		return 0, fmt.Errorf("k must be >= 0")
	}
	return FallingFactorial(n, k)
}
