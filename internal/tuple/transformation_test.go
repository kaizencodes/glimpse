package tuple

import (
	"math"
	"testing"

	"github.com/kaizencodes/glimpse/internal/matrix"
)

func TestTranslate(t *testing.T) {
	point := Tuple{-3, 4, 5, 1}
	expected := Tuple{2, 1, 7, 1}
	var x, y, z float64
	x, y, z = 5, -3, 2
	if got := point.Translate(x, y, z); got.String() != expected.String() {
		t.Errorf("translation(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, expected)
	}

	inv, err := matrix.Translation(x, y, z).Inverse()
	if err != nil {
		t.Error(err)
	}
	expected = Tuple{-8, 7, 3, 1}
	if got, _ := Multiply(inv, point); got.String() != expected.String() {
		t.Errorf("inverse translation(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, expected)
	}

	vector := Tuple{-3, 4, 5, 0}
	if got := vector.Translate(x, y, z); got.String() != vector.String() {
		t.Errorf("vector translation(%f, %f, %f) changed vector,\na:\n%s\n\ngot:\n%s", x, y, z, vector, got)
	}
}

func TestScale(t *testing.T) {
	point := Tuple{-4, 6, 8, 1}
	expected := Tuple{-8, 18, 32, 1}
	var x, y, z float64
	x, y, z = 2, 3, 4
	if got := point.Scale(x, y, z); got.String() != expected.String() {
		t.Errorf("scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, expected)
	}

	inv, err := matrix.Scaling(x, y, z).Inverse()
	if err != nil {
		t.Error(err)
	}
	expected = Tuple{-2, 2, 2, 1}
	if got, _ := Multiply(inv, point); got.String() != expected.String() {
		t.Errorf("inverse scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, expected)
	}

	vector := Tuple{-4, 6, 8, 0}
	expected = Tuple{-8, 18, 32, 0}
	if got := vector.Scale(x, y, z); got.String() != expected.String() {
		t.Errorf("vector scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s", x, y, z, vector, got)
	}

}

func TestRotateX(t *testing.T) {
	point := Tuple{0, 1, 0, 1}
	r := math.Pi / 2
	expected := Tuple{0, 0.00000000000000006123233995736757, 1, 1}

	if got := point.RotateX(r); got.String() != expected.String() {
		t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

	r = math.Pi / 4
	expected = Tuple{0, 0.7071067811865476, 0.7071067811865475, 1}

	if got := point.RotateX(r); got.String() != expected.String() {
		t.Errorf("rotatingX(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

	inv, err := matrix.RotationX(r).Inverse()
	if err != nil {
		t.Error(err)
	}
	expected = Tuple{0, 0.7071067811865476, -0.7071067811865475, 1}
	if got, _ := Multiply(inv, point); got.String() != expected.String() {
		t.Errorf("inverse rotatingX(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

}

func TestRotateY(t *testing.T) {
	point := Tuple{0, 0, 1, 1}
	r := math.Pi / 2
	expected := Tuple{1, 0, 0.00000000000000006123233995736757, 1}

	if got := point.RotateY(r); got.String() != expected.String() {
		t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

	r = math.Pi / 4
	expected = Tuple{0.7071067811865475, 0, 0.7071067811865476, 1}

	if got := point.RotateY(r); got.String() != expected.String() {
		t.Errorf("rotatingY(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

	inv, err := matrix.RotationY(r).Inverse()
	if err != nil {
		t.Error(err)
	}
	expected = Tuple{-0.7071067811865475, 0, 0.7071067811865476, 1}
	if got, _ := Multiply(inv, point); got.String() != expected.String() {
		t.Errorf("inverse rotatingY(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

}

func TestRotateZ(t *testing.T) {
	point := Tuple{0, 1, 0, 1}
	r := math.Pi / 2
	expected := Tuple{-1, 0.00000000000000006123233995736757, 0, 1}

	if got := point.RotateZ(r); got.String() != expected.String() {
		t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

	r = math.Pi / 4
	expected = Tuple{-0.7071067811865475, 0.7071067811865476, 0, 1}

	if got := point.RotateZ(r); got.String() != expected.String() {
		t.Errorf("rotatingZ(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

	inv, err := matrix.RotationZ(r).Inverse()
	if err != nil {
		t.Error(err)
	}
	expected = Tuple{0.7071067811865475, 0.7071067811865476, 0, 1}
	if got, _ := Multiply(inv, point); got.String() != expected.String() {
		t.Errorf("inverse rotatingZ(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, expected)
	}

}

func TestShear(t *testing.T) {
	var tests = []struct {
		point                  Tuple
		xy, xz, yx, yz, zx, zy float64
		expected               Tuple
	}{
		{
			point: Tuple{2, 3, 4, 1},
			xy:    1, xz: 0, yx: 0, yz: 0, zx: 0, zy: 0,
			expected: Tuple{5, 3, 4, 1},
		},
		{
			point: Tuple{2, 3, 4, 1},
			xy:    0, xz: 1, yx: 0, yz: 0, zx: 0, zy: 0,
			expected: Tuple{6, 3, 4, 1},
		},
		{
			point: Tuple{2, 3, 4, 1},
			xy:    0, xz: 0, yx: 1, yz: 0, zx: 0, zy: 0,
			expected: Tuple{2, 5, 4, 1},
		},
		{
			point: Tuple{2, 3, 4, 1},
			xy:    0, xz: 0, yx: 0, yz: 1, zx: 0, zy: 0,
			expected: Tuple{2, 7, 4, 1},
		},
		{
			point: Tuple{2, 3, 4, 1},
			xy:    0, xz: 0, yx: 0, yz: 0, zx: 1, zy: 0,
			expected: Tuple{2, 3, 6, 1},
		},
		{
			point: Tuple{2, 3, 4, 1},
			xy:    0, xz: 0, yx: 0, yz: 0, zx: 0, zy: 1,
			expected: Tuple{2, 3, 7, 1},
		},
	}

	for _, test := range tests {
		got := test.point.Shear(test.xy, test.xz, test.yx, test.yz, test.zx, test.zy)
		if !got.Equal(test.expected) {
			t.Errorf("shearing,\npoint:\n%s\nwith:\n%f, %f, %f, %f, %f, %f, \ngot:\n%s\nexpected: \n%s", test.point, test.xy, test.xz, test.yx, test.yz, test.zx, test.zy, got, test.expected)
		}
	}

}

func TestChaining(t *testing.T) {
	point := Tuple{1, 0, 1, 1}
	got := point.RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7)
	expected := Tuple{15, 0, 7, 1}

	if !got.Equal(expected) {
		t.Errorf("chaining failed, \n got:\n%s\nexpected:\n%s", got, expected)
	}
}
