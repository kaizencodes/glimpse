package validator

import (
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

const schema = `
#Camera: {
	width: number
	height: number
	fov: number
}

camera: #Camera
`

// Validates the config file against the schema using cue.

func Validate(config []byte) (err error) {
	ctx := cuecontext.New()
	s := ctx.CompileString(schema)
	err = yaml.Validate(config, s)

	return err
}
