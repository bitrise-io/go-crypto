package crypto

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/pkg/errors"
)

// AES256GCMCipher ...
func AES256GCMCipher(textToEncrypt string, iv []byte, encryptKey string) ([]byte, error) {
	data := []byte(textToEncrypt)
	return AES256GCMCipherBytesInput(data, iv, []byte(encryptKey))
}

// AES256GCMCipherBytesInput ...
func AES256GCMCipherBytesInput(data []byte, iv []byte, encryptKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ciphertext := aesgcm.Seal(nil, iv, data, nil)

	return ciphertext, nil
}

// AES256GCMDecipher ...
func AES256GCMDecipher(encryptedText, iv []byte, encryptKey string) (string, error) {
	plaintext, err := AES256GCMDecipherByteOutput(encryptedText, iv, encryptKey)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(plaintext), nil
}

// AES256GCMDecipherByteOutput ...
func AES256GCMDecipherByteOutput(encryptedText, iv []byte, encryptKey string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(encryptKey))
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	plaintext, err := aesgcm.Open(nil, iv, encryptedText, nil)
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	return plaintext, nil
}
