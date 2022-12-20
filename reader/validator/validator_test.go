package validator

import (
	"testing"
)

func TestValidate(t *testing.T) {
	input := `camera:
  width: 250
  height: 125
  fov: 1.0471975512
  from: [0, 2, -7]
  to: [0, 1, 0]
  up: [0, 1, 0]
light:
  position: [0, 0, 0]
  intensity: [1, 1, 1]
shapes:
  - sphere:
      transform:
        - ["scale", 1, 2, 3]
        - ["translate", 1, 2, 3]
  - cube:
      transform:
        - ["scale", 1, 2, 3]
      material:
        ambient: 0.1
        diffuse: 0.9
        specular: 0.9
        shininnes: 200.0
        reflective: 0.0
        transparency: 0.0
        refractiveIndex: 1.0
  - model:
      path: foo/bar.yaml
`
	err := Validate([]byte(input))

	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestValidateInvalid(t *testing.T) {
	input := `camera:
  width: 250
  height: 125
  fov: 1.0471975512
  from: [0, 2, -7]
  to: [0, 1, 0]
  up: [0, 1, 0]
light:
  position: [0, 0, 0]
  intensity: [1, 1, 1]
shapes:
  - foobar:
      transform:
        - ["scale", 1, 2, 3]
        - ["translate", 1, 2, 3]
`
	err := Validate([]byte(input))

	if err == nil {
		t.Errorf("%s", "No error was raised for invalid shape `foobar`")
	}
}
