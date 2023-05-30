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
	localIntersect(r *ray.Ray) Intersections
	Parent() Shape
	SetParent(Shape)
}

func ColorAt(scenePoint tuple.Tuple, shape Shape) color.Color {
	invShapeTransform := shape.Transform().Inverse()
	objectPoint := tuple.Multiply(invShapeTransform, scenePoint)
	invPatternTransform := shape.Material().Transform().Inverse()
	patternPoint := tuple.Multiply(invPatternTransform, objectPoint)

	return shape.Material().ColorAt(patternPoint)
}

func Intersect(s Shape, r *ray.Ray) Intersections {
	transform := s.Transform().Inverse()
	origin := tuple.Multiply(transform, r.Origin)
	direction := tuple.Multiply(transform, r.Direction)
	localRay := ray.NewRay(origin, direction)

	return s.localIntersect(localRay)
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

	inverse := s.Transform().Inverse()
	result := tuple.Multiply(inverse, p)
	return result
}

func normalToScene(v tuple.Tuple, s Shape) tuple.Tuple {
	transposed := s.Transform().Inverse().Transpose()
	normal := tuple.Multiply(transposed, v).ToVector().Normalize()

	if s.Parent() != nil {
		normal = normalToScene(normal, s.Parent())
	}

	return normal
}
