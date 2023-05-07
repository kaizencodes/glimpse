package reader

import (
	cfg "glimpse/world/config"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	config, err := Read(`./examples/test.yml`)
	expectedConfig := cfg.Scene{
		Camera: cfg.Camera{
			Width:  250,
			Height: 125,
			Fov:    1.0471975512,
			From:   []float64{0, 2, -7},
			To:     []float64{0, 1, 0},
			Up:     []float64{0, 1, 0},
		},
		Light: cfg.Light{
			Position:  []float64{-10, 10, -10},
			Intensity: []float64{1, 1, 1},
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
				},
			},
		},
	}

	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("incorrect config. Actual: \n%#v\n Expected: \n%#v", config, expectedConfig)
	}
}

func TestReadInvalidFile(t *testing.T) {
	_, err := Read(`./examples/invalid.yml`)
	if err == nil {
		t.Errorf("%s", "No error was raised for invalid config")
	}
}
