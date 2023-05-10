package statistics

import (
	"fmt"
	"testing"
)

func ExampleDotProduct() {
	X1 := []uint8{1, 2, 3}
	Y1 := []uint8{3, 4, 5}
	product1 := DotProduct(X1, Y1)
	fmt.Printf("The dot product of %v and %v is %v.\n", X1, Y1, product1)

	X2 := []float64{1.1, 2.2, 3.3}
	Y2 := []float64{3.3, 4.4, 5.5}
	product2 := DotProduct(X2, Y2)
	fmt.Printf("The dot product of %v and %v is %v.\n", X2, Y2, product2)

	// Output:
	// The dot product of [1 2 3] and [3 4 5] is 26.
	// The dot product of [1.1 2.2 3.3] and [3.3 4.4 5.5] is 31.46.
}

func TestDotProduct(t *testing.T) {
	cases := []struct {
		name     string
		X        []int
		Y        []int
		expected string
	}{
		{
			name:     "test 1 - positives",
			X:        []int{1, 2, 3},
			Y:        []int{3, 4, 5},
			expected: "int 26",
		},
		{
			name:     "test 2 - zero",
			X:        []int{1, 2, 3},
			Y:        []int{0, 0, 0},
			expected: "int 0",
		},
		{
			name:     "test 3 - unequal lengths",
			X:        []int{1, 2, 3},
			Y:        []int{4, 5},
			expected: "int 14",
		},
		{
			name:     "test 4 - unequal lengths",
			X:        []int{1, 2, 3},
			Y:        []int{3, 4, 5, 6},
			expected: "int 26",
		},
		{
			name:     "test 5 - negatives",
			X:        []int{1, -2, 3},
			Y:        []int{3, 4, -5},
			expected: "int -20",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := DotProduct(c.X, c.Y)
			r := fmt.Sprintf("%T %v", actual, actual)
			if r != c.expected {
				t.Logf("expected %v, got %v", c.expected, r)
				t.FailNow()
			}
		})
	}
}
