package matrix

import "math"

func Translation(x, y, z float64) Matrix {
	return Matrix{
		[]float64{1, 0, 0, x},
		[]float64{0, 1, 0, y},
		[]float64{0, 0, 1, z},
		[]float64{0, 0, 0, 1},
	}
}

func Scaling(x, y, z float64) Matrix {
	return Matrix{
		[]float64{x, 0, 0, 0},
		[]float64{0, y, 0, 0},
		[]float64{0, 0, z, 0},
		[]float64{0, 0, 0, 1},
	}
}

func RotationX(rad float64) Matrix {
	return Matrix{
		[]float64{1, 0, 0, 0},
		[]float64{0, math.Cos(rad), -math.Sin(rad), 0},
		[]float64{0, math.Sin(rad), math.Cos(rad), 0},
		[]float64{0, 0, 0, 1},
	}
}

func RotationY(rad float64) Matrix {
	return Matrix{
		[]float64{math.Cos(rad), 0, math.Sin(rad), 0},
		[]float64{0, 1, 0, 0},
		[]float64{-math.Sin(rad), 0, math.Cos(rad), 0},
		[]float64{0, 0, 0, 1},
	}
}

func RotationZ(rad float64) Matrix {
	return Matrix{
		[]float64{math.Cos(rad), -math.Sin(rad), 0, 0},
		[]float64{math.Sin(rad), math.Cos(rad), 0, 0},
		[]float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1},
	}
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		[]float64{1, xy, xz, 0},
		[]float64{yx, 1, yz, 0},
		[]float64{zx, zy, 1, 0},
		[]float64{0, 0, 0, 1},
	}
}
