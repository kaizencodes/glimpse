package world

import (
    "glimpse/ray"
    "glimpse/tuple"
    "testing"
)

func TestIntersect(t *testing.T) {
    world := Default()
    r := ray.New(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
    sections := world.Intersect(r)
    expected := []float64{4, 4.5, 5.5, 6}
    for i, v := range expected {
        if sections[i].GetT() != v {
            t.Errorf("incorrect t of intersect:\n%s \n \ngot: \n%f. \nexpected: \n%f", r, sections[i].GetT(), v)
        }
    }
}
