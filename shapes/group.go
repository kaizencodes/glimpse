package shapes

import (
	"fmt"
	"glimpse/materials"
	"glimpse/matrix"
	"glimpse/tuple"
)

type Group struct {
	transform matrix.Matrix
	material  *materials.Material
	parent    Shape
	children  []Shape
}

func (s *Group) String() string {
	return fmt.Sprintf("Group(material: %s, transform: %s)", s.material, s.transform)
}

func (s *Group) SetTransform(transform matrix.Matrix) {
	s.transform = transform
}

func (s *Group) SetMaterial(mat *materials.Material) {
	s.material = mat
}

func (s *Group) Material() *materials.Material {
	return s.material
}

func (s *Group) Transform() matrix.Matrix {
	return s.transform
}

func (s *Group) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
	return tuple.Tuple{}
}

func NewGroup() *Group {
	return &Group{
		transform: matrix.DefaultTransform(),
		material:  materials.DefaultMaterial(),
		children:  []Shape{},
	}
}

func (g *Group) Parent() Shape {
	return g.parent
}

func (s *Group) SetParent(other Shape) {
	s.parent = other
}

func (g *Group) AddChild(s Shape) {
	s.SetParent(g)
	g.children = append(g.children, s)
}

func (g *Group) Children() []Shape {
	return g.children
}
