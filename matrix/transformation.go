package matrix

import "math"

func GetTranslation(x, y, z float64) Matrix {
	return Matrix{
		[]float64{1, 0, 0, x},
		[]float64{0, 1, 0, y},
		[]float64{0, 0, 1, z},
		[]float64{0, 0, 0, 1},
	}
}

func GetScaling(x, y, z float64) Matrix {
	return Matrix{
		[]float64{x, 0, 0, 0},
		[]float64{0, y, 0, 0},
		[]float64{0, 0, z, 0},
		[]float64{0, 0, 0, 1},
	}
}

func GetRotationX(rad float64) Matrix {
	return Matrix{
		[]float64{1, 0, 0, 0},
		[]float64{0, math.Cos(rad), -math.Sin(rad), 0},
		[]float64{0, math.Sin(rad), math.Cos(rad), 0},
		[]float64{0, 0, 0, 1},
	}
}

func GetRotationY(rad float64) Matrix {
	return Matrix{
		[]float64{math.Cos(rad), 0, math.Sin(rad), 0},
		[]float64{0, 1, 0, 0},
		[]float64{-math.Sin(rad), 0, math.Cos(rad), 0},
		[]float64{0, 0, 0, 1},
	}
}

func GetRotationZ(rad float64) Matrix {
	return Matrix{
		[]float64{math.Cos(rad), -math.Sin(rad), 0, 0},
		[]float64{math.Sin(rad), math.Cos(rad), 0, 0},
		[]float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1},
	}
}

func GetShearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		[]float64{1, xy, xz, 0},
		[]float64{yx, 1, yz, 0},
		[]float64{zx, zy, 1, 0},
		[]float64{0, 0, 0, 1},
	}
}
