package ray

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/tuple"
)

func TestPosition(t *testing.T) {
	var tests = []struct {
		ray      *Ray
		dist     float64
		expected tuple.Tuple
	}{
		{
			ray:      NewRay(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0)),
			dist:     0,
			expected: tuple.NewPoint(2, 3, 4),
		},
		{
			ray:      NewRay(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0)),
			dist:     1,
			expected: tuple.NewPoint(3, 3, 4),
		},
		{
			ray:      NewRay(tuple.NewPoint(2, 3, 4), tuple.NewVector(0, 1, 0)),
			dist:     -1,
			expected: tuple.NewPoint(2, 2, 4),
		},
		{
			ray:      NewRay(tuple.NewPoint(2, 3, 4), tuple.NewVector(0, 0, 1)),
			dist:     2.5,
			expected: tuple.NewPoint(2, 3, 6.5),
		},
	}

	for _, test := range tests {
		if got := test.ray.Position(test.dist); !got.Equal(test.expected) {
			t.Errorf("ray position:\n%s \n dist: %f. \ngot: \n%s. \nexpected: \n%s", test.ray, test.dist, got, test.expected)
		}
	}
}

func TestTranslate(t *testing.T) {
	r := NewRay(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	expected := NewRay(tuple.NewPoint(4, 6, 8), tuple.NewVector(0, 1, 0))
	x, y, z := 3.0, 4.0, 5.0

	if got := r.Translate(x, y, z); !got.Equal(expected) {
		t.Errorf("translation(%f, %f, %f),\nray:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, r, got, expected)
	}

	x, y, z = 2.0, 3.0, 4.0
	expected = NewRay(tuple.NewPoint(2, 6, 12), tuple.NewVector(0, 3, 0))

	if got := r.Scale(x, y, z); !got.Equal(expected) {
		t.Errorf("scale(%f, %f, %f),\nray:\n%s\n\ngot:\n%s", x, y, z, r, got)
	}
}

func TestScale(t *testing.T) {
	r := NewRay(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	expected := NewRay(tuple.NewPoint(2, 6, 12), tuple.NewVector(0, 3, 0))
	x, y, z := 2.0, 3.0, 4.0

	if got := r.Scale(x, y, z); !got.Equal(expected) {
		t.Errorf("scale(%f, %f, %f),\nray:\n%s\n\ngot:\n%s", x, y, z, r, got)
	}
}
