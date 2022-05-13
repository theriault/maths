package numbertheory

import (
	"fmt"
	"testing"
)

func TestSumOfDivisors(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		1:  1,
		11: 12,
		18: 39,
		24: 60,
		60: 168,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, SumOfDivisors(n))
		})
	}
}
