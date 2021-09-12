package ray

import (
	"fmt"
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/objects"
	"glimpse/tuple"
	"math"
	"strconv"
)

type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

type Intersection struct {
	t      float64
	object *objects.Sphere
}

type Intersections []Intersection

func New(origin, direction tuple.Tuple) Ray {
	return Ray{origin, direction}
}

func (r Ray) Position(dist float64) tuple.Tuple {
	return tuple.Add(r.origin, r.direction.Scalar(dist))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(origin: %s, direction: %s)", r.origin, r.direction)
}

func (r Ray) Equal(other Ray) bool {
	return r.origin.Equal(other.origin) && r.direction.Equal(other.direction)
}

func (c Intersections) String() string {
	var result string

	for _, section := range c {
		result += strconv.FormatFloat(section.t, 'f', -1, 64) + ", "
	}
	return result
}

func Intersect(r Ray, s *objects.Sphere) Intersections {
	transform, err := s.GetTransform().Inverse()
	if err != nil {
		panic(err)
	}
	origin, _ := tuple.Multiply(transform, r.origin)
	direction, _ := tuple.Multiply(transform, r.direction)
	ray2 := Ray{origin, direction}

	sphere_to_ray := tuple.Subtract(ray2.origin, tuple.NewPoint(0, 0, 0))

	a := tuple.Dot(ray2.direction, ray2.direction)
	b := 2 * tuple.Dot(ray2.direction, sphere_to_ray)
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

func (r Ray) Translate(x, y, z float64) Ray {
	origin, err := tuple.Multiply(matrix.GetTranslation(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	return Ray{origin: origin, direction: r.direction}
}

func (r Ray) Scale(x, y, z float64) Ray {
	origin, err := tuple.Multiply(matrix.GetScaling(x, y, z), r.origin)
	if err != nil {
		panic(err)
	}
	direction, err := tuple.Multiply(matrix.GetScaling(x, y, z), r.direction)
	if err != nil {
		panic(err)
	}
	return Ray{origin: origin, direction: direction}
}

func (r Ray) GetOrigin() tuple.Tuple {
	return r.origin
}

func (r Ray) GetDirection() tuple.Tuple {
	return r.direction
}

func (inter Intersection) Empty() bool {
	return inter.t == math.MaxFloat64
}

func (hit Intersection) GetT() float64 {
	return hit.t
}

func (hit Intersection) GetObject() *objects.Sphere {
	return hit.object
}

type PointLight struct {
	position  tuple.Tuple
	intensity color.Color
}

func NewPointLight(position tuple.Tuple, intensity color.Color) PointLight {
	return PointLight{position, intensity}
}

func (l PointLight) String() string {
	return fmt.Sprintf("PointLight(position: %f, intensity: %f)", l.position, l.intensity)
}

func Lighting(mat objects.Material, light PointLight, point, eyeV, normalV tuple.Tuple) color.Color {
	effectiveColor := color.HadamardProduct(mat.GetColor(), light.intensity)
	lightV := tuple.Subtract(light.position, point).Normalize()
	ambient := effectiveColor.Scalar(mat.GetAmbient())
	lightDotNormal := tuple.Dot(lightV, normalV)
	var diffuse, specular color.Color
	if lightDotNormal < 0 {
		diffuse = color.Black()
		specular = color.Black()
	} else {
		diffuse = effectiveColor.Scalar(mat.GetDiffuse() * lightDotNormal)
		reflectV := tuple.Reflect(lightV.Negate(), normalV)
		reflectDotEye := tuple.Dot(reflectV, eyeV)
		if reflectDotEye <= 0 {
			specular = color.Black()
		} else {
			factor := math.Pow(reflectDotEye, mat.GetShininess())
			specular = light.intensity.Scalar(mat.GetSpecular() * factor)
		}
	}

	return color.Add(color.Add(ambient, diffuse), specular)
}
