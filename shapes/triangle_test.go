package shapes

import (
	"glimpse/tuple"
	"testing"
)

func TestNewTriangle(t *testing.T) {
	a := tuple.NewPoint(0, 1, 0)
	b := tuple.NewPoint(-1, 0, 0)
	c := tuple.NewPoint(1, 0, 0)
	e1 := tuple.NewVector(-1, -1, 0)
	e2 := tuple.NewVector(1, -1, 0)
	normal := tuple.NewVector(0, 0, -1)

	triangle := NewTriangle(a, b, c)

	if triangle.A() != a {
		t.Errorf("Triangle A\ngot: \n%s. \nexpected: \n%s", triangle.A(), a)
	}
	if triangle.B() != b {
		t.Errorf("Triangle B\ngot: \n%s. \nexpected: \n%s", triangle.A(), b)
	}
	if triangle.C() != c {
		t.Errorf("Triangle C\ngot: \n%s. \nexpected: \n%s", triangle.A(), c)
	}
	if triangle.E1() != e1 {
		t.Errorf("Triangle E1\ngot: \n%s. \nexpected: \n%s", triangle.E1(), e1)
	}
	if triangle.E2() != e2 {
		t.Errorf("Triangle E2\ngot: \n%s. \nexpected: \n%s", triangle.E2(), e2)
	}
	if triangle.Normal() != normal {
		t.Errorf("Triangle normal\ngot: \n%s. \nexpected: \n%s", triangle.Normal(), normal)
	}
}

func TestLocalNormalAt(t *testing.T) {
	triangle := NewTriangle(
		tuple.NewPoint(0, 1, 0),
		tuple.NewPoint(-1, 0, 0),
		tuple.NewPoint(1, 0, 0),
	)

	var tests = []struct {
		point tuple.Tuple
	}{
		{
			point: tuple.NewPoint(0, 0.5, 0),
		},
		{
			point: tuple.NewPoint(-0.5, 0.75, 0),
		},
		{
			point: tuple.NewPoint(0.5, 0.25, 0),
		},
	}

	for _, test := range tests {
		if result := triangle.LocalNormalAt(test.point); !result.Equal(triangle.normal) {
			t.Errorf("Triangle normal:\n point: %s. \nresult: \n%s. \nexpected: \n%s", test.point, result, triangle.normal)
		}
	}
}
