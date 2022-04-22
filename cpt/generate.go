package cpt

import (
	"crypto/rand"
)

func Generate() []byte {
	// Generate a 32 bytes key
	key := make([]byte, 32)
	// Populate the key with cryptographically secure random bytes
	rand.Read(key)
	// Return the key
	return key
}
