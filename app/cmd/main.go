package main

/*
Path: app/cmd
File: main.go
Developer(s): Daniel J. Manning
Created: Sun, 05 Oct 2025
*/

import (
	// import packages
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Generate Bytes of a length given by the developer
func GenerateBytes(length int) ([]byte, error) {
	// Make a variable of b & store the amount length
	b := make([]byte, length)
	// If there's an error
	if _, err := rand.Read(b); err != nil {
		// Print the error
		return []byte(""), fmt.Errorf("random generation failed: %w", err)
	}

	// Ensure it's a []byte and encode it in a base64 string ("")
	return []byte(base64.StdEncoding.EncodeToString(b)), nil
}
