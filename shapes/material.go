package shapes

import (
	"fmt"
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/patterns"
	"glimpse/tuple"
)

type Material struct {
	pattern                               patterns.Pattern
	ambient, diffuse, specular, shininess float64
}

func (mat *Material) ColorAt(pos tuple.Tuple) color.Color {
	return mat.pattern.ColorAt(pos)
}

func (mat *Material) Ambient() float64 {
	return mat.ambient
}

func (mat *Material) Diffuse() float64 {
	return mat.diffuse
}

func (mat *Material) Specular() float64 {
	return mat.specular
}

func (mat *Material) Shininess() float64 {
	return mat.shininess
}

func (mat *Material) String() string {
	return fmt.Sprintf("Material(pattern: %s\n, ambient: %f, diffuse: %f, specular: %f, shininess: %f,)",
		mat.pattern,
		mat.ambient,
		mat.diffuse,
		mat.specular,
		mat.shininess,
	)
}

func DefaultMaterial() *Material {
	return &Material{
		pattern:   patterns.NewMonoPattern(color.White()),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200.0,
	}
}

func NewMaterial(pattern patterns.Pattern, ambient, diffuse, specular, shininess float64) *Material {
	return &Material{pattern, ambient, diffuse, specular, shininess}
}

func (m *Material) SetTransform(transform matrix.Matrix) {
	m.pattern.SetTransform(transform)
}

func (m *Material) SetPattern(pattern patterns.Pattern) {
	m.pattern = pattern
}

func (s *Material) Pattern() patterns.Pattern {
	return s.pattern
}

func (s *Material) Transform() matrix.Matrix {
	return s.pattern.Transform()
}
