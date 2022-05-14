package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleTotient() {
	r := Totient(60)
	fmt.Printf("Euler's totient of 60 is %v.\n", r)

	// Output:
	// Euler's totient of 60 is 16.
}

func TestTotient(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		12: 4,
		13: 12,
		38: 18,
		41: 40,
		60: 16,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, Totient(n))
		})
	}
}

func ExampleTotientK() {
	r := TotientK(60, 2)
	fmt.Printf("The J₂ of 60 is %v.", r)

	// Output:
	// The J₂ of 60 is 2304.
}

func TestTotientK2(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		17: 288,
		38: 1080,
		60: 2304,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, TotientK(n, 2))
		})
	}
}
