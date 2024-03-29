package reader

import (
	"testing"

	"github.com/kaizencodes/glimpse/internal/projectpath"
	cfg "github.com/kaizencodes/glimpse/internal/scenes/config"
	"github.com/kaizencodes/glimpse/internal/utils"
)

func TestRead(t *testing.T) {
	config, err := Read(projectpath.Root + `/internal/scenes/reader/examples/test_valid.yml`)
	expectedConfig := cfg.Scene{
		Camera: cfg.Camera{
			Width:  250,
			Height: 125,
			Fov:    1.0471975512,
			From:   []float64{0, 2, -7},
			To:     []float64{0, 1, 0},
			Up:     []float64{0, 1, 0},
		},
		Lights: []cfg.Light{
			{
				Position:  []float64{-10, 10, -10},
				Intensity: []float64{1, 1, 1},
			},
		},
		Objects: []cfg.Object{
			{
				Type: "sphere",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
					{
						Type:   "translate",
						Values: []float64{4.6, 0.4, 1},
					},
					{
						Type:   "rotate-x",
						Values: []float64{1},
					},
					{
						Type:   "rotate-y",
						Values: []float64{2},
					},
					{
						Type:   "rotate-z",
						Values: []float64{3},
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
			},
			{
				Type: "plane",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
				},
				Material: cfg.Material{
					Pattern: cfg.Pattern{
						Type: "stripe",
						Colors: [][]float64{
							{0.8, 0.5, 0.3},
							{0.1, 0.1, 0.1},
						},
						Transform: []cfg.Transform{
							{
								Type:   "scale",
								Values: []float64{0.4, 0.4, 0.4},
							},
						},
					},
					Ambient:         0.1,
					Diffuse:         0.9,
					Specular:        0.9,
					Shininess:       200.0,
					Reflective:      0.0,
					Transparency:    0.0,
					RefractiveIndex: 1.0,
				},
			},
			{
				Type: "cube",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
				},
			},
			{
				Type: "cylinder",
				Transform: []cfg.Transform{
					{
						Type:   "scale",
						Values: []float64{0.4, 0.4, 0.4},
					},
				},
				Minimum: 0,
				Maximum: 1,
				Closed:  true,
			},
		},
	}

	if err != nil {
		t.Errorf("Could not read file: %s", err.Error())
	} else {
		for _, diff := range utils.Compare(config, expectedConfig) {
			t.Errorf("Mismatch: %s", diff)
		}
	}
}

func TestReadMissingFile(t *testing.T) {
	if _, err := Read(`./examples/missing.yml`); err == nil {
		t.Errorf("%s", "No error was raised for invalid config")
	}
}

func TestReadInvalidFile(t *testing.T) {
	if _, err := Read(`./examples/test_invalid.yml`); err == nil {
		t.Errorf("%s", "No error was raised for invalid config")
	}
}
