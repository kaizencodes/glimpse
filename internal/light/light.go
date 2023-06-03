// Implementation of the Phong reflection model for lighting.
package light

import (
	"fmt"
	"math"

	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/shapes"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type Light struct {
	position  tuple.Tuple
	intensity color.Color
}

func (l Light) String() string {
	return fmt.Sprintf("Light(position: %f, intensity: %f)", l.position, l.intensity)
}

func (l Light) Position() tuple.Tuple {
	return l.position
}

func (l Light) Intensity() color.Color {
	return l.intensity
}

func NewLight(position tuple.Tuple, intensity color.Color) Light {
	return Light{position, intensity}
}

// The Phong reflection model is a method for approximating the illumination of points on a surface.
// It consists of three components:
//  1. Ambient reflection is background lighting, or light reflected from other objects in the environment.
//     The Phong model treats this as a constant, coloring all points on the surface equally.
//  2. Diffuse reflection is light reflected from a matte surface.
//     It depends only on the angle between the light source and the surface normal.
//  3. Specular reflection is the reflection of the light source itself and results in what is called a specular highlight—the bright spot on a curved surface.
//
// It depends only on the angle between the reflection vector and the eye vector and is controlled by a parameter that we’ll call shininess.
// The higher the shininess, the smaller and tighter the specular highlight.
func Lighting(shape shapes.Shape, light Light, point, eyeV, normalV tuple.Tuple, inShadow bool) color.Color {
	mat := shape.Material()
	coloring := shapes.ColorAt(point, shape)
	// combine the surface color with the light's color/intensity
	effectiveColor := color.HadamardProduct(coloring, light.intensity)
	// find the direction to the light source
	lightV := tuple.Subtract(light.position, point).Normalize()
	// compute the ambient contribution
	ambient := effectiveColor.Scalar(mat.Ambient)
	if inShadow {
		// in shadow, we only need to return the ambient contribution.
		return ambient
	}

	// lightDotNormal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface.
	lightDotNormal := tuple.Dot(lightV, normalV)

	var diffuse, specular color.Color
	if lightDotNormal < 0 {
		diffuse = color.Black()
		specular = color.Black()
	} else {
		// compute the diffuse contribution
		diffuse = effectiveColor.Scalar(mat.Diffuse * lightDotNormal)
		// reflectDotEye represents the cosine of the angle between the
		// reflection vector and the eye vector. A negative number means the
		// light reflects away from the eye.
		reflectV := tuple.Reflect(lightV.Negate(), normalV)
		reflectDotEye := tuple.Dot(reflectV, eyeV)
		if reflectDotEye <= 0 {
			specular = color.Black()
		} else {
			// compute the specular contribution
			factor := math.Pow(reflectDotEye, mat.Shininess)
			specular = light.intensity.Scalar(mat.Specular * factor)
		}
	}

	// Add the three contributions together to get the final shading.
	return color.Add(color.Add(ambient, diffuse), specular)
}
