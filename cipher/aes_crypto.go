package cipher

import (
	"encoding/hex"
	"io"
	"log"
	"program/crypto/error"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

/*
Function for AES encryption on file passed in from command line
key: string - string the must be 32-bit
message: string - text for encrypting
*/
func EncryptFile(key string, message string) string {
	// Key should be 32-bytes to select AES-256
	cipherBlock, err := aes.NewCipher([]byte(key))
	error.Check(err)

	// Createing an instance of GCM block cipher
	gcm, err := cipher.NewGCM(cipherBlock)
	error.Check(err)

	// Initializing a byte slice with ASCII representation of
	// the message to encrypt
	plainText := []byte(message)
	nonce := make([]byte, gcm.NonceSize())

	// Reads a slice of bytes from crypto secure rand num generator
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("nonce  err: %v", err.Error())
	}

	// Encypting all text from the file
	cipherTextAscii := gcm.Seal(nonce, nonce, plainText, nil)

	return hex.EncodeToString(cipherTextAscii)
}

/*
Function decrypting the AES enctypted file
key: string - string the must be 32-bit
cipherText: string - encrypted file text
*/
func DecryptFile(key string, cipherTextHex string) string {
	// Convert hex-encoded ciphertext to byte slice
	cipherText, err := hex.DecodeString(cipherTextHex)
	error.Check(err)

	// Key should be 32-byte for AES-256
	block, err := aes.NewCipher([]byte(key))
	error.Check(err)

	// Create a GCM instance
	gcm, err := cipher.NewGCM(block)
	error.Check(err)

	// Extract nonce from the cipherText
	nonceSize := gcm.NonceSize()

	nonce := cipherText[:nonceSize]
	cipherText = cipherText[nonceSize:]

	// Decrypt the ciphertext
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	error.Check(err)

	return string(plainText)
}
