package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleNumberOfDivisors() {
	n := NumberOfDivisors(48)
	fmt.Printf("The number-of-divisors of 48 is %v.\n", n)

	// Output:
	// The number-of-divisors of 48 is 10.
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
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, NumberOfDivisors(n))
		})
	}
}

func BenchmarkNumberOfDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOfDivisors(i)
	}
}
