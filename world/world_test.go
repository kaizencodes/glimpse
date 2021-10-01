package world

import (
    "glimpse/color"
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
    w.SetLight(ray.NewLight(tuple.NewPoint(0, 0.25, 0), color.New(1, 1, 1)))
    r = ray.New(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
    object = w.Objects()[1]
    i = ray.NewIntersection(0.5, object)
    comps = ray.PrepareComputations(i, r)

    result = w.shadeHit(comps)
    expected = color.New(0.9049844720832575, 0.9049844720832575, 0.9049844720832575)
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
