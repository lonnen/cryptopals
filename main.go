/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/base64"
	"encoding/hex"
)

func main() {
}

func hexToBase64(i string) (string, error) {
	bytes, err := hex.DecodeString(i)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
