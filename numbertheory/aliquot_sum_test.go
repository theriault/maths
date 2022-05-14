package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleAliquotSum() {
	n := AliquotSum(60)
	fmt.Printf("The aliquot sum of 60 is %v.\n", n)

	// Output:
	// The aliquot sum of 60 is 108.
}

func TestAliquotSum(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		1:  1,
		11: 1,
		18: 21,
		42: 54,
		60: 108,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, AliquotSum(n))
		})
	}
}
