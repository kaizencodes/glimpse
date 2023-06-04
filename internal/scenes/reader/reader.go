package reader

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kaizencodes/glimpse/internal/projectpath"
	cfg "github.com/kaizencodes/glimpse/internal/scenes/config"

	yaml "github.com/goccy/go-yaml"
)

// relative path to the project root.
const schemaPath = "/configs/schema.cue"

func Read(path string) (cfg.Scene, error) {
	config, err := os.ReadFile(path)
	if err != nil {
		return cfg.Scene{}, err
	}

	// Validate the config file against the schema using cue.
	cmd := exec.Command("cue", "vet", path, projectpath.Root+schemaPath)
	if output, err := cmd.CombinedOutput(); err != nil {
		return cfg.Scene{}, ValidationError{message: string(output)}
	}

	scene := cfg.Scene{}
	err = yaml.Unmarshal(config, &scene)
	if err != nil {
		panic(fmt.Sprintf("Unmarshaling failed: \n%s", err.Error()))
	}

	return scene, nil
}

// CustomError represents a custom error type.
type ValidationError struct {
	message string
}

// Error returns the error message for the CustomError.
func (e ValidationError) Error() string {
	return e.message
}
