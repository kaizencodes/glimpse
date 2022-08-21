package validator

import (
	"testing"
)

func TestValidate(t *testing.T) {
	input := `camera:
  width: 32
  height: 40
  fov: 1.4
`
	err := Validate([]byte(input))

	if err != nil {
		t.Errorf("%e", err)
	}
}
