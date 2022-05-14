package combinatorics

import (
	"fmt"
	"testing"
)

func ExamplePartialPermutations() {
	n, _ := PartialPermutations(8, 4)
	fmt.Printf("The partial permutations of 8 elements when given 4 at a time is %v.\n", n)

	// Output:
	// The partial permutations of 8 elements when given 4 at a time is 1680.
}

func TestPartialPermutations(t *testing.T) {
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
		{5, 6, 0, false},
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
			actual, err := PartialPermutations(c.x, c.n)
			compare(t, c.expected, actual)
			if c.err != (err != nil) {
				t.Logf("expected error %v, got error %v", c.err, err != nil)
				t.FailNow()
			}
		})
	}
}

func BenchmarkPartialPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartialPermutations(i, 5)
	}
}
