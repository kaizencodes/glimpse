package world

import (
	"glimpse/color"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/ray"
	"glimpse/shapes"
	"glimpse/tuple"
)

type World struct {
	shapes []shapes.Shape
	lights []ray.Light
}

func (w *World) Shapes() []shapes.Shape {
	return w.shapes
}

func (w *World) Lights() []ray.Light {
	return w.lights
}

func (w *World) SetShapes(shapes []shapes.Shape) {
	w.shapes = shapes
}

func (w *World) SetLights(lights []ray.Light) {
	w.lights = lights
}

func (w *World) ColorAt(r *ray.Ray) color.Color {
	hit := w.intersect(r).Hit()
	if hit.Empty() {
		return color.Black()
	}

	return w.shadeHit(ray.PrepareComputations(hit, r))
}

func (w *World) intersect(r *ray.Ray) ray.Intersections {
	coll := ray.Intersections{}
	for _, o := range w.shapes {
		coll = append(coll, r.Intersect(o)...)
	}
	coll.Sort()

	return coll
}

func (w *World) shadeHit(comps ray.Computations) color.Color {
	isShadowed := w.shadowAt(comps.OverPoint())
	c := ray.Lighting(
		comps.Shape(),
		w.Lights()[0],
		comps.Point(),
		comps.EyeV(),
		comps.NormalV(),
		isShadowed,
	)
	c = color.Add(c, w.reflectedColor(comps))

	for i, l := range w.Lights() {
		if i == 0 {
			continue
		}
		c = color.Add(c, ray.Lighting(
			comps.Shape(),
			l,
			comps.Point(),
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
		r := ray.New(point, v.Normalize())
		hit := w.intersect(r).Hit()

		if !hit.Empty() && hit.T() < dist {
			return true
		}
	}
	return false
}

func (w *World) reflectedColor(comps ray.Computations) color.Color {
	if comps.Shape().Material().Reflective() == 0 || comps.BounceLimit() < 1 {
		return color.Black()
	}

	r := ray.New(comps.OverPoint(), comps.ReflectV())
	r.SetBounceLimit(comps.BounceLimit() - 1)
	c := w.ColorAt(r)

	return c.Scalar(comps.Shape().Material().Reflective())
}

func Default() *World {
	o1 := shapes.NewSphere()
	o1.SetMaterial(materials.NewMaterial(color.New(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0, 0))
	o2 := shapes.NewSphere()
	o2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	return &World{
		shapes: []shapes.Shape{
			shapes.Shape(o1), shapes.Shape(o2),
		},
		lights: []ray.Light{
			ray.NewLight(tuple.NewPoint(-10, 10, -10), color.New(1, 1, 1)),
		},
	}
}

func New(shapes []shapes.Shape, lights []ray.Light) *World {
	return &World{shapes, lights}
}
