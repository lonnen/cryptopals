package cryptopals

import (
	"encoding/hex"
	"fmt"
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

func Set1Challenge6FindKeysize(plaintext string) int {
	cipherText := []byte(plaintext)

	return findKeysize(cipherText, 2, 40) // 2, 40 are magic numbers provided by prompt
}

func Set1Challenge6(plaintext string) string {
	cipherText := []byte(plaintext)

	bestLine := ""
	bestKey := ""
	bestCost := math.Inf(1)

	keySize := findKeysize(cipherText, 2, 40) // 2, 40 are magic numbers provided by prompt

	// chunks := make([][]byte, (len(cipherText) % keySize))
	// for i := range slices.Chunk(cipherText, keySize) {
	// 	chunks = append(chunks, i)
	// }

	chunks := slices.Collect(slices.Chunk(cipherText, keySize))
	transposed := transpose(chunks)

	// solve each block as single-character XOR
	keys := []byte{}
	totalCost := 0.0
	for b := range transposed {
		block := transposed[b]
		// for each block, the single-byte XOR key that produces the
		// best letter frequency distribution is the key for that block
		_, key, cost := findSingleByteXOR(block)
		keys = append(keys, key)
		totalCost += cost
	}

	if totalCost < bestCost {
		bestKey = string(keys)
		bestLine = string(repeatingKeyXOR(cipherText, keys))
		bestCost = totalCost

		println("NEW BEST")
		fmt.Printf("cost: %f, key: %s, line: %s\n", bestCost, bestKey, bestLine)
	}

	// put the keys together for all the blocks to get the key
	return bestKey
}
