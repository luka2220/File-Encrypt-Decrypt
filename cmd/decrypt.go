/*
Copyright © 2023 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"program/main/cipher"
	"program/main/error"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt [FILE]",
	Short: "Decrypting a [FILE] with AES encryption",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// Only one number of args
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decrypting")

		// Read encrypted file
		dat, err := os.ReadFile(args[0])
		error.Check(err)

		key := "this_must_be_of_32_byte_length!!"
		decryptedString := cipher.DecryptFile(key, string(dat))

		// Creating a new file to store decrytped text
		filePath := "/Users/luka/Desktop/encrypted/dat2"
		f, err := os.Create(filePath)
		error.Check(err)

		defer f.Close()

		text, err := f.WriteString(decryptedString)
		error.Check(err)
		fmt.Printf("Wrote %d bytes to %s.\n", text, filePath)

		f.Sync()
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Command flags
}
