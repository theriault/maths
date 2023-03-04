package numbertheory

import (
	"fmt"

	"github.com/theriault/maths"
)

var primorialMemoization [16][2]uint64 = [16][2]uint64{
	{47, 614889782588491410}, // 47#
	{43, 13082761331670030},  // 43#
	{41, 304250263527210},    // 41#
	{37, 7420738134810},      // 37#
	{31, 200560490130},       // 31#
	{29, 6469693230},         // 29#
	{23, 223092870},          // 23#
	{19, 9699690},            // 19#
	{17, 510510},             // 17#
	{13, 30030},              // 13#
	{11, 2310},               // 11#
	{7, 210},                 // 7#
	{5, 30},                  // 5#
	{3, 6},                   // 3#
	{2, 2},                   // 2#
	{1, 1},                   // 1#
}

// Primorial returns the product of all primes up to n. n must be 0 <= n <= 52.
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Primorial
//
// https://oeis.org/A002110
func Primorial[A maths.Integer](n A) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be >= 0")
	}
	if n >= 53 {
		return 0, fmt.Errorf("n >= 53 will overflow uint64")
	}
	r := uint64(1)
	for _, p := range primorialMemoization {
		if uint64(n) >= p[0] {
			r = p[1]
			break
		}
	}
	return r, nil
}
