package crypto

import "github.com/pkg/errors"

// GenerateIV ...
// Deprecated: We don't support this function and this library anymore, please use the `bitrise-io/bits` package at https://github.com/bitrise-io/bits
func GenerateIV() ([]byte, error) {
	secureRandomBytes, err := SecureRandomBytes(12)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return secureRandomBytes, nil
}
