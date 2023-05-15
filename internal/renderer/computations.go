package renderer

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/calc"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type Computations struct {
	t, n1, n2                                             float64
	shape                                                 shapes.Shape
	point, eyeV, normalV, reflectV, overPoint, underPoint tuple.Tuple
	bounceLimit                                           int
	inside                                                bool
}

func (c Computations) T() float64 {
	return c.t
}

func (c Computations) N1() float64 {
	return c.n1
}

func (c Computations) N2() float64 {
	return c.n2
}

func (c Computations) Shape() shapes.Shape {
	return c.shape
}

func (c Computations) Point() tuple.Tuple {
	return c.point
}

func (c Computations) EyeV() tuple.Tuple {
	return c.eyeV
}

func (c Computations) NormalV() tuple.Tuple {
	return c.normalV
}

func (c Computations) ReflectV() tuple.Tuple {
	return c.reflectV
}

func (c Computations) BounceLimit() int {
	return c.bounceLimit
}

func (c Computations) OverPoint() tuple.Tuple {
	return c.overPoint
}

func (c Computations) UnderPoint() tuple.Tuple {
	return c.underPoint
}

func (c Computations) Inside() bool {
	return c.inside
}

func (comps Computations) Schlick() float64 {
	// find the cosine of the angle between the eye and normal vectors
	cos := tuple.Dot(comps.eyeV, comps.normalV)
	// total internal reflection can only occur if n1 > n2
	if comps.n1 > comps.n2 {
		n := comps.n1 / comps.n2
		sin2T := math.Pow(n, 2) * (1.0 - math.Pow(cos, 2))
		if sin2T > 1.0 {
			return 1.0
		}

		// compute cosine of theta_t using trigonometric identity
		cosT := math.Sqrt(1.0 - sin2T)
		// when n1 > n2, use cos(theta_t) instead
		cos = cosT
	}

	r0 := math.Pow(((comps.n1 - comps.n2) / (comps.n1 + comps.n2)), 2)

	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

func PrepareComputations(hit shapes.Intersection, r *ray.Ray, xs shapes.Intersections) Computations {
	point := r.Position(hit.T())
	normalV := shapes.NormalAt(point, hit.Shape(), hit)
	eyeV := r.Direction().Negate()

	inside := false
	if tuple.Dot(normalV, eyeV) < 0 {
		inside = true
		normalV = normalV.Negate()
	}
	overPoint := tuple.Add(point, normalV.Scalar(calc.EPSILON))
	underPoint := tuple.Subtract(point, normalV.Scalar(calc.EPSILON))

	container := []shapes.Shape{}
	n1, n2 := 1.0, 1.0
	for _, val := range xs {
		if val == hit {
			if len(container) == 0 {
				n1 = 1.0
			} else {
				n1 = container[len(container)-1].Material().RefractiveIndex()
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
				n2 = container[len(container)-1].Material().RefractiveIndex()
			}
			break
		}
	}
	return Computations{
		t:           hit.T(),
		shape:       hit.Shape(),
		point:       point,
		eyeV:        eyeV,
		normalV:     normalV,
		reflectV:    tuple.Reflect(r.Direction(), normalV),
		inside:      inside,
		overPoint:   overPoint,
		underPoint:  underPoint,
		bounceLimit: r.BounceLimit(),
		n1:          n1,
		n2:          n2,
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
