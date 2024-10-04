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

func Set1Challenge6(plaintext string, key string) string {
	cipherText := []byte(plaintext)

	keyDistances := make(map[int]float64)
	for keySize := 2; keySize <= 40; keySize++ {
		chunks := slices.Collect(slices.Chunk(cipherText, keySize))

		maxChunks := max(len(chunks)-1, 2)
		totalDistance := 0.0
		for i := 0; i < maxChunks; i++ {
			distance, _ := hammingDistance(chunks[i], chunks[i+1])
			totalDistance += float64(distance) / float64(keySize)
		}
		normalizedDistance := totalDistance
		keyDistances[keySize] = normalizedDistance
	}
	// the KEYSIZE with the smallest normalized hamming distance is probably the key

	bestLine := ""
	bestKey := ""
	bestCost := math.Inf(-1)

	// break the ciphertext into blocks of KEYSIZE length
	keySize := 3
	chunks := make([][]byte, (len(cipherText) % keySize))
	for i := range slices.Chunk(cipherText, keySize) {
		chunks = append(chunks, i)
	}
	chunkCount := len(chunks)

	println(chunkCount)

	// transpose the blocks (a block of the first byte of each block, then the second byte of each block, etc)
	transposed := transpose(chunks)
	println(len(transposed))

	// solve each block as single-character XOR
	keys := []byte{}
	totalCost := math.Inf(-1)
	for b := range transposed {
		block := transposed[b]
		// for each block, the single-byte XOR key that produces the
		// best letter frequency distribution is the key for that block
		_, key, cost := findSingleByteXOR(block)
		keys = append(keys, key)
		totalCost += cost
	}

	if totalCost > bestCost {
		bestKey = string(keys)
		bestCost = totalCost
		bestLine = string(repeatingKeyXOR(cipherText, keys))
		println("NEW BEST")
		fmt.Printf("key: %s, cost: %f, line: %s", bestKey, bestCost, bestLine)
	}

	fmt.Printf("key: %s", string(bestKey))

	// put the keys together for all the blocks to get the key
	return bestLine
}
