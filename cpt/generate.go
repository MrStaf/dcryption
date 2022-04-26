package cpt

import (
	"crypto/rand"
	"io/ioutil"
)

func Generate() []byte {
	// Generate a 32 bytes key
	key := make([]byte, 32)
	// Populate the key with cryptographically secure random bytes
	rand.Read(key)
	// Return the key
	return key
}

func Save(key []byte, filepath string) {
	// Save the key to the file
	ioutil.WriteFile(filepath, key, 0644)
}
