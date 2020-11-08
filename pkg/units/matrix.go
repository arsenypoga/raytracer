package units

import (
	"fmt"
	"math"
)

// Matrix is a matrix /shrug
type Matrix struct {
	Width  int
	Height int
	matrix [][]float64
}

// NewMatrix Creates new matrix from the matrix passed in
func NewMatrix(matrix [][]float64) *Matrix {
	returnMatrix := Matrix{
		Width:  len(matrix[0]),
		Height: len(matrix),
	}

	returnMatrix.matrix = make([][]float64, returnMatrix.Height)
	for i := range returnMatrix.matrix {
		returnMatrix.matrix[i] = make([]float64, returnMatrix.Width)
		copy(returnMatrix.matrix[i], matrix[i])
	}
	return &returnMatrix
}

// DefaultMatrix creates new empty matrix of given width and height
func DefaultMatrix(width, height int) *Matrix {
	matrix := make([][]float64, height)
	for i := range matrix {
		matrix[i] = make([]float64, width)
	}

	return &Matrix{
		Width:  width,
		Height: height,
		matrix: matrix,
	}
}

// Dot multiplies two matrieces returning a new matrix
func (m Matrix) Dot(m1 *Matrix) *Matrix {
	if m.Width != m1.Height {
		panic("Matrices need to be same size") // TODO: Provide accurate message, I am lazy lol
	}

	returnMatrix := DefaultMatrix(m1.Width, m.Height)
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m1.Width; j++ {
			for k := 0; k < m1.Height; k++ {
				returnMatrix.matrix[i][j] += m.matrix[i][k] * m1.matrix[k][j]
			}
		}
	}
	return returnMatrix
}

// TupleMultiply Multiply Matrix by a Tuple resulting in a new Tuple
func (m Matrix) TupleMultiply(t *Tuple) *Tuple {
	tupleMatrix := NewMatrix([][]float64{{t.X}, {t.Y}, {t.Z}, {t.W}})
	result := m.Dot(tupleMatrix)
	return &Tuple{result.matrix[0][0], result.matrix[1][0], result.matrix[2][0], result.matrix[3][0]}
}

func copyMatrix(m *Matrix) *Matrix {
	matrixCopy := DefaultMatrix(m.Width, m.Height)
	for i := 0; i < matrixCopy.Height; i++ {
		for j := 0; j < matrixCopy.Width; j++ {
			matrixCopy.matrix[i][j] = m.matrix[i][j]
		}
	}

	return matrixCopy
}

// Transpose transposes the matrix
func (m Matrix) Transpose() *Matrix {
	matrixCopy := DefaultMatrix(m.Height, m.Width)

	for i := range m.matrix {
		for j := range m.matrix[0] {
			matrixCopy.matrix[j][i] = m.matrix[i][j]
		}
	}
	return matrixCopy
}

// Submatrix returns a modified matrix without given row, col
func (m Matrix) Submatrix(i, j int) *Matrix {
	matrixCopy := copyMatrix(&m)
	matrixCopy.Width--
	matrixCopy.Height--

	for row := range matrixCopy.matrix {
		matrixCopy.matrix[row] = append(matrixCopy.matrix[row][:j], matrixCopy.matrix[row][j+1:]...)
	}
	matrixCopy.matrix = append(matrixCopy.matrix[:i], matrixCopy.matrix[i+1:]...)

	return matrixCopy
}

// Minor computes determinant of submatrix at i, j
func (m Matrix) Minor(i, j int) float64 {
	submatrix := m.Submatrix(i, j)
	return submatrix.Determinant()
}

// Cofactor computes cofactor of submatrix at i, j
func (m Matrix) Cofactor(i, j int) float64 {
	det := m.Minor(i, j)

	if (i+j)%2 != 0 {
		det = -det
	}
	return det
}

// Get reads the value at the given x and y
func (m Matrix) Get(i, j int) float64 {
	return m.matrix[i][j]
}

// Determinant computes determinant of a matrix
func (m Matrix) Determinant() float64 {
	det := 0.
	if m.Width == 2 {
		det = m.matrix[0][0]*m.matrix[1][1] - m.matrix[0][1]*m.matrix[1][0]
	} else {
		for row := 0; row < m.Width; row++ {
			det += m.matrix[0][row] * m.Cofactor(0, row)
		}
	}

	return det
}

