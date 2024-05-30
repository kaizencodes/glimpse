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

func (g *Group) String() string {
	return fmt.Sprintf("Group(material: %s, transform: %s)", g.material, g.transform)
}

func (g *Group) SetTransform(transform matrix.Matrix) {
	g.transform = transform
}

func (g *Group) SetMaterial(mat *materials.Material) {
	g.material = mat
}

func (g *Group) Material() *materials.Material {
	return g.material
}

func (g *Group) Transform() matrix.Matrix {
	return g.transform
}

func (g *Group) localNormalAt(point tuple.Tuple, _hit Intersection) tuple.Tuple {
	panic("localNormalAt called on group. Groups do not have normals")
}

func (g *Group) localIntersect(r *ray.Ray) Intersections {
	if !BoxIntersection(g.boundingBox, r) {
		return Intersections{}
	}

	xs := Intersections{}
	for i := 0; i < len(g.children); i++ {
		xs = append(xs, Intersect(g.children[i], r)...)
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

func (g *Group) SetParent(other Shape) {
	g.parent = other
}

func (g *Group) AddChild(shapes ...Shape) {
	for i := 0; i < len(shapes); i++ {
		shapes[i].SetParent(g)
	}
	g.children = append(g.children, shapes...)
}

func (g *Group) RemoveChild(s Shape) {
	s.SetParent(nil)
	for i := 0; i < len(g.children); i++ {
		if g.children[i] == s {
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
	for i := 0; i < len(g.children); i++ {
		g.boundingBox.AddBox(g.children[i].BoundingBox())
	}
}

func (g *Group) CalculateBoundingBoxCascade() {
	for i := 0; i < len(g.children); i++ {
		g.children[i].CalculateBoundingBox()
		g.boundingBox.AddBox(g.children[i].BoundingBox())
	}
}

func (g *Group) BoundingBox() *BoundingBox {
	return g.boundingBox
}

func (g *Group) Partition() (left, right []Shape) {
	leftBox, _ := g.boundingBox.Split()

	for i := 0; i < len(g.children); i++ {
		if leftBox.ContainsBox(g.children[i].BoundingBox()) {
			left = append(left, g.children[i])
		} else {
			right = append(right, g.children[i])
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
	for i := 0; i < len(g.children); i++ {
		if group, ok := g.children[i].(*Group); ok {
			group.Divide(threshold)
		}
		if model, ok := g.children[i].(*Model); ok {
			model.Divide(threshold)
		}
	}
}
