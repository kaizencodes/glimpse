package tuple

import "glimpse/matrix"

func (t Tuple) Translate(x, y, z float64) (Tuple, error) {
	return Multiply(matrix.GetTranslation(x, y, z), t)
}

func (t Tuple) Scale(x, y, z float64) (Tuple, error) {
	return Multiply(matrix.GetScaling(x, y, z), t)
}

func (t Tuple) RotateX(r float64) (Tuple, error) {
	return Multiply(matrix.GetRotationX(r), t)
}

func (t Tuple) RotateY(r float64) (Tuple, error) {
	return Multiply(matrix.GetRotationY(r), t)
}

func (t Tuple) RotateZ(r float64) (Tuple, error) {
	return Multiply(matrix.GetRotationZ(r), t)
}
