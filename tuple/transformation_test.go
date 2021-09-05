package tuple

import (
    "glimpse/matrix"
    "math"
    "testing"
)

func TestTranslate(t *testing.T) {
    point := Tuple{-3, 4, 5, 1}
    want := Tuple{2, 1, 7, 1}
    var x, y, z float64
    x, y, z = 5, -3, 2
    if got, _ := point.Translate(x, y, z); got.String() != want.String() {
        t.Errorf("translation(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    inv, err := matrix.GetTranslation(x, y, z).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = Tuple{-8, 7, 3, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse translation(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    vector := Tuple{-3, 4, 5, 0}
    if got, _ := vector.Translate(x, y, z); got.String() != vector.String() {
        t.Errorf("vector translation(%f, %f, %f) changed vector,\na:\n%s\n\ngot:\n%s", x, y, z, vector, got)
    }

}

func TestScale(t *testing.T) {
    point := Tuple{-4, 6, 8, 1}
    want := Tuple{-8, 18, 32, 1}
    var x, y, z float64
    x, y, z = 2, 3, 4
    if got, _ := point.Scale(x, y, z); got.String() != want.String() {
        t.Errorf("scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    inv, err := matrix.GetScaling(x, y, z).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = Tuple{-2, 2, 2, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", x, y, z, point, got, want)
    }

    vector := Tuple{-4, 6, 8, 0}
    want = Tuple{-8, 18, 32, 0}
    if got, _ := vector.Scale(x, y, z); got.String() != want.String() {
        t.Errorf("vector scaling(%f, %f, %f),\na:\n%s\n\ngot:\n%s", x, y, z, vector, got)
    }

}

func TestRotateX(t *testing.T) {
    point := Tuple{0, 1, 0, 1}
    r := math.Pi / 2
    want := Tuple{0, 0.00000000000000006123233995736757, 1, 1}

    if got, _ := point.RotateX(r); got.String() != want.String() {
        t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    r = math.Pi / 4
    want = Tuple{0, 0.7071067811865476, 0.7071067811865475, 1}

    if got, _ := point.RotateX(r); got.String() != want.String() {
        t.Errorf("rotatingX(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    inv, err := matrix.GetRotationX(r).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = Tuple{0, 0.7071067811865476, -0.7071067811865475, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse rotatingX(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

}

func TestRotateY(t *testing.T) {
    point := Tuple{0, 0, 1, 1}
    r := math.Pi / 2
    want := Tuple{1, 0, 0.00000000000000006123233995736757, 1}

    if got, _ := point.RotateY(r); got.String() != want.String() {
        t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    r = math.Pi / 4
    want = Tuple{0.7071067811865475, 0, 0.7071067811865476, 1}

    if got, _ := point.RotateY(r); got.String() != want.String() {
        t.Errorf("rotatingY(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    inv, err := matrix.GetRotationY(r).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = Tuple{-0.7071067811865475, 0, 0.7071067811865476, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse rotatingY(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

}

func TestRotateZ(t *testing.T) {
    point := Tuple{0, 1, 0, 1}
    r := math.Pi / 2
    want := Tuple{-1, 0.00000000000000006123233995736757, 0, 1}

    if got, _ := point.RotateZ(r); got.String() != want.String() {
        t.Errorf("rotating(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    r = math.Pi / 4
    want = Tuple{-0.7071067811865475, 0.7071067811865476, 0, 1}

    if got, _ := point.RotateZ(r); got.String() != want.String() {
        t.Errorf("rotatingZ(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

    inv, err := matrix.GetRotationZ(r).Inverse()
    if err != nil {
        t.Error(err)
    }
    want = Tuple{0.7071067811865475, 0.7071067811865476, 0, 1}
    if got, _ := Multiply(inv, point); got.String() != want.String() {
        t.Errorf("inverse rotatingZ(%f),\na:\n%s\n\ngot:\n%s\nexpected: \n%s", r, point, got, want)
    }

}

func TestShear(t *testing.T) {
    var tests = []struct {
        point                  Tuple
        xy, xz, yx, yz, zx, zy float64
        want                   Tuple
    }{
        {
            point: Tuple{2, 3, 4, 1},
            xy:    1, xz: 0, yx: 0, yz: 0, zx: 0, zy: 0,
            want: Tuple{5, 3, 4, 1},
        },
        {
            point: Tuple{2, 3, 4, 1},
            xy:    0, xz: 1, yx: 0, yz: 0, zx: 0, zy: 0,
            want: Tuple{6, 3, 4, 1},
        },
        {
            point: Tuple{2, 3, 4, 1},
            xy:    0, xz: 0, yx: 1, yz: 0, zx: 0, zy: 0,
            want: Tuple{2, 5, 4, 1},
        },
        {
            point: Tuple{2, 3, 4, 1},
            xy:    0, xz: 0, yx: 0, yz: 1, zx: 0, zy: 0,
            want: Tuple{2, 7, 4, 1},
        },
        {
            point: Tuple{2, 3, 4, 1},
            xy:    0, xz: 0, yx: 0, yz: 0, zx: 1, zy: 0,
            want: Tuple{2, 3, 6, 1},
        },
        {
            point: Tuple{2, 3, 4, 1},
            xy:    0, xz: 0, yx: 0, yz: 0, zx: 0, zy: 1,
            want: Tuple{2, 3, 7, 1},
        },
    }

    for _, test := range tests {
        got, _ := test.point.Shear(test.xy, test.xz, test.yx, test.yz, test.zx, test.zy)
        if !got.Equal(test.want) {
            t.Errorf("shearing,\npoint:\n%s\nwith:\n%f, %f, %f, %f, %f, %f, \ngot:\n%s\nexpected: \n%s", test.point, test.xy, test.xz, test.yx, test.yz, test.zx, test.zy, got, test.want)
        }
    }

}
