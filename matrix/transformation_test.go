package matrix

import (
    "glimpse/tuple"
    "math"
    "testing"
)

func TestTranslate(t *testing.T) {
    point := tuple.Tuple{-3, 4, 5, 1}
    want := tuple.Tuple{2, 1, 7, 1}
    var x, y, z float64
    x, y, z = 5, -3, 2
    if got, _ := Translate(point, x, y, z); got.String() != want.String() {
        t.Errorf("translation(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    inv, err := translation_matrix(x, y, z).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = tuple.Tuple{-8, 7, 3, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse translation(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    vector := tuple.Tuple{-3, 4, 5, 0}
    if got, _ := Translate(vector, x, y, z); got.String() != vector.String() {
        t.Errorf("vector translation(%f, %f, %f) changed vector,\na:\n%s\n\ngot:\n%s", x, y, z, vector, got)
    }

}

func TestScale(t *testing.T) {
    point := tuple.Tuple{-4, 6, 8, 1}
    want := tuple.Tuple{-8, 18, 32, 1}
    var x, y, z float64
    x, y, z = 2, 3, 4
    if got, _ := Scale(point, x, y, z); got.String() != want.String() {
        t.Errorf("scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    inv, err := scaling_matrix(x, y, z).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = tuple.Tuple{-2, 2, 2, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    vector := tuple.Tuple{-4, 6, 8, 0}
    want = tuple.Tuple{-8, 18, 32, 0}
    if got, _ := Scale(vector, x, y, z); got.String() != want.String() {
        t.Errorf("vector scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s", x, y, z, vector, got)
    }

}

func TestRotateX(t *testing.T) {
    point := tuple.Tuple{0, 1, 0, 1}
    r := math.Pi / 2
    want := tuple.Tuple{0, 0.00000000000000006123233995736757, 1, 1}

    if got, _ := RotateX(point, r); got.String() != want.String() {
        t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    r = math.Pi / 4
    want = tuple.Tuple{0, 0.7071067811865476, 0.7071067811865475, 1}

    if got, _ := RotateX(point, r); got.String() != want.String() {
        t.Errorf("rotatingX(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    inv, err := rotation_x_matrix(r).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = tuple.Tuple{0, 0.7071067811865476, -0.7071067811865475, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse rotatingX(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

}

func TestRotateY(t *testing.T) {
    point := tuple.Tuple{0, 0, 1, 1}
    r := math.Pi / 2
    want := tuple.Tuple{1, 0, 0.00000000000000006123233995736757, 1}

    if got, _ := RotateY(point, r); got.String() != want.String() {
        t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    r = math.Pi / 4
    want = tuple.Tuple{0.7071067811865475, 0, 0.7071067811865476, 1}

    if got, _ := RotateY(point, r); got.String() != want.String() {
        t.Errorf("rotatingY(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    inv, err := rotation_y_matrix(r).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = tuple.Tuple{-0.7071067811865475, 0, 0.7071067811865476, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse rotatingY(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

}

func TestRotateZ(t *testing.T) {
    point := tuple.Tuple{0, 1, 0, 1}
    r := math.Pi / 2
    want := tuple.Tuple{-1, 0.00000000000000006123233995736757, 0, 1}

    if got, _ := RotateZ(point, r); got.String() != want.String() {
        t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    r = math.Pi / 4
    want = tuple.Tuple{-0.7071067811865475, 0.7071067811865476, 0, 1}

    if got, _ := RotateZ(point, r); got.String() != want.String() {
        t.Errorf("rotatingZ(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    inv, err := rotation_z_matrix(r).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = tuple.Tuple{0.7071067811865475, 0.7071067811865476, 0, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse rotatingZ(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

}
