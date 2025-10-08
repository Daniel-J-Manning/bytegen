package main

// File: bytegen.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain

import (
	// import packages
	"bytegen/bytegen"
	"fmt"
)

func main() {
	// Create token and error variables for generating 32 random bytes
	token, err := bytegen.GenerateRandomBytes(32)

	// If there's an error, print it
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the token of 32 random bytes
	fmt.Println("Generated token:", token)
}

// MIT License
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: MIT (See LICENSE file in repository)
