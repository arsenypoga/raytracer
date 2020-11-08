package units

import (
	"math"
)

// EPSILON is a value used for comparisson of small values
const EPSILON = 0.00001

// FloatEqual compares two small values
func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

// TupleEqual Compares two tuples.
func TupleEqual(a, b *Tuple) bool {
	return FloatEqual(a.X, b.X) && FloatEqual(a.Y, b.Y) && FloatEqual(a.Z, b.Z) && FloatEqual(a.W, b.W)
}
