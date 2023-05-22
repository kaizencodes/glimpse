package renderer

import (
	"math"
	"sync"

	"github.com/kaizencodes/glimpse/internal/camera"
	"github.com/kaizencodes/glimpse/internal/canvas"
	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/light"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/scenes"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

func Render(c *camera.Camera, w *scenes.Scene) canvas.Canvas {
	img := canvas.New(c.Width, c.Height)
	var wg sync.WaitGroup

	for y := 0; y < c.Height-1; y++ {
		for x := 0; x < c.Width-1; x++ {
			wg.Add(1)

			go func(x, y int) {
				defer wg.Done()

				r := c.RayForPixel(x, y)
				col := colorAt(w, r)
				img[x][y] = col
			}(x, y)
		}
	}
	wg.Wait()

	return img
}

func colorAt(scene *scenes.Scene, r *ray.Ray) color.Color {
	intersections := intersect(scene, r)
	hit := intersections.Hit()
	if hit.Empty() {
		return color.Black()
	}

	return shadeHit(scene, PrepareComputations(hit, r, intersections))
}

func intersect(scene *scenes.Scene, r *ray.Ray) shapes.Intersections {
	coll := shapes.Intersections{}
	for _, o := range scene.Shapes {
		coll = append(coll, shapes.Intersect(o, r)...)
	}
	coll.Sort()

	return coll
}

func shadeHit(scene *scenes.Scene, comps Computations) color.Color {
	isShadowed := shadowAt(scene, comps.OverPoint)
	c := light.Lighting(
		comps.Shape,
		scene.Lights[0],
		comps.OverPoint,
		comps.EyeV,
		comps.NormalV,
		isShadowed,
	)
	reflected := reflectedColor(scene, comps)
	refracted := refractedColor(scene, comps)
	mat := comps.Shape.Material()
	if mat.Reflective > 0 && mat.Transparency > 0 {
		reflectance := comps.Schlick()
		reflected = reflected.Scalar(reflectance)
		refracted = refracted.Scalar(1 - reflectance)
	}
	c = color.Add(c, reflected)
	c = color.Add(c, refracted)

	for i, l := range scene.Lights {
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

func shadowAt(scene *scenes.Scene, point tuple.Tuple) bool {
	for _, l := range scene.Lights {
		v := tuple.Subtract(l.Position(), point)
		dist := v.Magnitude()
		r := ray.NewRay(point, v.Normalize())
		hit := intersect(scene, r).Hit()

		if !hit.Empty() && hit.T() < dist {
			return true
		}
	}
	return false
}

func reflectedColor(scene *scenes.Scene, comps Computations) color.Color {
	if comps.Shape.Material().Reflective == 0 || comps.BounceLimit < 1 {
		return color.Black()
	}

	r := ray.NewRay(comps.OverPoint, comps.ReflectV)
	r.BounceLimit = comps.BounceLimit - 1
	c := colorAt(scene, r)

	return c.Scalar(comps.Shape.Material().Reflective)
}

func refractedColor(scene *scenes.Scene, comps Computations) color.Color {
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
	return colorAt(scene, refractedRay).Scalar(comps.Shape.Material().Transparency)
}
