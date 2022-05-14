package numbertheory

import (
	"fmt"
	"testing"
)

func ExamplePoliteness() {
	n := Politeness(42)
	fmt.Printf("The politeness of 42 is %v.\n", n)

	// Output:
	// The politeness of 42 is 3.
}

func TestPoliteness(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		1:  0,
		2:  0,
		4:  0,
		5:  1,
		9:  2,
		21: 3,
		27: 3,
		99: 5,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, Politeness(n))
		})
	}
}
