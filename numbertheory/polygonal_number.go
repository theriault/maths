package numbertheory

import (
	"fmt"
	"math"
)

// Polygonal number
//
// https://en.wikipedia.org/wiki/Polygonal_number
func PolygonalNumber(s uint64, n uint64) (uint64, error) {
	if s < 3 {
		return 0, fmt.Errorf("s must be >= 3")
	}
	return (s-2)*(n*(n-1))/2 + n, nil
}

// Polygonal root
//
// https://en.wikipedia.org/wiki/Polygonal_number
func PolygonalRoot(s uint64, x uint64) (float64, error) {
	if s < 3 {
		return 0, fmt.Errorf("s must be >= 3")
	}
	s2 := float64(s) - 2
	s4 := float64(s) - 4
	return (math.Sqrt(8*s2*float64(x)+math.Pow(s4, 2)) + s4) / (2 * s2), nil
}

// Polygonal sides
//
// https://en.wikipedia.org/wiki/Polygonal_number
func PolygonalSides(n uint64, x uint64) float64 {
	return 2 + (2/float64(n))*float64(x-n)/float64(n-1)
}
