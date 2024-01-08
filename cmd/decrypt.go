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
	"github.com/spf13/viper"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt [FILE] --path=[path]",
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

		// Get the flag value
		outputPath := cmd.Flags().Lookup("path").Value.String()

		// Read encrypted file
		dat, err := os.ReadFile(args[0])
		error.Check(err)

		key := viper.GetString("SECRET_KEY")
		decryptedString := cipher.DecryptFile(key, string(dat))

		f, err := os.Create(outputPath)
		error.Check(err)

		defer f.Close()

		text, err := f.WriteString(decryptedString)
		error.Check(err)
		fmt.Printf("Wrote %d bytes to %s.\n", text, outputPath)

		f.Sync()
	},
}

func init() {
	// Required command flags
	decryptCmd.Flags().StringP("path", "p", "", "File path to store the decryption result.")
	_ = encryptCmd.MarkFlagRequired("path")

	rootCmd.AddCommand(decryptCmd)
}
