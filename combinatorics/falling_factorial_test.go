package combinatorics

import (
	"fmt"
	"testing"
)

func ExampleFallingFactorial() {
	n, _ := FallingFactorial(8, 3)
	fmt.Printf("The number of different gold-silver-bronze podiums in an 8 person race is %v.\n", n)

	// Output:
	// The number of different gold-silver-bronze podiums in an 8 person race is 336.
}

func TestFallingFactorial(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := []struct {
		x        int
		n        int
		expected uint64
		err      bool
	}{
		{5, -1, 0, true},
		{5, 0, 1, false},
		{0, 1, 0, false},
		{5, 1, 5, false},
		{12, 2, 12 * 11, false},
		{13, 3, 13 * 12 * 11, false},
		{50, 5, 50 * 49 * 48 * 47 * 46, false},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test x=%v, n=%v", c.x, c.n), func(t *testing.T) {
			t.Parallel()
			actual, err := FallingFactorial(c.x, c.n)
			compare(t, c.expected, actual)
			if c.err != (err != nil) {
				t.Logf("expected error %v, got error %v", c.err, err != nil)
				t.FailNow()
			}
		})
	}
}

func BenchmarkFallingFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FallingFactorial(i, 5)
	}
}
