package renderer

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

type Computations struct {
	T, N1, N2                                             float64
	Shape                                                 shapes.Shape
	Point, EyeV, NormalV, ReflectV, OverPoint, UnderPoint tuple.Tuple
	BounceLimit                                           int
	Inside                                                bool
}

func (comps Computations) Schlick() float64 {
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

func PrepareComputations(hit shapes.Intersection, r *ray.Ray, xs shapes.Intersections) Computations {
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

	container := []shapes.Shape{}
	n1, n2 := 1.0, 1.0
	for _, val := range xs {
		if val == hit {
			if len(container) == 0 {
				n1 = 1.0
			} else {
				n1 = container[len(container)-1].Material().RefractiveIndex
			}
		}

		ok, at := contains(container, val.Shape())
		if ok {
			container = remove(container, at)
		} else {
			container = append(container, val.Shape())
		}

		if val == hit {
			if len(container) == 0 {
				n2 = 1.0
			} else {
				n2 = container[len(container)-1].Material().RefractiveIndex
			}
			break
		}
	}
	return Computations{
		T:           hit.T(),
		Shape:       hit.Shape(),
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

func contains(collection []shapes.Shape, shape shapes.Shape) (bool, int) {

	for i, e := range collection {
		if shape == e {
			return true, i
		}
	}
	return false, math.MaxInt
}

func remove(collection []shapes.Shape, i int) []shapes.Shape {
	collection[i] = collection[len(collection)-1]
	return collection[:len(collection)-1]
}
