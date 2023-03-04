package combinatorics

import (
	"fmt"

	"github.com/theriault/maths"
)

var factorialMemoization [21]uint64 = [21]uint64{
	1,                   // 0!
	1,                   // 1!
	2,                   // 2!
	6,                   // 3!
	24,                  // 4!
	120,                 // 5!
	720,                 // 6!
	5040,                // 7!
	40320,               // 8!
	362880,              // 9!
	3628800,             // 10!
	39916800,            // 11!
	479001600,           // 12!
	6227020800,          // 13!
	87178291200,         // 14!
	1307674368000,       // 15!
	20922789888000,      // 16!
	355687428096000,     // 17!
	6402373705728000,    // 18!
	121645100408832000,  // 19!
	2432902008176640000, // 20!
}

// Factorial returns the product of all positive integers less than or equal to n. n must be 0 <= n <= 20.
//
// Time complexity: O(1) - values are already precomputed for uint64
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Factorial
//
// https://oeis.org/A000142
func Factorial[A maths.Integer](n A) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be >= 0")
	}
	if n >= 21 {
		return 0, fmt.Errorf("n >= 21 will overflow uint64")
	}
	return factorialMemoization[int(n)], nil
}
