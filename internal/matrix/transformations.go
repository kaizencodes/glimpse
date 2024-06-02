package matrix

import "math"

// Translation moves the object in 3D space.
func Translation(x, y, z float64) Matrix {
	return Matrix{
		data: [16]float64{
			1, 0, 0, x,
			0, 1, 0, y,
			0, 0, 1, z,
			0, 0, 0, 1,
		},
		row_size: 4,
		col_size: 4,
	}
}

// // Scaling scales the object in 3D space.
// // Scaling by a positive > 1 value will enlarge the object.
// // Scaling by a positive < 1 value will shrink the object.
// // Scaling by a negative value will flip the object.
func Scaling(x, y, z float64) Matrix {
	return Matrix{
		data: [16]float64{
			x, 0, 0, 0,
			0, y, 0, 0,
			0, 0, z, 0,
			0, 0, 0, 1,
		},
		row_size: 4,
		col_size: 4,
	}
}

// // RotationX rotates the object around the x-axis.
func RotationX(rad float64) Matrix {
	return Matrix{
		data: [16]float64{
			1, 0, 0, 0,
			0, math.Cos(rad), -math.Sin(rad), 0,
			0, math.Sin(rad), math.Cos(rad), 0,
			0, 0, 0, 1,
		},
		row_size: 4,
		col_size: 4,
	}
}

// // RotationY rotates the object around the y-axis.
func RotationY(rad float64) Matrix {
	return Matrix{
		data: [16]float64{
			math.Cos(rad), 0, math.Sin(rad), 0,
			0, 1, 0, 0,
			-math.Sin(rad), 0, math.Cos(rad), 0,
			0, 0, 0, 1,
		},
		row_size: 4,
		col_size: 4,
	}
}

// // RotationZ rotates the object around the z-axis.
func RotationZ(rad float64) Matrix {
	return Matrix{
		data: [16]float64{
			math.Cos(rad), -math.Sin(rad), 0, 0,
			math.Sin(rad), math.Cos(rad), 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		},
		row_size: 4,
		col_size: 4,
	}
}

// // Shearing skews the object in 3D space.
func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		data: [16]float64{
			1, xy, xz, 0,
			yx, 1, yz, 0,
			zx, zy, 1, 0,
			0, 0, 0, 1,
		},
		row_size: 4,
		col_size: 4,
	}
}

var Identity = Matrix{
	[16]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	},
	4, 4,
}

// DefaultTransform returns the identity matrix.
// Multiplying by the identity matrix does not change the object.
func DefaultTransform() Matrix {
	return Identity
}
