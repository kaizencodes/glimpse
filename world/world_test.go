package world

import (
    "glimpse/color"
    "glimpse/matrix"
    "glimpse/objects"
    "glimpse/ray"
    "glimpse/tuple"
    "testing"
)

func TestIntersect(t *testing.T) {
    w := Default()
    r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
    sections := w.intersect(r)
    expected := []float64{4, 4.5, 5.5, 6}
    for i, v := range expected {
        if sections[i].T() != v {
            t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", r, sections[i].T(), v)
        }
    }
}

func TestShadeHit(t *testing.T) {
    w := Default()
    r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
    object := w.Objects()[0]
    i := ray.NewIntersection(4, object)
    comps := ray.PrepareComputations(i, r)

    result := w.shadeHit(comps)
    expected := color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
    if result != expected {
        t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
    }

    w = Default()
    w.SetLights([]ray.Light{
        ray.NewLight(tuple.NewPoint(0, 0.25, 0), color.New(1, 1, 1)),
    })
    r = ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
    object = w.Objects()[1]
    i = ray.NewIntersection(0.5, object)
    comps = ray.PrepareComputations(i, r)

    result = w.shadeHit(comps)
    expected = color.New(0.9049844720832575, 0.9049844720832575, 0.9049844720832575)
    if result != expected {
        t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
    }

    w = Default()
    w.SetLights([]ray.Light{
        ray.NewLight(tuple.NewPoint(0, 0.25, 0), color.New(1, 1, 1)),
        ray.NewLight(tuple.NewPoint(1, 0, 1), color.New(0.9, 0.7, 0)),
    })
    r = ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
    object = w.Objects()[1]
    i = ray.NewIntersection(0.5, object)
    comps = ray.PrepareComputations(i, r)

    result = w.shadeHit(comps)
    expected = color.New(0.19, 0.16999999999999998, 0.1)
    if result != expected {
        t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
    }

    w = Default()
    w.SetLights([]ray.Light{
        ray.NewLight(tuple.NewPoint(0, 0, -10), color.New(1, 1, 1)),
    })
    s1 := objects.NewSphere()
    s2 := objects.NewSphere()
    s2.SetTransform(matrix.Translation(0, 0, 10))
    w.SetObjects([]objects.Object{s1, s2})

    r = ray.New(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
    i = ray.NewIntersection(4, s2)
    comps = ray.PrepareComputations(i, r)

    result = w.shadeHit(comps)
    expected = color.New(0.1, 0.1, 0.1)
    if result != expected {
        t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
    }
}

func TestColorAt(t *testing.T) {
    w := Default()
    r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 1, 0))
    result := w.ColorAt(r)
    expected := color.Black()
    if result != expected {
        t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
    }

    w = Default()
    r = ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
    result = w.ColorAt(r)
    expected = color.New(0.38066119308103435, 0.47582649135129296, 0.28549589481077575)
    if result != expected {
        t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
    }

    w = Default()
    outer := w.Objects()[0]
    m := outer.Material()
    outer.SetMaterial(objects.NewMaterial(m.Color(), 1, m.Diffuse(), m.Specular(), m.Shininess()))
    inner := w.Objects()[1]
    m = inner.Material()
    inner.SetMaterial(objects.NewMaterial(m.Color(), 1, m.Diffuse(), m.Specular(), m.Shininess()))

    r = ray.New(tuple.NewPoint(0, 0, 0.75), tuple.NewVector(0, 0, -1))
    result = w.ColorAt(r)
    expected = inner.Material().Color()
    if result != expected {
        t.Errorf("incorrect Shading:\nresult: \n%s. \nexpected: \n%s", result, expected)
    }
}

func TestViewTransformation(t *testing.T) {
    var tests = []struct {
        from     tuple.Tuple
        to       tuple.Tuple
        up       tuple.Tuple
        expected matrix.Matrix
    }{
        {
            from:     tuple.NewPoint(0, 0, 0),
            to:       tuple.NewPoint(0, 0, -1),
            up:       tuple.NewVector(0, 1, 0),
            expected: matrix.NewIdentity(4),
        },
        {
            from:     tuple.NewPoint(0, 0, 0),
            to:       tuple.NewPoint(0, 0, 1),
            up:       tuple.NewVector(0, 1, 0),
            expected: matrix.Scaling(-1, 1, -1),
        },
        {
            from:     tuple.NewPoint(0, 0, 8),
            to:       tuple.NewPoint(0, 0, 0),
            up:       tuple.NewVector(0, 1, 0),
            expected: matrix.Translation(0, 0, -8),
        },
        {
            from: tuple.NewPoint(1, 3, 2),
            to:   tuple.NewPoint(4, -2, 8),
            up:   tuple.NewVector(1, 1, 0),
            expected: matrix.Matrix{
                []float64{-0.5070925528371099, 0.5070925528371099, 0.6761234037828132, -2.366431913239846},
                []float64{0.7677159338596801, 0.6060915267313263, 0.12121830534626524, -2.8284271247461894},
                []float64{-0.35856858280031806, 0.5976143046671968, -0.7171371656006361, 0},
                []float64{0, 0, 0, 1},
            },
        },
    }

    for _, test := range tests {
        result := ViewTransformation(test.from, test.to, test.up)
        if result.String() != test.expected.String() {
            t.Errorf("ViewTransformation,\nto:\n%s\nfrom:\n%s\nup:\n%s\nresult:\n%s\nexpected: \n%s", test.to, test.from, test.up, result, test.expected)
        }
    }
}

func TestShadowAt(t *testing.T) {
    w := Default()
    var tests = []struct {
        w        *World
        point    tuple.Tuple
        expected bool
    }{
        {
            w:        w,
            point:    tuple.NewPoint(0, 10, 0),
            expected: false,
        },
        {
            w:        w,
            point:    tuple.NewPoint(10, -10, 10),
            expected: true,
        },
        {
            w:        w,
            point:    tuple.NewPoint(-20, 20, -20),
            expected: false,
        },
        {
            w:        w,
            point:    tuple.NewPoint(-2, 2, -2),
            expected: false,
        },
    }

    for _, test := range tests {
        result := test.w.shadowAt(test.point)
        if result != test.expected {
            t.Errorf("ShadowAt,\npoint:\n%s\nresult:\n%t\nexpected: \n%t", test.point, result, test.expected)
        }
    }
}
