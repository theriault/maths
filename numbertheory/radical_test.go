package numbertheory

import (
	"fmt"
	"testing"
)

func TestRadical(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		1:  1,
		12: 6,
		53: 53,
		60: 30,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, Radical(n))
		})
	}
}
