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

func GetRotationX(r float64) Matrix {
	return Matrix{
		[]float64{1, 0, 0, 0},
		[]float64{0, math.Cos(r), -math.Sin(r), 0},
		[]float64{0, math.Sin(r), math.Cos(r), 0},
		[]float64{0, 0, 0, 1},
	}
}

func GetRotationY(r float64) Matrix {
	return Matrix{
		[]float64{math.Cos(r), 0, math.Sin(r), 0},
		[]float64{0, 1, 0, 0},
		[]float64{-math.Sin(r), 0, math.Cos(r), 0},
		[]float64{0, 0, 0, 1},
	}
}

func GetRotationZ(r float64) Matrix {
	return Matrix{
		[]float64{math.Cos(r), -math.Sin(r), 0, 0},
		[]float64{math.Sin(r), math.Cos(r), 0, 0},
		[]float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1},
	}
}
