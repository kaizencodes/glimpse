package matrix

import (
	"fmt"
	"glimpse/tuple"
)

func Translate(t Transformable, x, y, z float64) (Transformable, error) {
	switch t := t.(type) {
	case tuple.Tuple:
		mat := New(4, 1)
		for i, v := range t.ToSlice() {
			mat[i][0] = float64(v)
		}
		mat, err := multiply_matrices(translation_matrix(x, y, z), mat)
		if err != nil {
			return nil, err
		}
		mat = mat.Transpose()
		return tuple.Tuple{mat[0][0], mat[0][1], mat[0][2], mat[0][3]}, nil

	default:
		return nil, fmt.Errorf("incompatible type for matrix multiplication: %T", t)
	}
}

func Scale(t Transformable, x, y, z float64) (Transformable, error) {
	switch t := t.(type) {
	case tuple.Tuple:
		mat := New(4, 1)
		for i, v := range t.ToSlice() {
			mat[i][0] = float64(v)
		}
		mat, err := multiply_matrices(scaling_matrix(x, y, z), mat)
		if err != nil {
			return nil, err
		}
		mat = mat.Transpose()
		return tuple.Tuple{mat[0][0], mat[0][1], mat[0][2], mat[0][3]}, nil

	default:
		return nil, fmt.Errorf("incompatible type for matrix multiplication: %T", t)
	}
}

func translation_matrix(x, y, z float64) Matrix {
	return Matrix{
		[]float64{1, 0, 0, x},
		[]float64{0, 1, 0, y},
		[]float64{0, 0, 1, z},
		[]float64{0, 0, 0, 1},
	}
}

func scaling_matrix(x, y, z float64) Matrix {
	return Matrix{
		[]float64{x, 0, 0, 0},
		[]float64{0, y, 0, 0},
		[]float64{0, 0, z, 0},
		[]float64{0, 0, 0, 1},
	}
}
