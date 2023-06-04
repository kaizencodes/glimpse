package shapes

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
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

	for _, diff := range utils.Compare(box, expected) {
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

	for _, diff := range utils.Compare(box, expected) {
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

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForTestShape(t *testing.T) {
	//  A test shape has a bounding box
	ts := NewTestShape()
	box := BoundFor(ts)
	expected := NewBoundingBox(tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoundForGroups(t *testing.T) {
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
	box := BoundFor(g)
	expected := NewBoundingBox(tuple.NewPoint(-4.5, -3, -5), tuple.NewPoint(4, 7, 4.5))
	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestAddBox(t *testing.T) {
	// Merging two bounding boxes
	box1 := NewBoundingBox(tuple.NewPoint(-5, -2, 0), tuple.NewPoint(7, 4, 4))
	box2 := NewBoundingBox(tuple.NewPoint(8, -7, -2), tuple.NewPoint(14, 2, 8))
	box1.AddBox(box2)
	expected := NewBoundingBox(tuple.NewPoint(-5, -7, -2), tuple.NewPoint(14, 4, 8))

	for _, diff := range utils.Compare(box1, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestContainsPoint(t *testing.T) {
	// Check if a bounding box contains a given point
	box := NewBoundingBox(tuple.NewPoint(5, -2, 0), tuple.NewPoint(11, 4, 7))
	var tests = []struct {
		point    tuple.Tuple
		expected bool
	}{
		{
			point:    tuple.NewPoint(5, -2, 0),
			expected: true,
		}, {
			point:    tuple.NewPoint(11, 4, 7),
			expected: true,
		}, {
			point:    tuple.NewPoint(8, 1, 3),
			expected: true,
		}, {
			point:    tuple.NewPoint(3, 0, 3),
			expected: false,
		}, {
			point:    tuple.NewPoint(8, -4, 3),
			expected: false,
		}, {
			point:    tuple.NewPoint(8, 1, -1),
			expected: false,
		}, {
			point:    tuple.NewPoint(13, 1, 3),
			expected: false,
		}, {
			point:    tuple.NewPoint(8, 5, 3),
			expected: false,
		}, {
			point:    tuple.NewPoint(8, 1, 8),
			expected: false,
		},
	}

	for _, test := range tests {
		if result := box.ContainsPoint(test.point); result != test.expected {
			if test.expected {
				t.Errorf("Expected bounding box to contain point %v", test.point)
			} else {
				t.Errorf("Expected bounding box to not contain point %v", test.point)
			}
		}
	}
}

func TestContainsBox(t *testing.T) {
	// Check if a bounding box contains a given box
	box := NewBoundingBox(tuple.NewPoint(5, -2, 0), tuple.NewPoint(11, 4, 7))
	var tests = []struct {
		box      *BoundingBox
		expected bool
	}{
		{
			box:      NewBoundingBox(tuple.NewPoint(5, -2, 0), tuple.NewPoint(11, 4, 7)),
			expected: true,
		}, {
			box:      NewBoundingBox(tuple.NewPoint(6, -1, 1), tuple.NewPoint(10, 3, 6)),
			expected: true,
		}, {
			box:      NewBoundingBox(tuple.NewPoint(4, -3, -1), tuple.NewPoint(10, 3, 6)),
			expected: false,
		}, {
			box:      NewBoundingBox(tuple.NewPoint(6, -1, 1), tuple.NewPoint(12, 5, 8)),
			expected: false,
		},
	}

	for _, test := range tests {
		if result := box.ContainsBox(test.box); result != test.expected {
			if test.expected {
				t.Errorf("Expected bounding box to contain box %v", test.box)
			} else {
				t.Errorf("Expected bounding box to not contain box %v", test.box)
			}
		}
	}
}

func TestTransformBoundingBox(t *testing.T) {
	// Transforming a bounding box
	box := NewBoundingBox(tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))
	matrix := matrix.Multiply(matrix.RotationX(math.Pi/4), matrix.RotationY(math.Pi/4))
	result := Transform(box, matrix)
	expected := NewBoundingBox(
		tuple.NewPoint(-1.414213562373095, -1.7071067811865475, -1.7071067811865475),
		tuple.NewPoint(1.414213562373095, 1.7071067811865475, 1.7071067811865475))

	for _, diff := range utils.Compare(result, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestTransformedBoundFor(t *testing.T) {
	// Querying a shape's bounding box in its parent's space
	shape := NewSphere()
	shape.SetTransform(
		matrix.Multiply(matrix.Translation(1, -3, 5), matrix.Scaling(0.5, 2, 4)),
	)
	box := TransformedBoundFor(shape)
	expected := NewBoundingBox(tuple.NewPoint(0.5, -5, 1), tuple.NewPoint(1.5, -1, 9))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}
