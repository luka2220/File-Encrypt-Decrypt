/*
Copyright © 2023 Luka Piplica piplicaluka64@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cipher_cli",
	Short: "Encrypt and decrypt secret messages in seconds!!!",
	Long: `This application encrypts and decrypts secret messages with ease. 

Try out the Caesar and Bacon Cipher options to generate secret messages and share with your inner circle

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=caesar --key=54

cipher_cli encrypt "Welcome to the hallowed chambers" --algorithm=bacon`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

/*
- Function to add in flags for command
- Declaring a persistent flag which is available to all sub commands
*/
func init() {
	rootCmd.PersistentFlags().StringP("key", "k", "", "The key to pass for the algorithm")
}
