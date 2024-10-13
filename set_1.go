package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"math"
	"slices"
	"strings"
)

func Set1Challenge1(i string) string {
	b64, _ := hexToBase64(i)
	return b64
}

func Set1Challenge2(p string, q string) string {
	left, _ := hex.DecodeString(p)
	right, _ := hex.DecodeString(q)
	z, _ := xorCipherStrings(left, right)
	return string(z)
}

func Set1Challenge3(hexEncoded string) (string, byte, float64) {
	d, _ := hex.DecodeString(hexEncoded)
	line, key, score := findSingleByteXOR(d)
	return string(line), key, score
}

func Set1Challenge4(hexEncodedFile string) (string, byte, float64) {
	lines := strings.Split(hexEncodedFile, "\n")

	bestLine := make([]byte, len(lines[0]))
	bestKey := byte(0)
	bestScore := math.Inf(-1)
	for _, line := range lines {
		decoded, _ := hex.DecodeString(line)
		decoded, key, score := findSingleByteXOR(decoded)
		if score > bestScore {
			bestLine = decoded
			bestKey = key
			bestScore = score
		}
	}

	return string(bestLine), bestKey, bestScore
}

func Set1Challenge5(plaintext string, key string) string {
	return hex.EncodeToString(repeatingKeyXOR([]byte(plaintext), []byte(key)))
}

func Set1Challenge6Hamming(a string, b string) int {
	distance, _ := hammingDistance([]byte(a), []byte(b))
	return distance
}

func Set1Challenge6FindKeysize(plaintext string) []KeyScore {
	cipherText, _ := base64.StdEncoding.DecodeString(plaintext)
	return findKeysize(cipherText, 2, 40) // 2, 40 are magic numbers provided by prompt
}

func Set1Challenge6(plaintext string) string {
	cipherText, _ := base64.StdEncoding.DecodeString(plaintext)

	//bestLine := ""
	bestKey := ""
	bestScore := math.Inf(-1)

	for _, keySize := range findKeysize(cipherText, 2, 40) { // 2, 40 are magic numbers provided by prompt

		// if the message is too short, pad it with a 1 followed by 0s until it divides evenly by the key
		padding := make([]byte, keySize.Key-(len(cipherText)%keySize.Key))
		if len(padding) > 0 {
			padding[0] = 1
			cipherText = append(cipherText, padding...)
		}
		chunks := slices.Collect(slices.Chunk(cipherText, keySize.Key))

		transposed := transpose(chunks)

		// solve each block as single-character XOR
		keys := []byte{}
		totalScore := 0.0

		for b := range transposed {
			block := transposed[b]
			// for each block, the single-byte XOR key that produces the
			// best letter frequency distribution is the key for that block
			_, key, score := findSingleByteXOR(block)
			keys = append(keys, key)
			totalScore += score
		}

		if totalScore > bestScore {
			bestKey = string(keys)
			//bestLine = string(repeatingKeyXOR(cipherText, keys))
			bestScore = totalScore
		}
	}

	// put the keys together for all the blocks to get the key
	return bestKey
}
