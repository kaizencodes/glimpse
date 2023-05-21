package reader

import (
	"fmt"
	"os"

	cfg "github.com/kaizencodes/glimpse/internal/scenes/config"
	"github.com/kaizencodes/glimpse/internal/scenes/reader/validator"

	yaml "github.com/goccy/go-yaml"
)

func Read(path string) (cfg.Scene, error) {
	config, err := os.ReadFile(path)
	if err != nil {
		return cfg.Scene{}, err
	}

	err = validator.Validate(config)
	if err != nil {
		return cfg.Scene{}, err
	}

	scene := cfg.Scene{}
	err = yaml.Unmarshal(config, &scene)
	if err != nil {
		panic(fmt.Sprintf("Unmarshaling failed: \n%s", err.Error()))
	}

	return scene, nil
}
