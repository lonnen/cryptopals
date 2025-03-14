package cryptopals

import (
	"cryptopals/grams"
	"errors"
	"math"
	"slices"
	"sort"
)

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

func transpose(arr [][]byte) [][]byte {
	height := len(arr)
	length := len(arr[0])

	newArr := make([][]byte, length)
	for i := 0; i < length; i++ {
		newArr[i] = make([]byte, height)
		for j := 0; j < height; j++ {
			newArr[i][j] = arr[j][i]
		}
	}
	return newArr
}

type KeyScore struct {
	Key   int
	Value float64
}

type KeyScores []KeyScore

func (k KeyScores) Len() int           { return len(k) }
func (k KeyScores) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k KeyScores) Less(i, j int) bool { return k[i].Value < k[j].Value }

func findKeysize(cipherText []byte, lowerBound int, upperBound int) []KeyScore {
	keyDistances := make(KeyScores, (upperBound - lowerBound))
	for keySize := lowerBound; keySize < upperBound; keySize++ {
		chunks := slices.Collect(slices.Chunk(cipherText, keySize))

		maxChunks := min(len(chunks)-1, 4)
		totalDistance := 0.0
		for i := 0; i < maxChunks; i++ {
			distance, _ := hammingDistance(chunks[i], chunks[i+1])
			totalDistance += float64(distance) / float64(keySize)
		}
		normalizedDistance := totalDistance / float64(maxChunks)
		keyDistances[keySize-lowerBound] = KeyScore{keySize, normalizedDistance}
	}

	sort.Sort(keyDistances)

	return keyDistances[:3]
}

func padToMultipleOf(text []byte, multipleLength int) []byte {
	// if the message is too short, pad it with a 1 followed by 0s until it divides evenly by the key
	padding := make([]byte, multipleLength-(len(text)%multipleLength))
	if len(padding) > 0 {
		padding[0] = 1
		text = append(text, padding...)
	}

	return text
}

func padPKCS7(text []byte, multipleLength int) []byte {
	// if the message is too short, pad it with a 1 followed by 0s until it divides evenly by the key
	padding := make([]byte, multipleLength-(len(text)%multipleLength))
	for i := range padding {
		padding[i] = 0x04
	}

	return text
}
