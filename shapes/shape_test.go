package shapes

import (
    "fmt"
    "glimpse/matrix"
    "glimpse/tuple"
    "math"
    "testing"
)

func TestNormal(t *testing.T) {
    var tests = []struct {
        shape Shape
        point tuple.Tuple
        want  tuple.Tuple
    }{
        {
            shape: NewTestShape(),
            point: tuple.NewPoint(1, 0, 0),
            want:  tuple.NewVector(1, 0, 0),
        },
        {
            shape: NewTestShape(),
            point: tuple.NewPoint(0, 1, 0),
            want:  tuple.NewVector(0, 1, 0),
        },
        {
            shape: NewTestShape(),
            point: tuple.NewPoint(0, 0, 1),
            want:  tuple.NewVector(0, 0, 1),
        },
        {
            shape: NewTestShape(),
            point: tuple.NewPoint(1, 0, 0),
            want:  tuple.NewVector(1, 0, 0),
        },
        {
            shape: NewTestShape(),
            point: tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
            want:  tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
        },
        {
            shape: NewTestShape(),
            point: tuple.NewPoint(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0),
            want:  tuple.NewVector(math.Sqrt(3)/3.0, math.Sqrt(3)/3.0, math.Sqrt(3)/3.0).Normalize(),
        },
    }

    for _, test := range tests {
        if got := NormalAt(test.point, test.shape); !got.Equal(test.want) {
            t.Errorf("shpere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", test.shape, test.point, got, test.want)
        }
    }

    // Computing the normal on a translated shape
    shape := NewTestShape()
    shape.SetTransform(matrix.Translation(0, 1, 0))
    point := tuple.NewPoint(0, 1.70711, -0.70711)
    want := tuple.NewVector(0, 0.7071067811865475, -0.7071067811865476)

    if got := NormalAt(point, shape); !got.Equal(want) {
        t.Errorf("shpere normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", shape, point, got, want)
    }

    // Computing the normal on a transformed shape
    shape = NewTestShape()
    transform, _ := matrix.Multiply(matrix.Scaling(1, 0.5, 1), matrix.RotationZ(math.Pi/5.0))
    shape.SetTransform(transform)
    point = tuple.NewPoint(0, math.Sqrt(2)/2.0, -math.Sqrt(2)/2.0)
    want = tuple.NewVector(0, 0.9701425001453319, -0.24253562503633294)

    if got := NormalAt(point, shape); !got.Equal(want) {
        t.Errorf("shape normal:\n%s \n point: %s. \ngot: \n%s. \nexpected: \n%s", shape, point, got, want)
    }
}

type TestShape struct {
    transform matrix.Matrix
    material  Material
}

func (s *TestShape) String() string {
    return fmt.Sprintf("Shape(material: %s\n, transform: %s)", s.Material(), s.Transform())
}

func (s *TestShape) SetTransform(transform matrix.Matrix) {
    s.transform = transform
}

func (s *TestShape) SetMaterial(mat Material) {
    s.material = mat
}

func (s *TestShape) Material() Material {
    return s.material
}

func (s *TestShape) Transform() matrix.Matrix {
    return s.transform
}

func (s *TestShape) LocalNormalAt(point tuple.Tuple) tuple.Tuple {
    return point.ToVector()
}

func NewTestShape() *TestShape {
    return &TestShape{
        transform: DefaultTransform(),
        material:  DefaultMaterial(),
    }
}
