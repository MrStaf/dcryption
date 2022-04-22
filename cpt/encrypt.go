package cpt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func Encrypt(filepath string, key []byte, output string) {
	fmt.Println("Encryption Program v0.01")

	// Check if filepath is valid and if it exists
	if _, err := ioutil.ReadFile(filepath); err != nil {
		fmt.Println(err)
	}

	// CHeck if key is valid
	if len(key) != 32 {
		fmt.Println("Key must be 32 bytes long")
	}

	// Check if output is valid
	if _, err := ioutil.ReadFile(output); err == nil {
		fmt.Println("Output file already exists")
	}

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.

	// Open file from file path
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	// Encrypt file
	ciphertext := gcm.Seal(nonce, nonce, file, nil)

	MkAll(output)

	// Write to output file
	err = ioutil.WriteFile(output, ciphertext, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}
