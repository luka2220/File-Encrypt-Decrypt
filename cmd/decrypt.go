/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
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
		log.SetFlags(log.LstdFlags | log.Llongfile)
		log.Println(": inputted file working")

		key := "this_must_be_of_32_byte_length!!"
		decryptedString := cipher.DecryptFile(key, string(dat))

		// Creating a new file to store decrytped text
		f, err := os.Create("/Users/luka/Desktop/encrypted/message")
		error.Check(err)

		log.Println("Create file path working")

		defer f.Close()

		text, err := f.WriteString(decryptedString)
		error.Check(err)
		fmt.Printf("Wrote %d bytes.\n", text)

		f.Sync()
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
