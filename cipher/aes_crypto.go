package cipher

import (
	"io"
	"log"
	"program/main/error"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

/*
Function for AES encryption on file passed in from command line
key: string - string the must be 32-bit
message: string - text for encrypting
*/
func EncryptFile(key string, message string) string {
	// Key should be 32-bit to select AES-256
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
	cipherText := gcm.Seal(nonce, nonce, plainText, nil)

	return hex.EncodeToString(cipherText)
}

/*
Function decrypting the AES enctypted file
key: string - string the must be 32-bit
message: string - text for decrypting
*/
func DecryptFile(key string, message string) string {
	txt, _ := hex.DecodeString(message)

	cipherBlock, err := aes.NewCipher([]byte(key))

	error.Check(err)

	msgByte := make([]byte, len(txt))
	cipherBlock.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])
	return msg
}
