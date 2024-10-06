package cryptopals

import (
	"cryptopals/grams"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
)

func hexToBase64(i string) (string, error) {
	bytes, err := hex.DecodeString(i)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
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

func xorCipherStrings(left []byte, right []byte) ([]byte, error) {
	decoded, e := xorCipher(left, right)
	reencoded := make([]byte, hex.EncodedLen(len(decoded)))
	hex.Encode(reencoded, decoded)
	return reencoded, e
}

func xorCipherSingleByte(cipher []byte, key byte) []byte {
	result := make([]byte, len(cipher))
	for i := 0; i < len(cipher); i++ {
		result[i] = cipher[i] ^ key
	}
	return result
}

func findSingleByteXOR(hexed []byte) ([]byte, byte, float64) {
	bestKey := byte(0)
	bestScore := math.Inf(-1)
	bestMessage := make([]byte, len(hexed))
	for b := range 256 {
		key := byte(b)
		decoded := xorCipherSingleByte(hexed, key)
		score := grams.Score(decoded, grams.Bigrams)
		if score > bestScore {
			bestScore = score
			bestKey = key
			bestMessage = decoded
		}
	}

	return bestMessage, bestKey, bestScore
}

func repeatingKeyXOR(plaintext []byte, key []byte) []byte {
	xorText := []byte{}
	for i, b := range plaintext {
		xorText = append(xorText, b^key[i%len(key)])
	}

	return xorText
}

func hammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strings ")
	}

	if len(a) == 0 {
		return 0, nil
	}

	distance := 0
	for i := range a {
		// 1-bits in the XOR'd bytes represent differing bits
		diff := a[i] ^ b[i]

		for j := 0; j < 8; j++ {
			distance += int(diff & 1)
			diff >>= 1
		}
	}

	return distance, nil
}

func transpose(a [][]byte) [][]byte {
	newArr := make([][]byte, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}
