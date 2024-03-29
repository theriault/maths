package numbertheory

import (
	"fmt"
	"math"

	"github.com/theriault/maths"
)

// PolygonalNumber returns the nth s-gon number. If s < 3, returns an error.
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Polygonal_number
func PolygonalNumber[A maths.Integer, B maths.Integer](s A, n B) (uint64, error) {
	if s < 3 {
		return 0, fmt.Errorf("s must be >= 3")
	}
	si := uint64(s)
	ni := uint64(n)
	return (si-2)*(ni*(ni-1))/2 + ni, nil
}

// PolygonalRoot returns the s-gonal root of x. If s < 3, returns an error.
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Polygonal_number
func PolygonalRoot[A maths.Integer, B maths.Integer](s A, x B) (float64, error) {
	if s < 3 {
		return 0, fmt.Errorf("s must be >= 3")
	}
	s2 := float64(s) - 2
	s4 := float64(s) - 4
	return (math.Sqrt(8*s2*float64(x)+math.Pow(s4, 2)) + s4) / (2 * s2), nil
}

// PolygonalSides returns the number of sides that the nth polygonal number with
// a value of x has.
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// https://en.wikipedia.org/wiki/Polygonal_number
func PolygonalSides[A maths.Integer, B maths.Integer](n A, x B) float64 {
	return 2 + (2/float64(n))*(float64(x)-float64(n))/float64(n-1)
}
