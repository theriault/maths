package combinatorics

import "fmt"

// PartialPermutations returns the number of ordered arrangements in which no element occurs more than once, but without
// the requirement of using all the elements from a given set (using only k elements)
//
//        n!
//    ----------
//     (n - k)!
//
// https://en.wikipedia.org/wiki/Permutation
func PartialPermutations[
	A int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
	B int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
](n A, k B) (uint64, error) {
	if k < 0 {
		return 0, fmt.Errorf("k must be >= 0")
	}
	return FallingFactorial(n, k)
}
