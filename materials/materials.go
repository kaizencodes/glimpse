package materials

import (
	"fmt"
	"glimpse/color"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Material struct {
	pattern                                           *Pattern
	ambient, diffuse, specular, shininess, reflective float64
}

func (mat *Material) ColorAt(pos tuple.Tuple) color.Color {
	return mat.pattern.colorAt(pos)
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

func (mat *Material) Reflective() float64 {
	return mat.reflective
}

func (mat *Material) SetAmbient(ambient float64) {
	mat.ambient = ambient
}

func (mat *Material) SetDiffuse(diffuse float64) {
	mat.diffuse = diffuse
}

func (mat *Material) SetSpecular(specular float64) {
	mat.specular = specular
}

func (mat *Material) SetShininess(shininess float64) {
	mat.shininess = shininess
}

func (mat *Material) SetReflective(reflective float64) {
	mat.reflective = reflective
}

func (mat *Material) String() string {
	return fmt.Sprintf("Material(ambient: %f, diffuse: %f, specular: %f, shininess: %f,)",
		mat.ambient,
		mat.diffuse,
		mat.specular,
		mat.shininess,
	)
}

func DefaultMaterial() *Material {
	return &Material{
		pattern:    NewPattern(Base, color.White()),
		ambient:    0.1,
		diffuse:    0.9,
		specular:   0.9,
		shininess:  200.0,
		reflective: 0.0,
	}
}

func NewMaterial(c color.Color, ambient, diffuse, specular, shininess, reflective float64) *Material {
	return &Material{NewPattern(Base, c), ambient, diffuse, specular, shininess, reflective}
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
