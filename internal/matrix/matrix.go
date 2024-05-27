// matrix package contains the matrix type and related methods and transformations.
package matrix

import (
	"fmt"
	"strconv"

	"github.com/kaizencodes/glimpse/internal/utils"
)

type Matrix struct {
	data               [16]float64
	row_size, col_size int
}

func NewEmpty(n, m int) Matrix {
	return Matrix{
		[16]float64{},
		n, m,
	}
}

func New(n int, m int, data [16]float64) Matrix {
	return Matrix{
		data,
		n, m,
	}
}

func NewIdentity(size int) Matrix {
	identity := NewEmpty(size, size)
	for i := 0; i < size; i++ {
		identity.data[i*size+i] = 1
	}
	return identity
}

func (m Matrix) String() string {
	var result string

	elem_count := 0
	for i := range m.data {
		result += strconv.FormatFloat(m.data[i], 'f', -1, 64)
		result += ", "
		elem_count += 1
		if elem_count == m.col_size {
			elem_count = 0
			result += string('\n')
		}
	}
	return result
}

// Transposing the matrix swaps the rows and columns.
func (a Matrix) Transpose() Matrix {
	mat := NewEmpty(a.col_size, a.row_size)

	for i := 0; i < mat.row_size; i++ {
		for j := 0; j < mat.col_size; j++ {
			mat.data[i*mat.col_size+j] = a.At(j, i)
		}
	}
	return mat
}

func (m Matrix) At(row, col int) float64 {
	return m.data[row*m.col_size+col]
}

func (m Matrix) Equal(other Matrix) bool {
	if len(m.data) != len(other.data) || m.col_size != other.col_size || m.row_size != other.row_size {
		return false
	}

	for i, elem := range m.data {
		if !utils.FloatEquals(elem, other.data[i]) {
			return false
		}
	}
	return true
}

// multiplications are used to perform transformations: scaling, rotating, translating.
func Multiply(a, b Matrix) Matrix {
	if a.col_size != b.row_size {
		panic(fmt.Errorf("incompatible matrices: len: col a: %d, row: b  %d", a.col_size, b.row_size))
	}

	new_mat := NewEmpty(a.row_size, b.col_size)
	for i := 0; i < a.row_size; i++ {
		for j := 0; j < b.col_size; j++ {
			new_mat.data[i*new_mat.col_size+j] = dot(a, b, i, j)
		}
	}
	return new_mat
}

func dot(a, b Matrix, row, col int) float64 {
	var sum float64
	for i := 0; i < a.col_size; i++ {
		sum += a.At(row, i) * b.At(i, col)
	}

	return sum
}
