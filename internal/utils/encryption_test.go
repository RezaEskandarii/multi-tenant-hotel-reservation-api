package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {

	sampleTexts := [5]string{
		"Reza", "this is sample text", "hello", "reservation api", "Golang",
	}

	for _, item := range sampleTexts {
		encryptedText := Encrypt(item)
		assert.NotEqual(t, encryptedText, item)
		assert.Equal(t, Decrypt(encryptedText), item)
	}

}
