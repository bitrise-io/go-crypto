package crypto

import (
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
)

// GenerateIV ...
func GenerateIV() ([]byte, error) {
	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		// this could never could/should happen
		return nil, errors.WithStack(err)
	}
	return iv, nil
}
