package shapes

import (
	"fmt"

	"github.com/kaizencodes/glimpse/internal/materials"
	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
)

type Group struct {
	transform   matrix.Matrix
	material    *materials.Material
	parent      Shape
	children    []Shape
	boundingBox *BoundingBox
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
	if !BoxIntersection(s.boundingBox, r) {
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
		transform:   matrix.DefaultTransform(),
		material:    materials.DefaultMaterial(),
		children:    []Shape{},
		boundingBox: DefaultBoundingBox(),
	}
}

func (g *Group) Parent() Shape {
	return g.parent
}

func (s *Group) SetParent(other Shape) {
	s.parent = other
}

func (g *Group) AddChild(shapes ...Shape) {
	for _, s := range shapes {
		s.SetParent(g)
	}
	g.children = append(g.children, shapes...)
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

func (g *Group) CalculateBoundingBox() {
	//TODO: use i:=0; i<... format over the range everywhere
	for _, child := range g.Children() {
		g.boundingBox.AddBox(child.BoundingBox())
	}
}

func (g *Group) CalculateBoundingBoxCascade() {
	for _, child := range g.Children() {
		child.CalculateBoundingBox()
		g.boundingBox.AddBox(child.BoundingBox())
	}
}

func (g *Group) BoundingBox() *BoundingBox {
	return g.boundingBox
}

func (g *Group) Partition() (left, right []Shape) {
	leftBox, _ := g.boundingBox.Split()

	for _, child := range g.children {
		if leftBox.ContainsBox(child.BoundingBox()) {
			left = append(left, child)
		} else {
			right = append(right, child)
		}
	}
	return left, right
}

func (g *Group) Divide(threshold int) {
	// TODO: recursively check each children and if a group, sum it's children as well.
	// models can have thousands of children and it's not reflected on the parent groups.
	if len(g.children) >= threshold {
		left, right := g.Partition()

		if len(left) > 0 && len(right) > 0 {
			g.children = []Shape{}

			leftGroup := NewGroup()
			leftGroup.AddChild(left...)
			leftGroup.CalculateBoundingBox()
			g.AddChild(leftGroup)

			rightGroup := NewGroup()
			rightGroup.AddChild(right...)
			rightGroup.CalculateBoundingBox()
			g.AddChild(rightGroup)
		}

	}
	for _, child := range g.children {
		if group, ok := child.(*Group); ok {
			group.Divide(threshold)
		}
		if model, ok := child.(*Model); ok {
			model.Divide(threshold)
		}
	}
}
