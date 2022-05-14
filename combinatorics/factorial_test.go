package combinatorics

import (
	"fmt"
	"testing"
)

func ExampleFactorial() {
	n, _ := Factorial(5)
	fmt.Printf("The factorial of 5 is %v.\n", n)

	// Output:
	// The factorial of 5 is 120.
}

func TestFactorial(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[int64]struct {
		expected uint64
		err      bool
	}{
		-1: {0, true},
		0:  {1, false},
		5:  {120, false},
		12: {479001600, false},
		13: {6227020800, false},
		50: {0, true},
	}
	for n, c := range cases {
		n, c := n, c
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			actual, err := Factorial(n)
			compare(t, c.expected, actual)
			if c.err != (err != nil) {
				t.Logf("expected error %v, got error %v", c.err, err != nil)
				t.FailNow()
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(i)
	}
}
