package shapes

import (
	"glimpse/color"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Shape interface {
	Material() *materials.Material
	SetMaterial(m *materials.Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	LocalNormalAt(point tuple.Tuple) tuple.Tuple
	Parent() Shape
	SetParent(Shape)
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

func WorldToObject(s Shape, p tuple.Tuple) tuple.Tuple {
	if s.Parent() != nil {
		p = WorldToObject(s.Parent(), p)
	}

	inverse, _ := s.Transform().Inverse()
	result, _ := tuple.Multiply(inverse, p)
	return result
}

func NormalToWorld(s Shape, v tuple.Tuple) tuple.Tuple {
	inv, _ := s.Transform().Inverse()
	normal, _ := tuple.Multiply(inv.Transpose(), v)
	normal = normal.ToVector().Normalize()

	if s.Parent() != nil {
		normal = NormalToWorld(s.Parent(), normal)
	}

	return normal
}
