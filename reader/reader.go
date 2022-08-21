package reader

import (
	"glimpse/reader/validator"
	cfg "glimpse/worldconfig"

	"os"

	"github.com/goccy/go-yaml"
)

func Read(path string) (cfg.WorldConfig, error) {
	rawConfig, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := cfg.WorldConfig{}
	err = validator.Validate(rawConfig)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		panic(err)
	}

	return config, nil
}
