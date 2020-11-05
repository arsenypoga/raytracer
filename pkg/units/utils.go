package units

import (
	"math"
)

// EPSILON is a value used for comparisson of small values
const EPSILON = 0.00001

// Equal compares two small values
func Equal(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

// TupleEqual Compares two tuples.
func TupleEqual(a, b *Tuple) bool {
	return Equal(a.X, b.X) && Equal(a.Y, b.Y) && Equal(a.Z, b.Z) && Equal(a.W, b.W)
}
