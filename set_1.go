package cryptopals

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"math"
	"slices"
	"strings"
)

func Set1Challenge1(i string) string {
	decoded, _ := hex.DecodeString(i)
	return base64.StdEncoding.EncodeToString(decoded)
}

func Set1Challenge2(p string, q string) string {
	left, _ := hex.DecodeString(p)
	right, _ := hex.DecodeString(q)

	decoded, _ := xorCipher(left, right)
	reencoded := make([]byte, hex.EncodedLen(len(decoded)))
	hex.Encode(reencoded, decoded)

	return string(reencoded)
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
	decoded, _ := base64.StdEncoding.DecodeString(plaintext)

	//bestLine := ""
	bestKey := ""
	bestScore := math.Inf(-1)

	for _, keySize := range findKeysize(decoded, 2, 40) { // 2, 40 are magic numbers provided by prompt

		cipherText := padToMultipleOf(decoded, keySize.Key)

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

func Set1Challenge7(text, key string) string {
	blockSize := 16

	decoded, _ := base64.StdEncoding.DecodeString(text)

	cipherText := padToMultipleOf(decoded, len(key))

	plainText := make([]byte, len(cipherText))
	cipher, _ := aes.NewCipher([]byte(key))

	srcChunks := slices.Collect(slices.Chunk(cipherText, blockSize))
	dstChunks := slices.Collect(slices.Chunk(plainText, blockSize))

	for i := range srcChunks {
		cipher.Decrypt(dstChunks[i], srcChunks[i])
	}

	return string(plainText)
}

func Set1Challenge8(text string) int {
	blockSize := 16

	lines := strings.Split(strings.TrimSpace(string(text)), "\n")

	bestCount := 0
	bestLine := -1

	for line_number, line := range lines {
		decoded, _ := base64.StdEncoding.DecodeString(line)
		cipherText := padToMultipleOf(decoded, blockSize)
		blocks := slices.Collect(slices.Chunk(cipherText, blockSize))

		// ECB lacks diffusion, so some patterns in the original text are preserved
		// they can be detected by looking for unevenly-distributed non-noisy values

		sameBlockCount := 0
		for i := 0; i < len(blocks); i++ {
			for j := 0; j < len(blocks); j++ {
				if bytes.Equal(blocks[i], blocks[j]) {
					sameBlockCount = sameBlockCount + 1
				}
			}
		}

		if sameBlockCount > bestCount {
			bestCount = sameBlockCount
			bestLine = line_number
		}
	}

	return bestLine
}
