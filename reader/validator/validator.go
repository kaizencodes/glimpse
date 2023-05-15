package validator

import (
	"os"

	"github.com/kaizencodes/glimpse/projectpath"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

// relative path to the project root.
const schemaPath = "/reader/validator/schema.cue"

// Validates the config file against the schema using cue.
func Validate(config []byte) error {
	ctx := cuecontext.New()
	schema, err := os.ReadFile(projectpath.Root + schemaPath)
	if err != nil {
		panic(err)
	}
	s := ctx.CompileBytes(schema)
	err = yaml.Validate(config, s)

	return err
}
