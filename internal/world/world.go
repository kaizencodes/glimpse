package world

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/light"
	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/renderer"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type World struct {
	shapes []shapes.Shape
	lights []light.Light
}

func (w *World) Shapes() []shapes.Shape {
	return w.shapes
}

func (w *World) Lights() []light.Light {
	return w.lights
}

func (w *World) SetShapes(shapes []shapes.Shape) {
	w.shapes = shapes
}

func (w *World) SetLights(lights []light.Light) {
	w.lights = lights
}

func (w *World) ColorAt(r *ray.Ray) color.Color {
	intersections := w.intersect(r)
	hit := intersections.Hit()
	if hit.Empty() {
		return color.Black()
	}

	return w.shadeHit(renderer.PrepareComputations(hit, r, intersections))
}

func (w *World) intersect(r *ray.Ray) shapes.Intersections {
	coll := shapes.Intersections{}
	for _, o := range w.shapes {
		coll = append(coll, shapes.Intersect(o, r)...)
	}
	coll.Sort()

	return coll
}

func (w *World) shadeHit(comps renderer.Computations) color.Color {
	isShadowed := w.shadowAt(comps.OverPoint())
	c := light.Lighting(
		comps.Shape(),
		w.Lights()[0],
		comps.OverPoint(),
		comps.EyeV(),
		comps.NormalV(),
		isShadowed,
	)
	reflected := w.reflectedColor(comps)
	refracted := w.refractedColor(comps)
	mat := comps.Shape().Material()
	if mat.Reflective > 0 && mat.Transparency > 0 {
		reflectance := comps.Schlick()
		reflected = reflected.Scalar(reflectance)
		refracted = refracted.Scalar(1 - reflectance)
	}
	c = color.Add(c, reflected)
	c = color.Add(c, refracted)

	for i, l := range w.Lights() {
		if i == 0 {
			continue
		}
		c = color.Add(c, light.Lighting(
			comps.Shape(),
			l,
			comps.OverPoint(),
			comps.EyeV(),
			comps.NormalV(),
			isShadowed))
	}

	return c
}

func (w *World) shadowAt(point tuple.Tuple) bool {
	for _, l := range w.lights {
		v := tuple.Subtract(l.Position(), point)
		dist := v.Magnitude()
		r := ray.NewRay(point, v.Normalize())
		hit := w.intersect(r).Hit()

		if !hit.Empty() && hit.T() < dist {
			return true
		}
	}
	return false
}

func (w *World) reflectedColor(comps renderer.Computations) color.Color {
	if comps.Shape().Material().Reflective == 0 || comps.BounceLimit() < 1 {
		return color.Black()
	}

	r := ray.NewRay(comps.OverPoint(), comps.ReflectV())
	r.SetBounceLimit(comps.BounceLimit() - 1)
	c := w.ColorAt(r)

	return c.Scalar(comps.Shape().Material().Reflective)
}

func (w *World) refractedColor(comps renderer.Computations) color.Color {
	if comps.Shape().Material().Transparency == 0 || comps.BounceLimit() < 1 {
		return color.Black()
	}

	// Find the ration of first index of refraction to the second.
	nRatio := comps.N1() / comps.N2()
	// cos(theta i) is the same as the dot product of the two vectors.
	cosI := tuple.Dot(comps.EyeV(), comps.NormalV())
	// Find sin(theta t)^2 via trigonometric identity
	sin2T := math.Pow(nRatio, 2) * (1 - math.Pow(cosI, 2))
	if sin2T > 1 {
		return color.Black()
	}

	// Find cos(theta t) via trigonometric identity
	cosT := math.Sqrt(1.0 - sin2T)

	// Compute the direction of the refracted ray.
	direction := tuple.Subtract(comps.NormalV().Scalar((nRatio*cosI)-cosT), comps.EyeV().Scalar(nRatio))
	refactedRay := ray.NewRay(comps.UnderPoint(), direction)
	refactedRay.SetBounceLimit(comps.BounceLimit() - 1)

	// Find the color of the refracted ray, making sure to multiply by the transparency
	// value to account for any opacity.
	return w.ColorAt(refactedRay).Scalar(comps.Shape().Material().Transparency)
}

func Default() *World {
	o1 := shapes.NewSphere()
	o1.SetMaterial(materials.NewMaterial(color.New(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0, 0, 0, 1))
	o2 := shapes.NewSphere()
	o2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	return &World{
		shapes: []shapes.Shape{
			shapes.Shape(o1), shapes.Shape(o2),
		},
		lights: []light.Light{
			light.NewLight(tuple.NewPoint(-10, 10, -10), color.New(1, 1, 1)),
		},
	}
}

func New(shapes []shapes.Shape, lights []light.Light) *World {
	return &World{shapes, lights}
}
