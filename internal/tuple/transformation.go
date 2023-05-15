package tuple

import "github.com/kaizencodes/glimpse/internal/matrix"

func (t Tuple) Translate(x, y, z float64) Tuple {
	result, err := Multiply(matrix.Translation(x, y, z), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) Scale(x, y, z float64) Tuple {
	result, err := Multiply(matrix.Scaling(x, y, z), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) RotateX(r float64) Tuple {
	result, err := Multiply(matrix.RotationX(r), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) RotateY(r float64) Tuple {
	result, err := Multiply(matrix.RotationY(r), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) RotateZ(r float64) Tuple {
	result, err := Multiply(matrix.RotationZ(r), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) Shear(xy, xz, yx, yz, zx, zy float64) Tuple {
	result, err := Multiply(matrix.Shearing(xy, xz, yx, yz, zx, zy), t)
	if err != nil {
		panic(err)
	}
	return result
}