// Invert inverts matrix if possible an fails if not
func (m Matrix) Invert() *Matrix {
	det := m.Determinant()

	if det == 0 {
		panic("Matrix is not invertible")
	}

	returnMatrix := DefaultMatrix(m.Width, m.Height)

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			c := m.Cofactor(i, j)
			returnMatrix.matrix[j][i] = c / det
		}
	}

	return returnMatrix

}

// IdentityMatrix returns identity matrix
func IdentityMatrix() *Matrix {
	return NewMatrix([][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})
}

// Translate translates the matrix with given coordinates
func (m Matrix) Translate(x, y, z float64) *Matrix {
	returnMatrix := IdentityMatrix()

	returnMatrix.matrix[0][3] = x
	returnMatrix.matrix[1][3] = y
	returnMatrix.matrix[2][3] = z

	return m.Dot(returnMatrix)
}

// Scale scales the matrix with the given coordinates
func (m Matrix) Scale(x, y, z float64) *Matrix {
	returnMatrix := IdentityMatrix()

	returnMatrix.matrix[0][0] = x
	returnMatrix.matrix[1][1] = y
	returnMatrix.matrix[2][2] = z

	return m.Dot(returnMatrix)
}

// RotateX rotates matrix on X axis by r radians
func (m Matrix) RotateX(r float64) *Matrix {
	returnMatrix := IdentityMatrix()

	returnMatrix.matrix[1][1] = math.Cos(r)
	returnMatrix.matrix[1][2] = -math.Sin(r)
	returnMatrix.matrix[2][1] = math.Sin(r)
	returnMatrix.matrix[2][2] = math.Cos(r)

	return m.Dot(returnMatrix)
}

// RotateY rotates matrix on Y axis by r radians
func (m Matrix) RotateY(r float64) *Matrix {
	returnMatrix := IdentityMatrix()

	returnMatrix.matrix[0][0] = math.Cos(r)
	returnMatrix.matrix[0][2] = math.Sin(r)
	returnMatrix.matrix[2][0] = -math.Sin(r)
	returnMatrix.matrix[2][2] = math.Cos(r)

	return m.Dot(returnMatrix)
}

// RotateZ rotates matrix on Z axis by r radians
func (m Matrix) RotateZ(r float64) *Matrix {
	returnMatrix := IdentityMatrix()

	returnMatrix.matrix[0][0] = math.Cos(r)
	returnMatrix.matrix[0][1] = -math.Sin(r)
	returnMatrix.matrix[1][0] = math.Sin(r)
	returnMatrix.matrix[1][1] = math.Cos(r)

	return m.Dot(returnMatrix)
}

// Shear shears the matrix on given coordinates
func (m Matrix) Shear(xy, xz, yx, yz, zx, zy float64) *Matrix {
	returnMatrix := IdentityMatrix()

	returnMatrix.matrix[0][1] = xy
	returnMatrix.matrix[0][2] = xz
	returnMatrix.matrix[1][0] = yx
	returnMatrix.matrix[1][2] = yz
	returnMatrix.matrix[2][0] = zx
	returnMatrix.matrix[2][1] = zy

	return m.Dot(returnMatrix)
}

// Equal compares two matrices
func (m Matrix) Equal(m1 *Matrix) bool {
	if m.Width != m1.Width || m.Height != m1.Height {
		return false
	}

	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			if !FloatEqual(m.matrix[i][j], m1.matrix[i][j]) {
				return false
			}
		}
	}
	return true
}

func (m Matrix) String() string {
	returnString := "Matrix(["
	for i := 0; i < m.Height; i++ {
		returnString += "\n    ["
		for j := 0; j < m.Width; j++ {
			returnString += fmt.Sprintf("%f", m.matrix[i][j])
			if j != m.Width-1 {
				returnString += ", "
			}
		}
		if i != m.Height-1 {
			returnString += "],"
		} else {
			returnString += "]"
		}
	}
	returnString += "\n])"
	return returnString
}
