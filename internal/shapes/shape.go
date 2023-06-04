// This package contains 3d primitives that can be rendered.
// The following primitives are supported:
// - Sphere
// - Plane
// - Cube
// - Cylinder

package shapes

import (
	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

// Shape is an interface that defines the methods that all 3d shapes must implement.
type Shape interface {
	Material() *materials.Material
	SetMaterial(m *materials.Material)
	Transform() matrix.Matrix
	SetTransform(transform matrix.Matrix)
	localNormalAt(point tuple.Tuple, hit Intersection) tuple.Tuple
	localIntersect(r *ray.Ray) Intersections
	Parent() Shape
	SetParent(Shape)
}

func ColorAt(scenePoint tuple.Tuple, shape Shape) color.Color {
	// transform a point in scene(global) space to object(local) space
	objectPoint := sceneToObject(scenePoint, shape)
	invPatternTransform := shape.Material().Transform().Inverse()
	patternPoint := tuple.Multiply(invPatternTransform, objectPoint)

	return shape.Material().ColorAt(patternPoint)
}

func Intersect(s Shape, r *ray.Ray) Intersections {
	transform := s.Transform().Inverse()
	origin := tuple.Multiply(transform, r.Origin)
	direction := tuple.Multiply(transform, r.Direction)
	localRay := ray.New(origin, direction)

	return s.localIntersect(localRay)
}

// Calculates the normal vector on the surface of a shape at a given point (the hit).
func NormalAt(scenePoint tuple.Tuple, shape Shape, hit Intersection) tuple.Tuple {
	// transform a point in scene(global) space to object(local) space
	localPoint := sceneToObject(scenePoint, shape)
	// calculate the normal vector in object(local) space
	localNormal := shape.localNormalAt(localPoint, hit)
	// transform the normal vector in object(local) space to scene(global) space.
	return objectToScene(localNormal, shape)
}

func sceneToObject(p tuple.Tuple, s Shape) tuple.Tuple {
	if s.Parent() != nil {
		p = sceneToObject(p, s.Parent())
	}

	inverse := s.Transform().Inverse()
	result := tuple.Multiply(inverse, p)
	return result
}

func objectToScene(v tuple.Tuple, s Shape) tuple.Tuple {
	transposed := s.Transform().Inverse().Transpose()
	normal := tuple.Multiply(transposed, v).ToVector().Normalize()

	if s.Parent() != nil {
		normal = objectToScene(normal, s.Parent())
	}

	return normal
}
