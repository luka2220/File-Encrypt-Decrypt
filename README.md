# File-Encrypt-Decrypt
A command-line tool that encrypts and decrypts files using symmetric encryption algorithms.

***

## Project Setup
Note: Make sure you have GO installed on your system
* Clone the repo on your system
  
* Create a config file in your home directory to store an encryption key (must be a string of 32-byte length)
  - `touch ~/.cipher_cli.yaml`
  - Create the key in the file: `SECRET_KEY: "this_must_be_of_32_byte_length!!"`
  
* In the project directory run the following commands:
  - `go install`
  - `go build`
 
***

## Running the CLI commands
Once you have set up the project on your system, you can access the `crypto` command.

* ### Encryption command:
  - `crypto encrypt ~/Desktop/names.txt --path=/Users/luka/Desktop/names-encrypted.txt`
  - The argument `~/Desktop/names.txt` is the path of the file you want to be encrypted.
  - The path flag `/Users/luka/Desktop/names-encrypted.txt` is where you want to store the encrypted file result.

* ### Decryption command:
  - `crypto decrypt ~/Desktop/names-encrypted.txt --path=/Users/luka/Desktop/names.txt`
  - The argument `~/Desktop/names-encrypted.txt` is the path of the file you want to be decrypted.
  - The path flag `/Users/luka/Desktop/names.txt` is where you want to store the decrypted file result.

***

### Note
The path flag must be the full system path of where you want to store the result of the operation. The file is not required to exist, but the path must be valid.
