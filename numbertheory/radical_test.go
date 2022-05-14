package numbertheory

import (
	"fmt"
	"testing"
)

func ExampleRadical() {
	r := Radical(60)
	fmt.Printf("The radical of 60 is %v.\n", r)

	// Output:
	// The radical of 60 is 30.
}

func TestRadical(t *testing.T) {
	compare := func(t *testing.T, a uint64, b uint64) {
		if a != b {
			t.Logf("expected %v, got %v", a, b)
			t.FailNow()
		}
	}
	cases := map[uint64]uint64{
		1:  1,
		12: 6,
		53: 53,
		60: 30,
	}
	for n, expected := range cases {
		n, expected := n, expected
		t.Run(fmt.Sprintf("test %v", n), func(t *testing.T) {
			t.Parallel()
			compare(t, expected, Radical(n))
		})
	}
}

func BenchmarkRadical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Radical(i)
	}
}
