package units

import "math"

// Tuple is an interface describing behavior of tuple-like structs
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewPoint creates new point struct
func NewPoint(X, Y, Z float64) *Tuple {
	return &Tuple{X, Y, Z, 1.}
}

// NewVector creates new vector struct
func NewVector(X, Y, Z float64) *Tuple {
	return &Tuple{X, Y, Z, 0.}
}

// IsPoint returns if Tuple is point
func (t Tuple) IsPoint() bool {
	return t.W == 1.
}

// IsVector returns if Tuple is vector
func (t Tuple) IsVector() bool {
	return t.W == 0.
}

// Add adds two tuples
func (t Tuple) Add(t1 *Tuple) *Tuple {
	return &Tuple{
		X: t.X + t1.X,
		Y: t.Y + t1.Y,
		Z: t.Z + t1.Z,
		W: t.W + t1.W,
	}
}

// Substract substracts two tuples
func (t Tuple) Substract(t1 *Tuple) *Tuple {
	return &Tuple{
		X: t.X - t1.X,
		Y: t.Y - t1.Y,
		Z: t.Z - t1.Z,
		W: t.W - t1.W,
	}
}

// Negate inverts the Tuple
func (t Tuple) Negate() *Tuple {
	return &Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

// Multiply multiplies Tuple by a scalar
func (t Tuple) Multiply(s float64) *Tuple {
	return &Tuple{
		X: t.X * s,
		Y: t.Y * s,
		Z: t.Z * s,
		W: t.W * s,
	}
}

// Divide multiplies Tuple by a scalar
func (t Tuple) Divide(s float64) *Tuple {
	return &Tuple{
		X: t.X / s,
		Y: t.Y / s,
		Z: t.Z / s,
		W: t.W / s,
	}
}

// Magnitude computes magnitude of a Tuple
func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2))
}

// Normalize normalizes a vector
func (t Tuple) Normalize() *Tuple {
	magn := t.Magnitude()
	return &Tuple{
		X: t.X / magn,
		Y: t.Y / magn,
		Z: t.Z / magn,
		W: t.W / magn,
	}
}

// Dot computes dot product
func (t Tuple) Dot(t1 *Tuple) float64 {
	return t.X*t1.X + t.Y*t1.Y + t.Z*t1.Z + t.W*t1.W
}

// Cross computes cross produt
func (t Tuple) Cross(t1 *Tuple) *Tuple {
	return NewVector(
		t.Y*t1.Z-t.Z*t1.Y,
		t.Z*t1.X-t.X*t1.Z,
		t.X*t1.Y-t.Y*t1.X,
	)
}

// MatrixMultiply Multiplies Tuple by matrix
func (t Tuple) MatrixMultiply(m *Matrix) *Tuple {
	return m.TupleMultiply(&t)
}

// Translate translates tuple by x, y, z
func (t Tuple) Translate(x, y, z float64) *Tuple {
	return t.MatrixMultiply(IdentityMatrix().Translate(x, y, z))
}

// Scale scales the tuple by x, y, z
func (t Tuple) Scale(x, y, z float64) *Tuple {
	return t.MatrixMultiply(IdentityMatrix().Scale(x, y, z))
}
