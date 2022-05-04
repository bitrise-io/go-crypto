package crypto

import "github.com/pkg/errors"

// GenerateIV ...
func GenerateIV() ([]byte, error) {
	secureRandomBytes, err := SecureRandomBytes(12)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return secureRandomBytes, nil
}

// GenerateIV256 generates a 256 bit secure random byte sequence
func GenerateIV256() ([]byte, error) {
	secureRandomBytes, err := SecureRandomBytes(32)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return secureRandomBytes, nil
}
