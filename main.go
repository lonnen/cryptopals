package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
	"strings"
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

func xorCipherSingleByte(cipher []byte, key byte) []byte {
	result := make([]byte, len(cipher))
	for i := 0; i < len(cipher); i++ {
		result[i] = cipher[i] ^ key
	}
	return result
}

// source: https://github.com/piersy/ascii-char-frequency-english/blob/main/ascii_freq.txt
var englishLetterFrequency = map[string]float64{
	"a": 11.7,
	"b": 4.4,
	"c": 5.2,
	"d": 3.2,
	"e": 2.8,
	"f": 4,
	"g": 1.6,
	"h": 4.2,
	"i": 7.3,
	"j": 0.51,
	"k": 0.86,
	"l": 2.4,
	"m": 3.8,
	"n": 2.3,
	"o": 7.6,
	"p": 4.3,
	"q": 0.22,
	"r": 2.8,
	"s": 6.7,
	"t": 16,
	"u": 1.2,
	"v": 0.82,
	"w": 5.5,
	"x": 0.045,
	"y": 0.76,
	"z": 0.045,
}

func isItEnglish(buffer []byte) float64 {
	score := 0.0
	for _, b := range buffer {
		score += math.Max(math.Log(englishLetterFrequency[strings.ToLower(string(b))]), -100.0)
	}
	return score
}

func findSingleByteXOR(hexed string) (string, byte) {
	decoded_hexed, _ := hex.DecodeString(hexed)

	bestKey := byte(0)
	bestScore := math.Inf(-1)
	bestMessage := ""
	for b := range 256 {
		key := byte(b)
		decoded := xorCipherSingleByte(decoded_hexed, key)
		score := isItEnglish(decoded)

		if score > bestScore {
			bestScore = score
			bestKey = key
			bestMessage = string(decoded)
		}
	}

	return bestMessage, bestKey
}
