package matrix

import (
	"fmt"
	"strconv"
)

type Matrix [][]float64

func New(size int) Matrix {
	m := make(Matrix, size)
	for i := 0; i < int(size); i++ {
		m[i] = make([]float64, size)
	}
	return m
}

func (m Matrix) String() string {
	var result string

	for _, row := range m {
		for _, val := range row {
			result += fmt.Sprintf("%s, ", strconv.FormatFloat(val, 'f', -1, 64))
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

func dot(a, b Matrix, row, col int) float64 {
	var sum float64
	for i, _ := range a {
		sum += a[row][i] * b[i][col]
	}

	return sum
}

func Transpose(a Matrix) Matrix {
	new_mat := New(len(a))
	for n := 0; n < len(a)-1/2; n++ {
		for m := 0; m < len(a[n])-1/2; m++ {
			new_mat[n][m], new_mat[m][n] = a[m][n], a[n][m]
		}
	}
	return new_mat
}

func Determinant(a Matrix) float64 {
	if len(a) == 2 {
		return a[0][0]*a[1][1] - a[0][1]*a[1][0]
	} else {
		deter := 0.0
		for n, elem := range a[0] {
			deter += elem * Cofactor(a, 0, n)
		}
		return deter
	}
}

func Submatrix(a Matrix, col, row int) Matrix {
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

func Minor(a Matrix, col, row int) float64 {
	return Determinant(Submatrix(a, col, row))
}

func Cofactor(a Matrix, col, row int) float64 {
	deter := Minor(a, col, row)
	if (col+row)%2 != 0 {
		deter *= -1
	}

	return deter
}

func Inverse(a Matrix) (Matrix, error) {
	det := Determinant(a)
	if det == 0 {
		return nil, fmt.Errorf("noninvertible matrix, determinant is zero")
	}

	inverse := New(len(a))
	for n, col := range inverse {
		for m, _ := range col {
			inverse[m][n] = Cofactor(a, n, m) / det
		}
	}
	return inverse, nil
}
