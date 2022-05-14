package numbertheory

import (
	"fmt"
	"testing"
)

func ExamplePrimeFactorization() {
	var n uint64 = 51561510
	r := PrimeFactorization(n)
	fmt.Printf("The prime factorization of %v is %v.\n", n, r)

	// Output:
	// The prime factorization of 51561510 is [2 3 5 7 11 13 17 101].
}

func TestPrimeFactorization(t *testing.T) {
	compare := func(t *testing.T, a []uint64, b []uint64) {
		if len(a) != len(b) {
			t.FailNow()
		}
		for k, n := range a {
			if b[k] != n {
				t.FailNow()
			}
		}
	}
	cases := map[uint64][]uint64{
		0:                                  {},
		1:                                  {},
		2:                                  {2},
		2 * 2:                              {2, 2},
		2 * 3 * 5:                          {2, 3, 5},
		2 * 3 * 5 * 7 * 11 * 13 * 17:       {2, 3, 5, 7, 11, 13, 17},
		19 * 31 * 41 * 43:                  {19, 31, 41, 43},
		2 * 3 * 5 * 7 * 11 * 13 * 17 * 101: {2, 3, 5, 7, 11, 13, 17, 101},
	}
	for n, c := range cases {
		n, c := n, c
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, c, PrimeFactorization(n))
		})
	}
}

func BenchmarkPrimeFactorization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactorization(i)
	}
}
