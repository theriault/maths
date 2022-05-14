package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleMobius() {
	n := Mobius(70)
	fmt.Printf("The möbius of 70 is %v.\n", n)

	// Output:
	// The möbius of 70 is -1.
}

func TestMobius(t *testing.T) {
	compare := func(t *testing.T, a int8, b int8) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]int8{
		13: -1,
		14: 1,
		60: 0,
		70: -1,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, Mobius(n))
		})
	}
}
