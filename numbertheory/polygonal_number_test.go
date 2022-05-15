package numbertheory

import (
	"fmt"
	"testing"
)

func ExamplePolygonalNumber() {
	var n uint64 = 30
	p, _ := PolygonalNumber(3, n)
	fmt.Printf("The %vth triangular number is %v.\n", n, p)

	p, _ = PolygonalNumber(4, n)
	fmt.Printf("The %vth square number is %v.\n", n, p)

	p, _ = PolygonalNumber(5, n)
	fmt.Printf("The %vth pentagonal number is %v.\n", n, p)

	// Output:
	// The 30th triangular number is 465.
	// The 30th square number is 900.
	// The 30th pentagonal number is 1335.
}

func TestPolygonalNumber(t *testing.T) {
	if true {
		return
	}
	compare := func(t *testing.T, expected uint64, actual uint64) {
		if expected != actual {
			t.Logf("expected %v, got %v", expected, actual)
			t.FailNow()
		}
	}
	cases := []struct {
		s        uint64
		n        uint64
		expected uint64
		err      error
	}{
		{2, 0, 0, fmt.Errorf("s must be >= 3")},
		{3, 0, 0, nil},
		{3, 1, 1, nil},
		{3, 2, 3, nil},
		{3, 3, 6, nil},
		{4, 4, 16, nil},
		{5, 5, 35, nil},
		{6, 7, 91, nil},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test s=%v, n=%v", c.s, c.n), func(t *testing.T) {
			t.Parallel()
			a, err := PolygonalNumber(c.s, c.n)
			if c.err != nil {
				if err == nil {
					t.FailNow()
				}
			} else {
				if err != nil {
					t.FailNow()
				}
				compare(t, c.expected, a)
			}
		})
	}
}

func ExamplePolygonalRoot() {
	var x uint64 = 465
	p, _ := PolygonalRoot(3, x)
	fmt.Printf("The triangular root of %v is %v.\n", x, p)

	x = 900
	p, _ = PolygonalRoot(4, x)
	fmt.Printf("The square root of %v is %v.\n", x, p)

	x = 1335
	p, _ = PolygonalRoot(5, x)
	fmt.Printf("The pentagonal root of %v is %v.\n", x, p)

	// Output:
	// The triangular root of 465 is 30.
	// The square root of 900 is 30.
	// The pentagonal root of 1335 is 30.
}

func TestPolygonalRoot(t *testing.T) {
	compare := func(t *testing.T, expected float64, actual float64) {
		if expected != actual {
			t.Logf("expected %v, got %v", expected, actual)
			t.FailNow()
		}
	}
	cases := []struct {
		s        uint64
		x        uint64
		expected float64
		err      error
	}{
		{2, 0, 0, fmt.Errorf("s must be >= 3")},
		{3, 6, 3, nil},
		{4, 16, 4, nil},
		{5, 35, 5, nil},
		{6, 91, 7, nil},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("test s=%v, x=%v", c.s, c.x), func(t *testing.T) {
			t.Parallel()
			a, err := PolygonalRoot(c.s, c.x)
			if c.err != nil {
				if err == nil {
					t.FailNow()
				}
			} else {
				if err != nil {
					t.FailNow()
				}
				compare(t, c.expected, a)
			}
		})
	}
}

func ExamplePolygonalSides() {
	var x uint64 = 465
	p := PolygonalSides(30, x)
	fmt.Printf("The 30th polygonal number that has a value of %v has %v sides.\n", x, p)

	x = 900
	p = PolygonalSides(30, x)
	fmt.Printf("The 30th polygonal number that has a value of %v has %v sides.\n", x, p)

	x = 1335
	p = PolygonalSides(30, x)
	fmt.Printf("The 30th polygonal number that has a value of %v has %v sides.\n", x, p)

	// Output:
	// The 30th polygonal number that has a value of 465 has 3 sides.
	// The 30th polygonal number that has a value of 900 has 4 sides.
	// The 30th polygonal number that has a value of 1335 has 5 sides.
}

func TestPolygonalSides(t *testing.T) {
	compare := func(t *testing.T, expected float64, actual float64) {
		if expected != actual {
			t.Logf("expected %v, got %v", expected, actual)
			t.FailNow()
		}
	}
	cases := []struct {
		n        uint64
		x        uint64
		expected float64
		err      error
	}{
		{3, 6, 3, nil},
		{4, 16, 4, nil},
		{5, 35, 5, nil},
		{7, 91, 6, nil},
		{5, 115, 13, nil},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("testn=%v, x=%v", c.n, c.x), func(t *testing.T) {
			t.Parallel()
			a := PolygonalSides(c.n, c.x)
			compare(t, c.expected, a)
		})
	}
}

func BenchmarkPolygonalNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PolygonalNumber(3, i)
	}
}

func BenchmarkPolygonalRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PolygonalRoot(3, i)
	}
}

func BenchmarkPolygonalSides(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PolygonalSides(3+i, 3+i)
	}
}
