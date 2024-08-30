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

func xor(left byte, right byte) byte {
	return left ^ right
}

func xorCipher(left []byte, right []byte) ([]byte, error) {
	if len(left) != len(right) {
		return nil, errors.New("buffers must have the same length")
	}

	result := make([]byte, len(left))
	for i := 0; i < len(left); i++ {
		result[i] = left[i] ^ right[i]
	}

	return result, nil
}

func xorCipherStrings(left string, right string) (string, error) {
	decoded_left, _ := hex.DecodeString(left)
	decoded_right, _ := hex.DecodeString(right)
	decoded, e := xorCipher(decoded_left, decoded_right)
	reencoded := make([]byte, hex.EncodedLen(len(decoded)))
	hex.Encode(reencoded, decoded)
	return string(reencoded), e
}
