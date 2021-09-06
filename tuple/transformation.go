package tuple

import "glimpse/matrix"

func (t Tuple) Translate(x, y, z float64) Tuple {
	result, err := Multiply(matrix.GetTranslation(x, y, z), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) Scale(x, y, z float64) Tuple {
	result, err := Multiply(matrix.GetScaling(x, y, z), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) RotateX(r float64) Tuple {
	result, err := Multiply(matrix.GetRotationX(r), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) RotateY(r float64) Tuple {
	result, err := Multiply(matrix.GetRotationY(r), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) RotateZ(r float64) Tuple {
	result, err := Multiply(matrix.GetRotationZ(r), t)
	if err != nil {
		panic(err)
	}
	return result
}

func (t Tuple) Shear(xy, xz, yx, yz, zx, zy float64) Tuple {
	result, err := Multiply(matrix.GetShearing(xy, xz, yx, yz, zx, zy), t)
	if err != nil {
		panic(err)
	}
	return result
}
