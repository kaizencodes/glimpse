package utils

import (
	cfg "glimpse/world/config"
	"testing"
)

func TestNestedStructs(t *testing.T) {
	obj1 := cfg.Scene{
		Objects: []cfg.Object{
			{
				Type: "cylinder",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
				},
				Material: cfg.Material{
					Color:           []float64{0.8, 0.5, 0.3},
					Ambient:         0.1,
					Diffuse:         0.9,
					Specular:        0.9,
					Shininess:       200.0,
					Reflective:      0.0,
					Transparency:    0.0,
					RefractiveIndex: 1.0,
				},
				Minimum: 0,
				Maximum: 1,
				Closed:  true,
			},
		},
	}
	obj2 := obj1

	for _, diff := range Compare(obj1, obj2) {
		t.Errorf("Mismatch: %s", diff)
	}
}

func TestCompareWithEmptySlices(t *testing.T) {
	// the material slice is empty in both objects, should return no diffs
	obj1 := cfg.Scene{
		Objects: []cfg.Object{
			{
				Type: "plane",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
				},
			},
		},
	}

	obj2 := obj1

	for _, diff := range Compare(obj1, obj2) {
		t.Errorf("Mismatch: %s", diff)
	}

}

func TestMismatchingTypes(t *testing.T) {
	obj1 := cfg.Scene{
		Objects: []cfg.Object{
			{
				Type: "plane",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
				},
			},
		},
	}

	obj2 := cfg.Scene{
		Objects: []cfg.Object{
			{
				Type: "sphere",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
				},
			},
		},
	}

	if len(Compare(obj1, obj2)) == 0 {
		t.Errorf("Expected mismatching types")
	}
}
