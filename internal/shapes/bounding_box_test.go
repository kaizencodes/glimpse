package shapes

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/tuple"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestAddPoint(t *testing.T) {
	// Adding points to an empty bounding box
	box := DefaultBoundingBox()
	box.AddPoint(tuple.NewPoint(-5, 2, 0))
	box.AddPoint(tuple.NewPoint(7, 2, -3))
	expected := NewBoundingBox(tuple.NewPoint(-5, 2, -3), tuple.NewPoint(7, 2, 0))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForSphere(t *testing.T) {
	//  A sphere has a bounding box
	s := NewSphere()
	box := BoundFor(s)
	expected := NewBoundingBox(tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForPlane(t *testing.T) {
	//  A plane has a bounding box
	p := NewPlane()
	box := BoundFor(p)
	expected := NewBoundingBox(
		tuple.NewPoint(math.Inf(-1), 0, math.Inf(-1)),
		tuple.NewPoint(math.Inf(1), 0, math.Inf(1)),
	)

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForCube(t *testing.T) {
	//  A cube has a bounding box
	c := NewCube()
	box := BoundFor(c)
	expected := NewBoundingBox(tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForCylinder(t *testing.T) {
	//  A cylinder has a bounding box
	c := NewCylinder()
	box := BoundFor(c)
	expected := NewBoundingBox(tuple.NewPoint(-1, math.Inf(-1), -1), tuple.NewPoint(1, math.Inf(1), 1))

	for _, diff := range utils.Compare(&box, &expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForLimitedCylinder(t *testing.T) {
	//  A cylinder has a bounding box
	c := NewCylinder()
	c.Minimum = -5
	c.Maximum = 3
	box := BoundFor(c)
	expected := NewBoundingBox(tuple.NewPoint(-1, -5, -1), tuple.NewPoint(1, 3, 1))

	for _, diff := range utils.Compare(&box, &expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForTriangles(t *testing.T) {
	//  A triangle has a bounding box
	p1 := tuple.NewPoint(-3, 7, 2)
	p2 := tuple.NewPoint(6, 2, -4)
	p3 := tuple.NewPoint(2, -1, -1)
	tri := NewTriangle(p1, p2, p3)
	box := BoundFor(tri)
	expected := NewBoundingBox(tuple.NewPoint(-3, -1, -4), tuple.NewPoint(6, 7, 2))

	for _, diff := range utils.Compare(&box, &expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForTestShape(t *testing.T) {
	//  A test shape has a bounding box
	ts := NewTestShape()
	box := BoundFor(ts)
	expected := NewBoundingBox(tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))

	for _, diff := range utils.Compare(&box, &expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}
