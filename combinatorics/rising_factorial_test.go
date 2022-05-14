package combinatorics

import (
	"fmt"
	"testing"
)

func ExampleRisingFactorial() {
	n, _ := RisingFactorial(2, 3)
	fmt.Printf("3 books on 2 different shelves can be arranged in %v ways.\n", n)

	// Output:
	// 3 books on 2 different shelves can be arranged in 24 ways.
}

func TestRisingFactorial(t *testing.T) {
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
		{12, 2, 12 * 13, false},
		{13, 3, 13 * 14 * 15, false},
		{50, 5, 50 * 51 * 52 * 53 * 54, false},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test x=%v, n=%v", c.x, c.n), func(t *testing.T) {
			t.Parallel()
			actual, err := RisingFactorial(c.x, c.n)
			compare(t, c.expected, actual)
			if c.err != (err != nil) {
				t.Logf("expected error %v, got error %v", c.err, err != nil)
				t.FailNow()
			}
		})
	}
}

func BenchmarkRisingFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RisingFactorial(i, 5)
	}
}
