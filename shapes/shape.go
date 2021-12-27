package shapes

import (
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Shape interface {
	Material() *Material
	SetMaterial(m *Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	LocalNormalAt(point tuple.Tuple) tuple.Tuple
}

func NormalAt(worldPoint tuple.Tuple, shape Shape) tuple.Tuple {
	inv_mat, err := shape.Transform().Inverse()
	if err != nil {
		panic(err)
	}
	localPoint, _ := tuple.Multiply(inv_mat, worldPoint)
	localNormal := shape.LocalNormalAt(localPoint)
	worldNormal, _ := tuple.Multiply(inv_mat.Transpose(), localNormal)
	return worldNormal.ToVector().Normalize()
}

func ColorAt(worldPoint tuple.Tuple, shape Shape) color.Color {
	invShapeTransform, err := shape.Transform().Inverse()
	if err != nil {
		panic(err)
	}

	objectPoint, err := tuple.Multiply(invShapeTransform, worldPoint)
	if err != nil {
		panic(err)
	}

	invPatternTransform, err := shape.Material().Transform().Inverse()
	if err != nil {
		panic(err)
	}

	patternPoint, err := tuple.Multiply(invPatternTransform, objectPoint)
	if err != nil {
		panic(err)
	}

	return shape.Material().ColorAt(patternPoint)
}
