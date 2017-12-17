package errors

import (
	"fmt"
	"io"

	"github.com/go-yaml/yaml"
)

// Parse parses a reader and assigns the yaml values to the interface passed in.
func Parse(r io.Reader, i interface{}) error {
	var body []byte
	b := make([]byte, 1)

	for {
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		body = append(body, b...)
	}

	err := yaml.Unmarshal(body, i)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal - %s", err.Error())
	}

	return nil
}
