package cookie

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"
	"crypto/rand"
	"encoding/base64"
)

func encrypt(key []byte, message string) (result string, err error) {
	plainText := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	result = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func decrypt(key []byte, message string) (result string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(message)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("CipherText block size is too short")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	result = string(cipherText)
	return
}
