// materials define the surface appearance of an object.
package materials

import (
	"fmt"

	"github.com/kaizencodes/glimpse/internal/color"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type Material struct {
	pattern *Pattern
	Ambient, Diffuse, Specular, Shininess, Reflective, Transparency,
	RefractiveIndex float64 // refractivity of the material, here are some refractive indices:
	//  Vacuum: 1
	//  Air: 1.00029
	//  Water: 1.333
	//  Glass: 1.52
	//  Diamond: 2.417
}

func (mat *Material) ColorAt(pos tuple.Tuple) color.Color {
	return mat.pattern.colorAt(pos)
}

func (mat *Material) String() string {
	return fmt.Sprintf("Material(Ambient: %f, Diffuse: %f, Specular: %f, Shininess: %f,)",
		mat.Ambient,
		mat.Diffuse,
		mat.Specular,
		mat.Shininess,
	)
}

func DefaultMaterial() *Material {
	return &Material{
		pattern:         NewPattern(Base, color.White()),
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200.0,
		Reflective:      0.0,
		Transparency:    0.0,
		RefractiveIndex: 1.0,
	}
}

func NewMaterial(c color.Color, ambient, diffuse, specular, shininess, reflective, transparency, refractiveIndex float64) *Material {
	return &Material{
		pattern:         NewPattern(Base, c),
		Ambient:         ambient,
		Diffuse:         diffuse,
		Specular:        specular,
		Shininess:       shininess,
		Reflective:      reflective,
		Transparency:    transparency,
		RefractiveIndex: refractiveIndex,
	}
}

func (m *Material) SetTransform(transform matrix.Matrix) {
	m.pattern.transform = (transform)
}

func (m *Material) SetPattern(pattern *Pattern) {
	m.pattern = pattern
}

func (s *Material) Pattern() *Pattern {
	return s.pattern
}

func (s *Material) Transform() matrix.Matrix {
	return s.pattern.transform
}
