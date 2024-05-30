package renderer

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

type Computations struct {
	T      float64 // The distance from the ray's origin to the intersection.
	N1, N2 float64 // N1 and N2 are the names given to the refractive indices of the materials
	// on either side of a ray-object intersection, with N1 belonging to the material being exited,
	// N2 belonging to the material being entered.
	Shape shapes.Shape
	Point tuple.Tuple // Point where the intersection occurred between thew shape and the ray.
	EyeV  tuple.Tuple // the eye vector, pointing from Point to the origin of the ray (usually, where
	// the eye exists that, is the camera looking at the scene).
	NormalV     tuple.Tuple // The surface normal, a vector that is perpendicular to the surface at Point.
	ReflectV    tuple.Tuple // The reflection vector, pointing in the direction that incoming light would bounce, or reflect.
	OverPoint   tuple.Tuple // The point is slightly over the surface.
	UnderPoint  tuple.Tuple // The point is slightly under the surface.
	BounceLimit int         // The maximum number of bounces that can occur.
	Inside      bool        // true if the intersection occurred from the inside of the shape.
}

// An approximation to Fresnel’s equations by Christophe Schlick.
func (comps Computations) schlick() float64 {
	// find the cosine of the angle between the eye and normal vectors
	cos := tuple.Dot(comps.EyeV, comps.NormalV)
	// total internal reflection can only occur if n1 > n2
	if comps.N1 > comps.N2 {
		n := comps.N1 / comps.N2
		sin2T := math.Pow(n, 2) * (1.0 - math.Pow(cos, 2))
		if sin2T > 1.0 {
			return 1.0
		}

		// compute cosine of theta_t using trigonometric identity
		cosT := math.Sqrt(1.0 - sin2T)
		// when n1 > n2, use cos(theta_t) instead
		cos = cosT
	}

	r0 := math.Pow(((comps.N1 - comps.N2) / (comps.N1 + comps.N2)), 2)

	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

func prepareComputations(hit shapes.Intersection, r *ray.Ray, xs shapes.Intersections) Computations {
	point := r.Position(hit.T())
	normalV := shapes.NormalAt(point, hit.Shape(), hit)
	eyeV := r.Direction.Negate()

	inside := false
	if tuple.Dot(normalV, eyeV) < 0 {
		inside = true
		normalV = normalV.Negate()
	}
	// after computing and (if appropriate) negating the normal vector we move the point slightly
	// over and under the surface.
	overPoint := tuple.Add(point, normalV.Scalar(utils.EPSILON))
	underPoint := tuple.Subtract(point, normalV.Scalar(utils.EPSILON))

	// contains objects encountered but not yet exited.
	container := []shapes.Shape{}
	n1, n2 := 1.0, 1.0
	for i := 0; i < len(xs); i++ {
		if xs[i] == hit {
			if len(container) == 0 {
				n1 = 1.0
			} else {
				// setting n1 to the last element in the container
				n1 = container[len(container)-1].Material().RefractiveIndex
			}
		}

		ok, at := contains(container, xs[i].Shape())
		if ok {
			container = remove(container, at)
		} else {
			container = append(container, xs[i].Shape())
		}

		if xs[i] == hit {
			if len(container) == 0 {
				n2 = 1.0
			} else {
				// // setting n2 to the last element in the container
				n2 = container[len(container)-1].Material().RefractiveIndex
			}
			break
		}
	}
	return Computations{
		T:           hit.T(),     // Copy the t value from the intersection for convenience.
		Shape:       hit.Shape(), // Copy the object from the intersection for convenience.
		Point:       point,
		EyeV:        eyeV,
		NormalV:     normalV,
		ReflectV:    tuple.Reflect(r.Direction, normalV),
		Inside:      inside,
		OverPoint:   overPoint,
		UnderPoint:  underPoint,
		BounceLimit: r.BounceLimit,
		N1:          n1,
		N2:          n2,
	}
}

// helper function to check if a shape is in a collection
func contains(collection []shapes.Shape, shape shapes.Shape) (bool, int) {
	for i := 0; i < len(collection); i++ {
		if shape == collection[i] {
			return true, i
		}
	}
	return false, math.MaxInt
}

// helper function to remove a shape from a collection
func remove(collection []shapes.Shape, i int) []shapes.Shape {
	collection[i] = collection[len(collection)-1]
	return collection[:len(collection)-1]
}
