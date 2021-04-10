// Package triangle provides utilities to deal with triangles
package triangle

import "math"

type Kind string

const (
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// KindFromSides calculates the kind of a triangle given its three sides
// it will also tell you if the three sides are not a valid triangle
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}

// isTriangle checks if the three sides form a valid triangle
func isTriangle(a, b, c float64) bool {
	if !isValidSide(a) || !isValidSide(b) || !isValidSide(c) {
		return false
	}

	if a+b < c || a+c < b || b+c < a {
		return false
	}

	return true
}

// isValidSide checks that a side is a positive number smaller than
// infinity
func isValidSide(side float64) bool {
	return !math.IsNaN(side) && !math.IsInf(side, 1) && side > 0
}
