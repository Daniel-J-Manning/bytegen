// Package: app/bytegen
// File: bytegen.go
// Author: Daniel J. Manning
//
//
// MIT License
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
//
// License: MIT (See LICENSE file in repository)
// GitHub: https://github.com/djmcodechain

package bytegen

import (
	// import packages
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Generate Random Bytes
// Returns a cryptographically secure random string of a given length, encoded in base64.
func GenerateRandomBytes(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("random generation failed: %w", err)
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
