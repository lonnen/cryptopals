package cryptopals

import (
	"encoding/hex"
	"math"
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

	distances := make([]int, 38)

	// KEYSIZE is the guessed length of the key; try 2 to 40
	for keySize := 2; keySize <= 40; keySize++ {
		// for each KEYSIZE, take the first two KEYSIZE of bytes
		// find the hamming distance and normalize it by KEYSIZE

		first := cipherText[0:keySize]
		second := cipherText[keySize+1 : keySize+keySize+1]

		distance, e := hammingDistance(first, second)
		if e != nil {
			continue
		}
		normalizedDistance := float64(distance) / float64(keySize)
		distances[keySize-1] = int(normalizedDistance)
	}
	// the KEYSIZE with the smallest normalized hamming distance is probably the key

	// break the ciphertext into blocks of KEYSIZE length
	// transpose the blocks (a block of the first byte of each block, then the second byte of each block, etc)
	// solve each block as single-character XOR
	// for each block, the single-byte XOR key that produces the best histogram is the key for that block
	// put the keys together for all the blocks to get the key
	return plaintext
}
