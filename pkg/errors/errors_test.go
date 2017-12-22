package errors_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/nanobox-io/nbx/pkg/errors"
)

var dataEN = `
boxfile-missing: |-
  Failed to find boxfile.yml. Please check your directory
  and project for one and try again.
boxfile-invalid: |-
  The current boxfile.yml contains invalid yaml syntax.
  Please lint it at https://yamllint.com/ and try again.
`
var badYaml = `
terrible: 
- :Failed
`

var parsedDataEN = map[string]string{
	"boxfile-missing": "Failed to find boxfile.yml. Please check your directory\nand project for one and try again.",
	"boxfile-invalid": "The current boxfile.yml contains invalid yaml syntax.\nPlease lint it at https://yamllint.com/ and try again.",
}

// TestParse tests the Parse() functionality.
func TestParse(t *testing.T) {
	r := strings.NewReader(dataEN)

	errs := map[string]string{}
	err := errors.Parse(r, &errs)
	if err != nil {
		t.Errorf("Failed to parse - %s", err.Error())
	}

	if !reflect.DeepEqual(errs, parsedDataEN) {
		t.Errorf("Unexpected results in parsed errors - %s", errs)
	}

	r = strings.NewReader(badYaml)

	err = errors.Parse(r, &errs)
	if err == nil {
		t.Errorf("Failed to parse - %s", err.Error())
	}

}
