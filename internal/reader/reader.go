package reader

import (
	"os"

	"github.com/kaizencodes/glimpse/internal/reader/validator"
	cfg "github.com/kaizencodes/glimpse/internal/scenes/config"

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
		panic(err)
	}

	return scene, nil
}
