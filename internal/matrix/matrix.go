// matrix package contains the matrix type and related methods and transformations.
package matrix

import (
	"fmt"
	"strconv"

	"github.com/kaizencodes/glimpse/internal/utils"
)

type Matrix [][]float64

func New(n, m int) Matrix {
	mat := make(Matrix, n)
	for i := 0; i < int(n); i++ {
		mat[i] = make([]float64, m)
	}
	return mat
}

func NewIdentity(size int) Matrix {
	identity := New(size, size)
	for n := range identity {
		identity[n][n] = 1
	}
	return identity
}

func (m Matrix) String() string {
	var result string

	for _, row := range m {
		for _, val := range row {
			result += strconv.FormatFloat(val, 'f', -1, 64)
			result += ", "
		}
		result += string('\n')
	}
	return result
}

// Transposing the matrix swaps the rows and columns.
func (a Matrix) Transpose() Matrix {
	mat := New(len(a[0]), len(a))
	for n := 0; n < len(mat); n++ {
		for m := 0; m < len(mat[0]); m++ {
			mat[n][m] = a[m][n]
		}
	}
	return mat
}

// Inverse of a matrix is a matrix that when multiplied by the original matrix results in the identity matrix.
func (a Matrix) Inverse() Matrix {
	det := a.determinant()
	if det == 0 {
		panic(fmt.Errorf("non-invertible matrix, determinant is zero for \n%s", a.String()))
	}

	inverse := New(len(a), len(a[0]))
	for n, col := range inverse {
		for m := range col {
			// the col and row are swapped here. n m vs m n
			inverse[m][n] = a.Cofactor(n, m) / det
		}
	}
	return inverse
}

func (m Matrix) At(row, col int) float64 {
	return m[row][col]
}

// multiplications are used to perform transformations: scaling, rotating, translating.
func Multiply(a, b Matrix) Matrix {
	if len(a[0]) != len(b) {
		panic(fmt.Errorf("incompatible matrices: len: col a: %d, col: b  %d", len(a[0]), len(b)))
	}

	new_mat := New(len(a), len(b[0]))
	for n, row := range new_mat {
		for m := range row {
			new_mat[n][m] = dot(a, b, n, m)
		}
	}
	return new_mat
}

func (m Matrix) Equal(other Matrix) bool {
	if len(m) != len(other) || len(m[0]) != len(other[0]) {
		return false
	}

	for i, row := range m {
		for j, val := range row {
			if !utils.FloatEquals(val, other[i][j]) {
				return false
			}
		}
	}
	return true
}

// Calculates the determinant of a matrix.
func (a Matrix) determinant() float64 {
	if len(a) == 2 {
		return float64(a[0][0]*a[1][1] - a[0][1]*a[1][0])
	} else {
		var deter float64
		for n, elem := range a[0] {
			deter += elem * a.Cofactor(0, n)
		}
		return deter
	}
}

func (a Matrix) submatrix(col, row int) Matrix {
	new_mat := New(len(a), len(a[0]))
	for n, col := range a {
		copy(new_mat[n], col)
	}

	new_mat = append(new_mat[:col], new_mat[col+1:]...)
	for n := 0; n < len(new_mat); n++ {
		new_mat[n] = append(new_mat[n][:row], new_mat[n][row+1:]...)
	}

	return new_mat
}

func (a Matrix) minor(col, row int) float64 {
	return a.submatrix(col, row).determinant()
}

func (a Matrix) Cofactor(col, row int) float64 {
	deter := a.minor(col, row)
	if (col+row)%2 != 0 {
		deter *= -1
	}

	return deter
}

func dot(a, b Matrix, row, col int) float64 {
	var sum float64
	for i := range a[0] {
		sum += a[row][i] * b[i][col]
	}

	return sum
}
