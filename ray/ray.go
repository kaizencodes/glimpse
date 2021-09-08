package ray

import (
	"fmt"
	"glimpse/tuple"
	"math"
	"strconv"
)

type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

type Sphere struct {
	center tuple.Tuple
	radius float64
}

type Intersection struct {
	t      float64
	object *Sphere
}

type Intersections []Intersection

func New(origin, direction tuple.Tuple) Ray {
	return Ray{origin, direction}
}

func NewShpere() Sphere {
	return Sphere{center: tuple.NewPoint(0, 0, 0), radius: 1}
}

func (r Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.origin, r.direction.Scalar(dist))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(origin: %s, direction: %s)", r.origin, r.direction)
}

func (s *Sphere) String() string {
	return fmt.Sprintf("Shpere(center: %s, radius: %f)", s.center, s.radius)
}

func (c Intersections) String() string {
	var result string

	for _, section := range c {
		result += strconv.FormatFloat(section.t, 'f', -1, 64) + ", "
	}
	return result
}

func Intersect(r Ray, s *Sphere) Intersections {
	sphere_to_ray := tuple.Subtract(r.origin, tuple.NewPoint(0, 0, 0))

	a := tuple.Dot(r.direction, r.direction)
	b := 2 * tuple.Dot(r.direction, sphere_to_ray)
	c := tuple.Dot(sphere_to_ray, sphere_to_ray) - 1

	disciminant := math.Pow(b, 2) - 4*a*c

	if disciminant < 0 {
		return Intersections{}
	}

	t1 := (-b - math.Sqrt(disciminant)) / (2 * a)
	t2 := (-b + math.Sqrt(disciminant)) / (2 * a)

	return Intersections{Intersection{t: t1, object: s}, Intersection{t: t2, object: s}}
}

func Hit(coll Intersections) Intersection {
	res := Intersection{t: math.MaxFloat64}
	for _, val := range coll {
		if val.t < 0 {
			continue
		}
		if val.t < res.t {
			res = val
		}
	}
	return res
}
