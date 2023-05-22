package tuple

import "github.com/kaizencodes/glimpse/internal/matrix"

func (t Tuple) Translate(x, y, z float64) Tuple {
	return Multiply(matrix.Translation(x, y, z), t)
}

func (t Tuple) Scale(x, y, z float64) Tuple {
	return Multiply(matrix.Scaling(x, y, z), t)
}

func (t Tuple) RotateX(r float64) Tuple {
	return Multiply(matrix.RotationX(r), t)
}

func (t Tuple) RotateY(r float64) Tuple {
	return Multiply(matrix.RotationY(r), t)
}

func (t Tuple) RotateZ(r float64) Tuple {
	return Multiply(matrix.RotationZ(r), t)
}

func (t Tuple) Shear(xy, xz, yx, yz, zx, zy float64) Tuple {
	return Multiply(matrix.Shearing(xy, xz, yx, yz, zx, zy), t)
}
