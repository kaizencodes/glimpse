package matrix

import (
	"fmt"
	"strconv"
)

type Element float64
type Matrix [][]Element

func New(size int) Matrix {
	m := make(Matrix, size)
	for i := 0; i < int(size); i++ {
		m[i] = make([]Element, size)
	}
	return m
}

func (e Element) String() string {
	return strconv.FormatFloat(float64(e), 'f', -1, 64)
}

func (m Matrix) String() string {
	var result string

	for _, row := range m {
		for _, val := range row {
			result += fmt.Sprintf("%s, ", val)
		}
		result += string('\n')
	}
	return result
}

func NewIdentity(size int) Matrix {
	identity := New(size)
	for n, _ := range identity {
		identity[n][n] = 1
	}
	return identity
}

func Multiply(a, b Matrix) (Matrix, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("incompatible matrices: len: %d, %d.", len(a), len(b))
	}

	new_mat := New(len(a))
	for n, row := range new_mat {
		for m, _ := range row {
			new_mat[n][m] = dot(a, b, n, m)
		}
	}
	return new_mat, nil
}

func dot(a, b Matrix, row, col int) Element {
	var sum Element
	for i, _ := range a {
		sum += a[row][i] * b[i][col]
	}

	return sum
}

func (a Matrix) Transpose() Matrix {
	new_mat := New(len(a))
	for n := 0; n < len(a)-1/2; n++ {
		for m := 0; m < len(a[n])-1/2; m++ {
			new_mat[n][m], new_mat[m][n] = a[m][n], a[n][m]
		}
	}
	return new_mat
}

func (a Matrix) Determinant() Element {
	if len(a) == 2 {
		return Element(a[0][0]*a[1][1] - a[0][1]*a[1][0])
	} else {
		var deter Element
		for n, elem := range a[0] {
			deter += elem * a.Cofactor(0, n)
		}
		return deter
	}
}

func (a Matrix) Submatrix(col, row int) Matrix {
	new_mat := New(len(a))
	for n, col := range a {
		copy(new_mat[n], col)
	}

	new_mat = append(new_mat[:col], new_mat[col+1:]...)
	for n := 0; n < len(new_mat); n++ {
		new_mat[n] = append(new_mat[n][:row], new_mat[n][row+1:]...)
	}

	return new_mat
}

func (a Matrix) Minor(col, row int) Element {
	return a.Submatrix(col, row).Determinant()
}

func (a Matrix) Cofactor(col, row int) Element {
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

	inverse := New(len(a))
	for n, col := range inverse {
		for m, _ := range col {
			inverse[m][n] = a.Cofactor(n, m) / det
		}
	}
	return inverse, nil
}
