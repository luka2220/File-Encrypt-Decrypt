/*
Copyright © 2023 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"program/main/cipher"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt [FILE]",
	Short: "Encrypting a [FILE] with AES encryption",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// Only one number of args
	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encrypting")

		// Read the file and check for an error
		dat, err := os.ReadFile(args[0])
		check(err)
		//fmt.Print(string(dat))

		key := "this_must_be_of_32_byte_length!!"
		encryptedString := cipher.EncryptFile(key, string(dat))

		// Storing encrypted text in new file
		f, err := os.Create("/Users/luka/Desktop/encrypted/dat2")
		check(err)

		defer f.Close()

		text, err := f.WriteString(encryptedString)
		check(err)
		megabytes := text / 1000000
		fmt.Printf("Wrote %d mb.\n", megabytes)

		f.Sync()

		// Open file for I/O operations
		// f, err := os.Open(args[0])
		// check(err)

		// Close file for I/O
		// f.Close()
	},
}

// Displays error if it exists
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
