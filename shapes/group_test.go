package shapes

import (
	"testing"

	"github.com/kaizencodes/glimpse/matrix"
	"github.com/kaizencodes/glimpse/ray"
	"github.com/kaizencodes/glimpse/tuple"
)

func TestAddChild(t *testing.T) {
	g := NewGroup()
	s := NewTestShape()

	g.AddChild(s)

	if g.Children()[0] != s {
		t.Errorf("group did not add child.")
	}
}

func TestEmptyGroupIntersect(t *testing.T) {
	g := NewGroup()
	r := ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	expected := Intersections{}

	testIntersection(t, g, r, expected)
}

func TestGroupIntersect(t *testing.T) {
	g := NewGroup()
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(matrix.Translation(0, 0, -3))
	s3 := NewSphere()
	s3.SetTransform(matrix.Translation(5, 0, 0))
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	expected := Intersections{
		NewIntersection(1, s2),
		NewIntersection(3, s2),
		NewIntersection(4, s1),
		NewIntersection(6, s1),
	}

	testIntersection(t, g, r, expected)
}

func TestGroupTransformation(t *testing.T) {
	g := NewGroup()
	g.SetTransform(matrix.Scaling(2, 2, 2))
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	g.AddChild(s)
	r := ray.NewRay(tuple.NewPoint(10, 0, -10), tuple.NewVector(0, 0, 1))
	if len(Intersect(g, r)) != 2 {
		t.Errorf("incorrect transformation")
	}
}
