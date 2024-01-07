/*
Copyright © 2023 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"program/crypto/cipher"
	"program/crypto/error"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt [FILE] --path=[path]",
	Short: "Encrypting a [FILE] with AES encryption",
	Long: `The encrypt command takes in a argumnet for the file to encrypt
	and a flag command for the path to store the result. The file path flag 
	must be the full path on you computer. For example:

	encrypt ~/Desktop/names.txt --path=/Users/luka/Desktop/encrypted/dat2
	decrypt ~/Desktop/decrypted/message.txt`,

	// Only one number of args
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encrypting")

		// Get the flag value
		outputPath := cmd.Flags().Lookup("path").Value.String()

		// dat stores raw bytes of file being read from the command argument args[0]
		dat, err := os.ReadFile(args[0])
		error.Check(err)

		key := "this_must_be_of_32_byte_length!!"
		encryptedString := cipher.EncryptFile(key, string(dat))

		f, err := os.Create(outputPath)
		error.Check(err)

		defer f.Close()

		bytes, err := f.WriteString(encryptedString)
		error.Check(err)
		fmt.Printf("Wrote %d bytes to %s.\n", bytes, outputPath)

		f.Sync()
	},
}

func init() {
	// Required command flags
	encryptCmd.Flags().StringP("path", "p", "", "File path to store the encryption result.")
	_ = encryptCmd.MarkFlagRequired("path")

	rootCmd.AddCommand(encryptCmd)
}
