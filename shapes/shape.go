package shapes

import (
	"glimpse/matrix"
	"glimpse/tuple"
)

type Shape interface {
	Material() Material
	SetMaterial(m Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	LocalNormalAt(point tuple.Tuple) tuple.Tuple
}

func DefaultTransform() matrix.Matrix {
	return matrix.NewIdentity(4)
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
