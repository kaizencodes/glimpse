package shapes

import (
	"fmt"

	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
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

func (s *Group) localNormalAt(point tuple.Tuple, _hit Intersection) tuple.Tuple {
	panic("localNormalAt called on group. Groups do not have normals")
}

func (s *Group) localIntersect(r *ray.Ray) Intersections {
	// TODO: check if we can save the bound calculation and reuse it
	if !BoxIntersection(BoundFor(s), r) {
		return Intersections{}
	}

	xs := Intersections{}
	for _, child := range s.Children() {
		xs = append(xs, Intersect(child, r)...)
	}
	xs.Sort()
	return xs
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

func (g *Group) RemoveChild(s Shape) {
	s.SetParent(nil)
	for i, child := range g.children {
		if child == s {
			// replace the child with the last element
			g.children[i] = g.children[len(g.children)-1]
			// remove the last element
			g.children = g.children[:len(g.children)-1]
			// this is a much faster solution and the ordering doesn't matter
			return
		}
	}
}

func (g *Group) Children() []Shape {
	return g.children
}

func (g *Group) Partition() (left, right []Shape) {
	box := BoundFor(g)
	leftBox, rightBox := box.Split()
	for _, child := range g.children {
		if leftBox.ContainsBox(TransformedBoundFor(child)) {
			left = append(left, child)
		} else if rightBox.ContainsBox(TransformedBoundFor(child)) {
			right = append(right, child)
		}
	}
	for _, child := range left {
		g.RemoveChild(child)
	}
	for _, child := range right {
		g.RemoveChild(child)
	}
	return left, right
}
