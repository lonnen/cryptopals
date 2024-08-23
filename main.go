/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func hexToBase64(i string) (string, error) {
	bytes, err := hex.DecodeString(i)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func Set1Challenge1(i string) string {
	b64, _ := hexToBase64(i)
	return b64
}
