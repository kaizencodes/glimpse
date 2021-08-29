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
