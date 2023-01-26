package utils

import (
	"crypto/aes"
	"crypto/cipher"
)

const (
	key = "example key 1234"
)

func Encrypt(text []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	b := text[0 : len(text)/2] // first half of the plaintext is used as IV

	ciphertext := make([]byte, aes.BlockSize+len(text))

	// copy IV to the beginning of ciphertext block
	copy(ciphertext[:aes.BlockSize], b)

	cfb := cipher.NewCFBEncrypter(block, b) // create CFB encrypter using IV and key

	// encrypt the rest of the plaintext
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text[aes.BlockSize:])

	return ciphertext // return encrypted data
}

func Decrypt(text []byte) []byte {

	block, err := aes.NewCipher([]byte(key)) // create new AES cipher using key

	if err != nil { // check for errors
		panic(err) // panic if error occurs
	}

	b := text[:aes.BlockSize] // get IV from beginning of ciphertext block

	plaintext := make([]byte, len(text)) // create empty byte array for decrypted data

	cfb := cipher.NewCFBDecrypter(block, b) // create CFB decrypter using IV and key

	cfb.XORKeyStream(plaintext, text[aes.BlockSize:]) // decrypt the rest of the ciphertext block

	return plaintext // return decrypted data
}
