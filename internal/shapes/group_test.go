package shapes

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestAddChild(t *testing.T) {
	g := NewGroup()
	s1 := NewTestShape()
	s2 := NewTestShape()
	g.AddChild(s1, s2)

	if g.Children()[0] != s1 || g.Children()[1] != s2 {
		t.Errorf("group did not add child.")
	}
}

func TestEmptyGroupIntersect(t *testing.T) {
	g := NewGroup()
	r := ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
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
	r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
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
	r := ray.New(tuple.NewPoint(10, 0, -10), tuple.NewVector(0, 0, 1))
	if len(Intersect(g, r)) != 2 {
		t.Errorf("incorrect transformation")
	}
}

func TestRemoveChild(t *testing.T) {
	g := NewGroup()
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(matrix.Translation(0, 0, -3))
	g.AddChild(s1)
	g.AddChild(s2)
	g.RemoveChild(s1)

	if len(g.Children()) != 1 && g.Children()[0] != s2 {
		t.Errorf("group did not remove child")
	}
}

func TestPartition(t *testing.T) {
	g := NewGroup()
	s1 := NewSphere()
	s1.SetTransform(matrix.Translation(-2, 0, 0))
	s2 := NewSphere()
	s2.SetTransform(matrix.Translation(2, 0, 0))
	s3 := NewSphere()
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)
	g.CalculateBoundingBoxCascade()
	left, right := g.Partition()

	if left[0] != s1 {
		t.Errorf("incorrect partition for left, expected \n%v\n, got \n%v\n", s1, left[0])
	}
	if right[0] != s2 {
		t.Errorf("incorrect partition for right, expected \n%v\n, got \n%v\n", s2, right[0])
	}
	if right[1] != s3 {
		t.Errorf("incorrect partition for right, expected \n%v\n, got \n%v", s3, right[1])
	}
}

func TestDivide(t *testing.T) {
	// dividing a group partitions its children
	g := NewGroup()
	s1 := NewSphere()
	s1.SetTransform(matrix.Translation(-2, -2, 0))
	s2 := NewSphere()
	s2.SetTransform(matrix.Translation(-2, 2, 0))
	s3 := NewSphere()
	s3.SetTransform(matrix.Scaling(4, 4, 4))
	g.AddChild(s1, s2, s3)
	g.CalculateBoundingBoxCascade()
	g.Divide(1)

	if len(g.Children()) != 2 {
		t.Errorf("incorrect number of children, expected 2, got %v", len(g.Children()))
	}

	subGroup := g.Children()[0].(*Group)
	if len(subGroup.Children()) != 2 {
		t.Errorf("incorrect number of children, expected 2, got %v", len(subGroup.Children()))
	}
	subGroupOfS1 := subGroup.Children()[0].(*Group)
	if len(subGroupOfS1.Children()) != 1 && subGroupOfS1.Children()[0] != s1 {
		t.Errorf("incorrect children, expected %v, got %v", s1, subGroupOfS1.Children())
	}
	subGroupOfS2 := subGroup.Children()[1].(*Group)
	if len(subGroupOfS2.Children()) != 1 && subGroupOfS2.Children()[0] != s2 {
		t.Errorf("incorrect children, expected %v, got %v", s2, subGroupOfS2.Children())
	}

	subGroup2 := g.Children()[1].(*Group)
	if len(subGroup2.Children()) != 1 {
		t.Errorf("incorrect number of children, expected 2, got %v", len(subGroup2.Children()))
	}
	if len(subGroup2.Children()) != 1 && subGroup2.Children()[0] != s3 {
		t.Errorf("incorrect children, expected %v, got %v", s3, subGroup2.Children())
	}

}

func TestBoundingBoxForGroups(t *testing.T) {
	// A group has a bounding box that contains its children
	s1 := NewSphere()
	s1.SetTransform(
		matrix.Multiply(
			matrix.Translation(2, 5, -3),
			matrix.Scaling(2, 2, 2),
		),
	)
	c := NewCylinder()
	c.Minimum = -2
	c.Maximum = 2
	c.SetTransform(
		matrix.Multiply(
			matrix.Translation(-4, -1, 4),
			matrix.Scaling(0.5, 1, 0.5),
		),
	)
	g := NewGroup()
	g.AddChild(s1)
	g.AddChild(c)
	g.CalculateBoundingBoxCascade()
	box := g.BoundingBox()
	expected := NewBoundingBox(tuple.NewPoint(-4.5, -3, -5), tuple.NewPoint(4, 7, 4.5))
	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}
