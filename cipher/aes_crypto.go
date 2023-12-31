package cipher

import (
	"crypto/aes"
	"encoding/hex"
)

/*
Function to implment AES encryption on file passed in from command line
*/
func EncryptFile(key string, message string) string {

	// Key should be 32-bit to select AES-256
	cipherBlock, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	// len() returns the number of bytes in the message
	msgByte := make([]byte, len(message))
	cipherBlock.Encrypt(msgByte, []byte(message))

	return hex.EncodeToString(msgByte)
}
