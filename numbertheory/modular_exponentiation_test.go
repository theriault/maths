package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleModExp() {
	n := ModExp(2, 256, 1_000_000_007)
	fmt.Printf("The modular exponentiation of 2^256 %% (10^6+7) is %v.\n", n)

	// Output:
	// The modular exponentiation of 2^256 % (10^6+7) is 792845266.
}

func TestModExp(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := []struct {
		Base     int
		Exponent int
		Modulus  int
		Expected int64
	}{
		{2, 256, 1_000_000_007, 792845266},
		{256, 20, 43, 21},
		{256, 256, 251, 63},
		{123456789, 123456789, 987654321, 598987215},
	}
	for n, c := range cases {
		n, c := n, c
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, uint64(c.Expected), uint64(ModExp(c.Base, c.Exponent, c.Modulus)))
		})
	}
}
