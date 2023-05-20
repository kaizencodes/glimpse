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
objects:
  - type: sphere
    transform:
      - type: "scale"
        values: [1, 2, 3]
      - type: "translate"
        values: [1, 2, 3]
      - type: "rotate-x"
        values: [1]
      - type: "rotate-y"
        values: [2]
      - type: "rotate-z"
        values: [3]
  - type: cube
    transform:
      - type: "scale"
        values: [1, 2, 3]
    material:
      color: [1, 0, 0]
      ambient: 0.1
      diffuse: 0.9
      specular: 0.9
      shininess: 200.0
      reflective: 0.0
      transparency: 0.0
      refractive_index: 1.0
  - type: plane
    transform:
      - type: "scale"
        values: [1, 2, 3]
    material:
      pattern:
        type: "stripe"
        colors:
          - [1, 0, 0]
          - [1, 0, 0]
        transform:
          - type: "scale"
            values: [1, 2, 3]
      ambient: 0.1
      diffuse: 0.9
      specular: 0.9
      shininess: 200.0
      reflective: 0.0
      transparency: 0.0
      refractive_index: 1.0
  - type: cylinder
    minimum: 1
    maximum: 2
    closed: true
    transform:
      - type: "scale"
        values: [1, 2, 3]
  - type: cube
  - type: model
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
objects:
  - type: "foobar"
    transform:
      - type: "scale"
        values: [1, 2, 3]
      - type: "translate"
        values: [1, 2, 3]
`
	err := Validate([]byte(input))

	if err == nil {
		t.Errorf("%s", "No error was raised for invalid shape `foobar`")
	}
}