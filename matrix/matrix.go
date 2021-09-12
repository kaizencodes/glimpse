package matrix

import (
	"fmt"
	"strconv"
)

type Matrix [][]float64

func (a Matrix) Transpose() Matrix {
	mat := New(len(a[0]), len(a))
	for n := 0; n < len(mat); n++ {
		for m := 0; m < len(mat[0]); m++ {
			mat[n][m] = a[m][n]
		}
	}
	return mat
}

func (a Matrix) Determinant() float64 {
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

func (a Matrix) Submatrix(col, row int) Matrix {
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

func (a Matrix) Minor(col, row int) float64 {
	return a.Submatrix(col, row).Determinant()
}

func (a Matrix) Cofactor(col, row int) float64 {
	deter := a.Minor(col, row)
	if (col+row)%2 != 0 {
		deter *= -1
	}

	return deter
}

func (a Matrix) Inverse() (Matrix, error) {
	det := a.Determinant()
	if det == 0 {
		return nil, fmt.Errorf("noninvertible matrix, determinant is zero")
	}

	inverse := New(len(a), len(a[0]))
	for n, col := range inverse {
		for m, _ := range col {
			inverse[m][n] = a.Cofactor(n, m) / det
		}
	}
	return inverse, nil
}

func New(n, m int) Matrix {
	mat := make(Matrix, n)
	for i := 0; i < int(n); i++ {
		mat[i] = make([]float64, m)
	}
	return mat
}

func (m Matrix) String() string {
	var result string

	for _, row := range m {
		for _, val := range row {
			result += strconv.FormatFloat(val, 'f', -1, 64)
		}
		result += string('\n')
	}
	return result
}

func NewIdentity(size int) Matrix {
	identity := New(size, size)
	for n, _ := range identity {
		identity[n][n] = 1
	}
	return identity
}

func Multiply(a, b Matrix) (Matrix, error) {
	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("incompatible matrices: len: col a: %d, col: b  %d.", len(a[0]), len(b))
	}

	new_mat := New(len(a), len(b[0]))
	for n, row := range new_mat {
		for m, _ := range row {
			new_mat[n][m] = dot(a, b, n, m)
		}
	}
	return new_mat, nil
}

func dot(a, b Matrix, row, col int) float64 {
	var sum float64
	for i, _ := range a[0] {
		sum += a[row][i] * b[i][col]
	}

	return sum
}
