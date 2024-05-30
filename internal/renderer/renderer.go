// All rendering logic is contained in this package.
package renderer

import (
	"fmt"
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

// The main function that renders the scene pixel by pixel.
func Render(c *camera.Camera, w *scenes.Scene) canvas.Canvas {
	total := c.Width * c.Height
	done := 0
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
				done += 1
				fmt.Printf("\rRendering: %d%%", int(math.Round((float64(done)/float64(total))*100)))
			}(x, y)
		}
	}
	wg.Wait()
	fmt.Printf("\nDone!")
	return img
}

// Computes the color of a pixel.
func colorAt(scene *scenes.Scene, r *ray.Ray) color.Color {
	intersections := intersect(scene, r)
	hit := intersections.Hit()
	if hit.Empty() {
		return color.Black()
	}

	return shadeHit(scene, prepareComputations(hit, r, intersections))
}

// helper method for colorAt.
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
		reflectance := comps.schlick()
		reflected = reflected.Scalar(reflectance)
		refracted = refracted.Scalar(1 - reflectance)
	}
	c = color.Add(c, reflected)
	c = color.Add(c, refracted)

	for i := 0; i < len(scene.Lights); i++ {
		if i == 0 {
			continue
		}
		c = color.Add(c, light.Lighting(
			comps.Shape,
			scene.Lights[i],
			comps.OverPoint,
			comps.EyeV,
			comps.NormalV,
			isShadowed))
	}

	return c
}

// Computes all intersections between a ray and the scene objects.
func intersect(scene *scenes.Scene, r *ray.Ray) shapes.Intersections {
	coll := shapes.Intersections{}
	for i := 0; i < len(scene.Shapes); i++ {
		coll = append(coll, shapes.Intersect(scene.Shapes[i], r)...)
	}
	// Sorting is helpful for reflections and refractions.
	coll.Sort()

	return coll
}

// Determines if a point is in shadow or not.
func shadowAt(scene *scenes.Scene, point tuple.Tuple) bool {
	for i := 0; i < len(scene.Lights); i++ {
		// Measure the distance from point to the light source by subtracting point from the light position
		v := tuple.Subtract(scene.Lights[i].Position(), point)
		// The magnitude of the resulting vector is the distance between the point and the light source.
		dist := v.Magnitude()
		// Create a ray from point toward the light source by normalizing the vector.
		r := ray.New(point, v.Normalize())
		// check for intersections between the point and the light source.
		hit := intersect(scene, r).Hit()

		// if there is an intersection then the point is in shadow.
		if !hit.Empty() && hit.T() < dist {
			return true
		}
	}
	return false
}

// Computes the color of a reflected ray.
func reflectedColor(scene *scenes.Scene, comps Computations) color.Color {
	if comps.Shape.Material().Reflective == 0 || comps.BounceLimit < 1 {
		return color.Black()
	}

	// use OverPoint to avoid shadow acne.
	r := ray.New(comps.OverPoint, comps.ReflectV)
	r.BounceLimit = comps.BounceLimit - 1
	c := colorAt(scene, r)

	return c.Scalar(comps.Shape.Material().Reflective)
}

// Computes the color of a refracted ray.
// Refraction describes how light bends when it passes from one transparent medium to another.
// uses Snell’s Law which describes the relationship between the angles of the light rays and the refractive indices of the two media.
func refractedColor(scene *scenes.Scene, comps Computations) color.Color {
	if comps.Shape.Material().Transparency == 0 || comps.BounceLimit < 1 {
		return color.Black()
	}

	// Find the ratio of first index of refraction to the second.
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
	refractedRay := ray.New(comps.UnderPoint, direction)
	refractedRay.BounceLimit = comps.BounceLimit - 1

	// Find the color of the refracted ray, making sure to multiply by the transparency
	// value to account for any opacity.
	return colorAt(scene, refractedRay).Scalar(comps.Shape.Material().Transparency)
}
