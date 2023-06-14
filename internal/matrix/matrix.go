// matrix package contains the matrix type and related methods and transformations.
package matrix

import (
	"log"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

// Hiding the implementation details of the matrix type with the inner attribute.
// Previously it was an internal solution, so kept the interface and changed the implementation
// to use gonum/mat package since it's much more efficient resulting in a 10x speedup.
type Matrix struct {
	inner mat.Matrix
}

func New(rows, cols int, data []float64) Matrix {
	return Matrix{inner: mat.NewDense(rows, cols, data)}
}

func (m Matrix) String() string {
	var result string

	rows, cols := m.inner.Dims()

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result += strconv.FormatFloat(m.inner.At(i, j), 'f', -1, 64)
			result += ", "
		}
		result += string('\n')
	}
	return result
}

func (m Matrix) Inverse() Matrix {
	var inv mat.Dense
	if err := inv.Inverse(m.inner); err != nil {
		log.Fatal(err)
	}
	return Matrix{inner: &inv}
}

func (m Matrix) Transpose() Matrix {
	return Matrix{inner: m.inner.T()}
}

func (m Matrix) At(row, col int) float64 {
	return m.inner.At(row, col)
}

// multiplications are used to perform transformations: scaling, rotating, translating.
func Multiply(a, b Matrix) Matrix {
	var m mat.Dense
	m.Mul(a.inner, b.inner)

	return Matrix{inner: &m}
}

func Equal(a, b Matrix) bool {
	return mat.Equal(a.inner, b.inner)
}
