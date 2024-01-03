package cipher

import (
	"program/main/error"

	"crypto/aes"
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

	//gcm, err = cipher.NewGCM()

	msgByte := make([]byte, len(message))
	cipherBlock.Encrypt(msgByte, []byte(message))

	return hex.EncodeToString(msgByte)
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
