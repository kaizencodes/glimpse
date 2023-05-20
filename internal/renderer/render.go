package renderer

import (
	"math"

	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/light"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/scene"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

func ColorAt(w *scene.Scene, r *ray.Ray) color.Color {
	intersections := intersect(w, r)
	hit := intersections.Hit()
	if hit.Empty() {
		return color.Black()
	}

	return shadeHit(w, PrepareComputations(hit, r, intersections))
}

func intersect(w *scene.Scene, r *ray.Ray) shapes.Intersections {
	coll := shapes.Intersections{}
	for _, o := range w.Shapes {
		coll = append(coll, shapes.Intersect(o, r)...)
	}
	coll.Sort()

	return coll
}

func shadeHit(w *scene.Scene, comps Computations) color.Color {
	isShadowed := shadowAt(w, comps.OverPoint)
	c := light.Lighting(
		comps.Shape,
		w.Lights[0],
		comps.OverPoint,
		comps.EyeV,
		comps.NormalV,
		isShadowed,
	)
	reflected := reflectedColor(w, comps)
	refracted := refractedColor(w, comps)
	mat := comps.Shape.Material()
	if mat.Reflective > 0 && mat.Transparency > 0 {
		reflectance := comps.Schlick()
		reflected = reflected.Scalar(reflectance)
		refracted = refracted.Scalar(1 - reflectance)
	}
	c = color.Add(c, reflected)
	c = color.Add(c, refracted)

	for i, l := range w.Lights {
		if i == 0 {
			continue
		}
		c = color.Add(c, light.Lighting(
			comps.Shape,
			l,
			comps.OverPoint,
			comps.EyeV,
			comps.NormalV,
			isShadowed))
	}

	return c
}

func shadowAt(w *scene.Scene, point tuple.Tuple) bool {
	for _, l := range w.Lights {
		v := tuple.Subtract(l.Position(), point)
		dist := v.Magnitude()
		r := ray.NewRay(point, v.Normalize())
		hit := intersect(w, r).Hit()

		if !hit.Empty() && hit.T() < dist {
			return true
		}
	}
	return false
}

func reflectedColor(w *scene.Scene, comps Computations) color.Color {
	if comps.Shape.Material().Reflective == 0 || comps.BounceLimit < 1 {
		return color.Black()
	}

	r := ray.NewRay(comps.OverPoint, comps.ReflectV)
	r.BounceLimit = comps.BounceLimit - 1
	c := ColorAt(w, r)

	return c.Scalar(comps.Shape.Material().Reflective)
}

func refractedColor(w *scene.Scene, comps Computations) color.Color {
	if comps.Shape.Material().Transparency == 0 || comps.BounceLimit < 1 {
		return color.Black()
	}

	// Find the ration of first index of refraction to the second.
	nRatio := comps.N1 / comps.N2
	// cos(theta i) is the same as the dot product of the two vectors.
	cosI := tuple.Dot(comps.EyeV, comps.NormalV)
	// Find sin(theta t)^2 via trigonometric identity
	sin2T := math.Pow(nRatio, 2) * (1 - math.Pow(cosI, 2))
	if sin2T > 1 {
		return color.Black()
	}

	// Find cos(theta t) via trigonometric identity
	cosT := math.Sqrt(1.0 - sin2T)

	// Compute the direction of the refracted ray.
	direction := tuple.Subtract(comps.NormalV.Scalar((nRatio*cosI)-cosT), comps.EyeV.Scalar(nRatio))
	refractedRay := ray.NewRay(comps.UnderPoint, direction)
	refractedRay.BounceLimit = comps.BounceLimit - 1

	// Find the color of the refracted ray, making sure to multiply by the transparency
	// value to account for any opacity.
	return ColorAt(w, refractedRay).Scalar(comps.Shape.Material().Transparency)
}
