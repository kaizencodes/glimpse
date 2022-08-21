package reader

import (
	cfg "glimpse/worldconfig"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	config, err := Read(`./examples/check.yml`)
	expectedConfig := cfg.WorldConfig{
		Camera: cfg.Camera{
			Width:  32,
			Height: 40,
			Fov:    1.4,
		},
	}
	if err != nil {
		t.Errorf("%e", err)
	}
	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("incorrect config. Actual: \n%#v\n Expected: \n%#v", config, expectedConfig)
	}
}
