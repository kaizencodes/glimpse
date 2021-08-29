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
