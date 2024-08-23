/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func hexToBase64(i string) (string, error) {
	bytes, err := hex.DecodeString(i)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func xor(left []byte, right []byte) ([]byte, error) {
	if len(left) != len(right) {
		return nil, errors.New("Buffers must have the same length")
	}

	result := make([]byte, len(left))
	for i := 0; i < len(left); i++ {
		result[i] = left[i] ^ right[i]
	}

	return result, nil
}

func Set1Challenge1(i string) string {
	b64, _ := hexToBase64(i)
	return b64
}

func Set1Challenge2(p string, q string) string {
	decoded_p, _ := hexToBase64(p)
	decoded_q, _ := hexToBase64(q)
	decoded, _ := xor([]byte(decoded_p), []byte(decoded_q))
	return string(decoded)
}
