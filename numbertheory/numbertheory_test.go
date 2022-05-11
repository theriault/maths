package numbertheory

import (
	"fmt"
	"testing"
)

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

	t.Run("test 0", func(t *testing.T) {
		t.Parallel()
		compare(t, []uint64{}, PrimeFactorization(0))
	})

	t.Run("test 1", func(t *testing.T) {
		t.Parallel()
		compare(t, []uint64{}, PrimeFactorization(1))
	})

	t.Run("test 2", func(t *testing.T) {
		t.Parallel()
		compare(t, []uint64{2}, PrimeFactorization(2))
	})

	t.Run("test 2*2", func(t *testing.T) {
		t.Parallel()
		compare(t, []uint64{2, 2}, PrimeFactorization(2*2))
	})

	t.Run("test 2*3*5", func(t *testing.T) {
		t.Parallel()
		compare(t, []uint64{2, 3, 5}, PrimeFactorization(2*3*5))
	})

	t.Run("test 2*3*5*7*11*13*17", func(t *testing.T) {
		t.Parallel()
		compare(t, []uint64{2, 3, 5, 7, 11, 13, 17}, PrimeFactorization(2*3*5*7*11*13*17))
	})

	t.Run("test 2*3*5*7*11*13*17*101", func(t *testing.T) {
		t.Parallel()
		compare(t, []uint64{2, 3, 5, 7, 11, 13, 17, 101}, PrimeFactorization(2*3*5*7*11*13*17*101))
	})
}

func TestNumberOfDivisors(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		11: 2,
		36: 9,
		48: 10,
	}
	for n, expected := range cases {
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, NumberOfDivisors(n))
		})
	}
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
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, Totient(n))
		})
	}
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
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, TotientK(n, 2))
		})
	}
}
