package shapes

import (
	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type Shape interface {
	Material() *materials.Material
	SetMaterial(m *materials.Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	LocalNormalAt(point tuple.Tuple, hit Intersection) tuple.Tuple
	LocalIntersect(r *ray.Ray) Intersections
	Parent() Shape
	SetParent(Shape)
}

func ColorAt(scenePoint tuple.Tuple, shape Shape) color.Color {
	invShapeTransform, err := shape.Transform().Inverse()
	if err != nil {
		panic(err)
	}

	objectPoint, err := tuple.Multiply(invShapeTransform, scenePoint)
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

func Intersect(s Shape, r *ray.Ray) Intersections {
	transform, err := s.Transform().Inverse()
	if err != nil {
		panic(err)
	}
	origin, _ := tuple.Multiply(transform, r.Origin)
	direction, _ := tuple.Multiply(transform, r.Direction)
	localRay := ray.NewRay(origin, direction)

	return s.LocalIntersect(localRay)
}

func NormalAt(scenePoint tuple.Tuple, shape Shape, hit Intersection) tuple.Tuple {
	localPoint := sceneToObject(scenePoint, shape)
	localNormal := shape.LocalNormalAt(localPoint, hit)
	return normalToScene(localNormal, shape)
}

func sceneToObject(p tuple.Tuple, s Shape) tuple.Tuple {
	if s.Parent() != nil {
		p = sceneToObject(p, s.Parent())
	}

	inverse, _ := s.Transform().Inverse()
	result, _ := tuple.Multiply(inverse, p)
	return result
}

func normalToScene(v tuple.Tuple, s Shape) tuple.Tuple {
	inv, _ := s.Transform().Inverse()
	normal, _ := tuple.Multiply(inv.Transpose(), v)
	normal = normal.ToVector().Normalize()

	if s.Parent() != nil {
		normal = normalToScene(normal, s.Parent())
	}

	return normal
}
