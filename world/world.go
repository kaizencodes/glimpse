package world

import (
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/objects"
	"glimpse/ray"
	"glimpse/tuple"
)

type World struct {
	objects []objects.Object
	lights  []ray.Light
}

func (w *World) Objects() []objects.Object {
	return w.objects
}

func (w *World) Lights() []ray.Light {
	return w.lights
}

func (w *World) SetObjects(objs []objects.Object) {
	w.objects = objs
}

func (w *World) SetLights(lights []ray.Light) {
	w.lights = lights
}

func (w *World) ColorAt(r ray.Ray) color.Color {
	hit := w.intersect(r).Hit()
	if hit.Empty() {
		return color.Black()
	}

	return w.shadeHit(ray.PrepareComputations(hit, r))
}

func (w *World) intersect(r ray.Ray) ray.Intersections {
	coll := ray.Intersections{}
	for _, o := range w.objects {
		coll = append(coll, r.Intersect(o)...)
	}
	coll.Sort()

	return coll
}

func (w *World) shadeHit(comps ray.Computations) color.Color {
	isShadowed := w.shadowAt(comps.OverPoint())
	c := ray.Lighting(
		comps.Object().Material(),
		w.Lights()[0],
		comps.Point(),
		comps.EyeV(),
		comps.NormalV(),
		isShadowed,
	)
	for i, l := range w.Lights() {
		if i == 0 {
			continue
		}
		c = color.Add(c, ray.Lighting(
			comps.Object().Material(),
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

func Default() *World {
	o1 := objects.NewSphere()
	o1.SetMaterial(objects.NewMaterial(color.New(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0))
	o2 := objects.NewSphere()
	o2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	return &World{
		objects: []objects.Object{
			objects.Object(o1), objects.Object(o2),
		},
		lights: []ray.Light{
			ray.NewLight(tuple.NewPoint(-10, 10, -10), color.New(1, 1, 1)),
		},
	}
}

func New(objects []objects.Object, lights []ray.Light) *World {
	return &World{objects, lights}
}

func ViewTransformation(from, to, up tuple.Tuple) matrix.Matrix {
	forward := tuple.Subtract(to, from).Normalize()
	left := tuple.Cross(forward, up.Normalize())
	trueUp := tuple.Cross(left, forward)

	orientation := matrix.Matrix{
		[]float64{left.X(), left.Y(), left.Z(), 0},
		[]float64{trueUp.X(), trueUp.Y(), trueUp.Z(), 0},
		[]float64{-forward.X(), -forward.Y(), -forward.Z(), 0},
		[]float64{0, 0, 0, 1},
	}

	result, _ := matrix.Multiply(orientation, matrix.Translation(-from.X(), -from.Y(), -from.Z()))

	return result
}
