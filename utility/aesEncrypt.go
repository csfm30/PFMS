package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func AESEncrypt(key string, input string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("Provided key must be non-empty")
	}

	isString := true
	if len(input) == 0 {
		return "", fmt.Errorf("Provided input must be a non-empty string or buffer")
	}

	// Create SHA-256 hash of key
	hash := sha256.New()
	hash.Write([]byte(key))
	keyHash := hash.Sum(nil)

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return "", err
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt
	stream := cipher.NewCTR(block, iv)
	buffer := []byte(input)
	ciphertext := make([]byte, len(buffer))
	stream.XORKeyStream(ciphertext, buffer)

	// Combine IV and ciphertext
	result := append(iv, ciphertext...)

	// Encode to base64 if the input was a string
	var encrypted string
	if isString {
		encrypted = base64.StdEncoding.EncodeToString(result)
	}

	return encrypted, nil
}

func AESDecrypt(key string, encrypted string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("Provided key must be non-empty")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < 17 {
		return "", fmt.Errorf("Encrypted data must decrypt to a non-empty string or buffer")
	}

	// Create SHA-256 hash of key
	hash := sha256.New()
	hash.Write([]byte(key))
	keyHash := hash.Sum(nil)

	// Initialization vector
	iv := ciphertext[:16]
	data := ciphertext[16:]

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return "", err
	}

	// Create AES-CTR mode cipher
	stream := cipher.NewCTR(block, iv)

	// Decrypt
	stream.XORKeyStream(data, data)

	return string(data), nil
}
