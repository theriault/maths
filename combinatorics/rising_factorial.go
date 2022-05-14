package combinatorics

import "fmt"

// RisingFactorial returns the number of ways of arranging n items in x ordered lists where each list may have 0..n
// items. It is possible that the result may overflow.
//
// https://en.wikipedia.org/wiki/Falling_and_rising_factorials
func RisingFactorial[
	A int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
	B int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64,
](x A, n B) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be >= 0")
	}
	if n == 0 {
		return 1, nil
	}
	product := uint64(x)
	for i, n := 1, int(n); i < n; i++ {
		product *= (uint64(x) + uint64(i))
	}
	return product, nil
}
