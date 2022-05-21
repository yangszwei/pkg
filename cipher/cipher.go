package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

// Encrypt encrypts the data with the key.
func Encrypt(data, key []byte) (string, error) {
	gcm, err := createGCM(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the ciphertext with the key.
func Decrypt(ciphertext string, key []byte) ([]byte, error) {
	gcm, err := createGCM(key)
	if err != nil {
		return nil, err
	}

	data, _ := hex.DecodeString(ciphertext)
	nonce, text := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, text, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// createGCM creates a new GCM cipher with the key.
func createGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}
