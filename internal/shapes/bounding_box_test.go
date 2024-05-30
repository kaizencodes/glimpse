package shapes

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
	"github.com/kaizencodes/glimpse/internal/ray"
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
	TransformBoundingBox(box, matrix)
	expected := NewBoundingBox(
		tuple.NewPoint(-1.414213562373095, -1.7071067811865475, -1.7071067811865475),
		tuple.NewPoint(1.414213562373095, 1.7071067811865475, 1.7071067811865475))

	for _, diff := range utils.Compare(box, expected) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestBoxIntersection(t *testing.T) {
	// Intersecting a ray with a bounding box at the origin
	box := NewBoundingBox(tuple.NewPoint(-1, -1, -1), tuple.NewPoint(1, 1, 1))
	var tests = []struct {
		origin, direction tuple.Tuple
		expected          bool
	}{
		{
			origin:    tuple.NewPoint(5, 0.5, 0),
			direction: tuple.NewVector(-1, 0, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(-5, 0.5, 0),
			direction: tuple.NewVector(1, 0, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(0.5, 5, 0),
			direction: tuple.NewVector(0, -1, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(0.5, -5, 0),
			direction: tuple.NewVector(0, 1, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(0.5, 0, 5),
			direction: tuple.NewVector(0, 0, -1),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(0.5, 0, -5),
			direction: tuple.NewVector(0, 0, 1),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(0, 0.5, 0),
			direction: tuple.NewVector(0, 0, 1),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(-2, 0, 0),
			direction: tuple.NewVector(2, 4, 6),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(0, -2, 0),
			direction: tuple.NewVector(6, 2, 4),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(0, 0, -2),
			direction: tuple.NewVector(4, 6, 2),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(2, 0, 2),
			direction: tuple.NewVector(0, 0, -1),
			expected:  false,
		},
		{
			origin:    tuple.NewPoint(0, 2, 2),
			direction: tuple.NewVector(0, -1, 0),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(2, 2, 0),
			direction: tuple.NewVector(-1, 0, 0),
			expected:  false,
		},
	}
	for _, test := range tests {
		r := ray.New(test.origin, test.direction)
		if result := BoxIntersection(box, r); result != test.expected {
			if test.expected {
				t.Errorf("Expected ray %v to intersect bounding box %v", r, box)
			} else {
				t.Errorf("Expected ray %v to not intersect bounding box %v", r, box)
			}
		}
	}
}

func TestBoxIntersection2(t *testing.T) {
	// Intersecting a ray with a non-cubic bounding box
	box := NewBoundingBox(tuple.NewPoint(5, -2, 0), tuple.NewPoint(11, 4, 7))
	var tests = []struct {
		origin, direction tuple.Tuple
		expected          bool
	}{
		{
			origin:    tuple.NewPoint(15, 1, 2),
			direction: tuple.NewVector(-1, 0, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(-5, -1, 4),
			direction: tuple.NewVector(1, 0, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(7, 6, 5),
			direction: tuple.NewVector(0, -1, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(9, -5, 6),
			direction: tuple.NewVector(0, 1, 0),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(8, 2, 12),
			direction: tuple.NewVector(0, 0, -1),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(6, 0, -5),
			direction: tuple.NewVector(0, 0, 1),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(8, 1, 3.5),
			direction: tuple.NewVector(0, 0, 1),
			expected:  true,
		}, {
			origin:    tuple.NewPoint(9, -1, -8),
			direction: tuple.NewVector(2, 4, 6),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(8, 3, -4),
			direction: tuple.NewVector(6, 2, 4),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(9, -1, -2),
			direction: tuple.NewVector(4, 6, 2),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(4, 0, 9),
			direction: tuple.NewVector(0, 0, 1),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(8, 6, -1),
			direction: tuple.NewVector(0, -1, 0),
			expected:  false,
		}, {
			origin:    tuple.NewPoint(12, 5, 4),
			direction: tuple.NewVector(-1, 0, 0),
			expected:  false,
		},
	}

	for _, test := range tests {
		r := ray.New(test.origin, test.direction)
		if result := BoxIntersection(box, r); result != test.expected {
			if test.expected {
				t.Errorf("Expected ray %v to intersect bounding box %v", r, box)
			} else {
				t.Errorf("Expected ray %v to not intersect bounding box %v", r, box)
			}
		}
	}
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		box, left, right *BoundingBox
	}{
		{
			// Splitting a perfect cube
			box:   NewBoundingBox(tuple.NewPoint(-1, -4, -5), tuple.NewPoint(9, 6, 5)),
			left:  NewBoundingBox(tuple.NewPoint(-1, -4, -5), tuple.NewPoint(4, 6, 5)),
			right: NewBoundingBox(tuple.NewPoint(4, -4, -5), tuple.NewPoint(9, 6, 5)),
		},
		{
			// Splitting an x-wide box
			box:   NewBoundingBox(tuple.NewPoint(-1, -2, -3), tuple.NewPoint(9, 5.5, 3)),
			left:  NewBoundingBox(tuple.NewPoint(-1, -2, -3), tuple.NewPoint(4, 5.5, 3)),
			right: NewBoundingBox(tuple.NewPoint(4, -2, -3), tuple.NewPoint(9, 5.5, 3)),
		},
		{
			// Splitting a y-wide box
			box:   NewBoundingBox(tuple.NewPoint(-1, -2, -3), tuple.NewPoint(5, 8, 3)),
			left:  NewBoundingBox(tuple.NewPoint(-1, -2, -3), tuple.NewPoint(5, 3, 3)),
			right: NewBoundingBox(tuple.NewPoint(-1, 3, -3), tuple.NewPoint(5, 8, 3)),
		},
		{
			// Splitting a z-wide box
			box:   NewBoundingBox(tuple.NewPoint(-1, -2, -3), tuple.NewPoint(5, 3, 7)),
			left:  NewBoundingBox(tuple.NewPoint(-1, -2, -3), tuple.NewPoint(5, 3, 2)),
			right: NewBoundingBox(tuple.NewPoint(-1, -2, 2), tuple.NewPoint(5, 3, 7)),
		},
	}

	for _, test := range tests {
		left, right := test.box.Split()
		for _, diff := range utils.Compare(left, test.left) {
			t.Errorf("Split left Mismatch: %s", diff)
		}
		for _, diff := range utils.Compare(right, test.right) {
			t.Errorf("Split right Mismatch: %s", diff)
		}

	}
}
