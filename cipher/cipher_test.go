package cipher_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yangszwei/pkg/cipher"
	"testing"
)

func TestCipher(t *testing.T) {
	var (
		data = []byte("Hello, world!")
		key  = []byte("1234567890123456")
	)

	ciphertext, err := cipher.Encrypt(data, key)
	require.NoError(t, err)

	plaintext, err := cipher.Decrypt(ciphertext, key)
	require.NoError(t, err)

	assert.Equal(t, data, plaintext)
}
