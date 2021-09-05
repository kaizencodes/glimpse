package matrix

import (
	"fmt"
	"glimpse/tuple"
	"math"
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

func RotateX(t Transformable, r float64) (Transformable, error) {
	switch t := t.(type) {
	case tuple.Tuple:
		mat := New(4, 1)
		for i, v := range t.ToSlice() {
			mat[i][0] = float64(v)
		}
		mat, err := multiply_matrices(rotation_x_matrix(r), mat)
		if err != nil {
			return nil, err
		}
		mat = mat.Transpose()
		return tuple.Tuple{mat[0][0], mat[0][1], mat[0][2], mat[0][3]}, nil

	default:
		return nil, fmt.Errorf("incompatible type for matrix multiplication: %T", t)
	}
}

func RotateY(t Transformable, r float64) (Transformable, error) {
	switch t := t.(type) {
	case tuple.Tuple:
		mat := New(4, 1)
		for i, v := range t.ToSlice() {
			mat[i][0] = float64(v)
		}
		mat, err := multiply_matrices(rotation_y_matrix(r), mat)
		if err != nil {
			return nil, err
		}
		mat = mat.Transpose()
		return tuple.Tuple{mat[0][0], mat[0][1], mat[0][2], mat[0][3]}, nil

	default:
		return nil, fmt.Errorf("incompatible type for matrix multiplication: %T", t)
	}
}

func RotateZ(t Transformable, r float64) (Transformable, error) {
	switch t := t.(type) {
	case tuple.Tuple:
		mat := New(4, 1)
		for i, v := range t.ToSlice() {
			mat[i][0] = float64(v)
		}
		mat, err := multiply_matrices(rotation_z_matrix(r), mat)
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

func rotation_x_matrix(r float64) Matrix {
	return Matrix{
		[]float64{1, 0, 0, 0},
		[]float64{0, math.Cos(r), -math.Sin(r), 0},
		[]float64{0, math.Sin(r), math.Cos(r), 0},
		[]float64{0, 0, 0, 1},
	}
}

func rotation_y_matrix(r float64) Matrix {
	return Matrix{
		[]float64{math.Cos(r), 0, math.Sin(r), 0},
		[]float64{0, 1, 0, 0},
		[]float64{-math.Sin(r), 0, math.Cos(r), 0},
		[]float64{0, 0, 0, 1},
	}
}

func rotation_z_matrix(r float64) Matrix {
	return Matrix{
		[]float64{math.Cos(r), -math.Sin(r), 0, 0},
		[]float64{math.Sin(r), math.Cos(r), 0, 0},
		[]float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1},
	}
}
