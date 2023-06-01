package hash_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {

	sampleTexts := [5]string{
		"exampleplaintex",
	}

	for _, item := range sampleTexts {
		encryptedText := Encrypt([]byte(item))
		assert.NotEqual(t, encryptedText, item)
		assert.Equal(t, Decrypt(encryptedText), item)
	}

}
