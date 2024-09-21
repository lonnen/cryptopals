package cryptopals

import "encoding/hex"

func Set1Challenge1(i string) string {
	b64, _ := hexToBase64(i)
	return b64
}

func Set1Challenge2(p string, q string) string {
	z, _ := xorCipherStrings(p, q)
	return z
}

func Set1Challenge3(hexEncoded string) (string, byte, float64) {
	return findSingleByteXOR(hexEncoded)
}

func Set1Challenge4(hexEncodedFile string) (string, byte, float64) {
	return detectSingleByteXOR(hexEncodedFile)
}

func Set1Challenge5(plaintext string, key string) string {
	return hex.EncodeToString(repeatingKeyXOR([]byte(plaintext), []byte(key)))
}

func Set1Challenge6Hamming(a string, b string) int {
	return 0
}

func Set1Challenge6(plaintext string, key string) string {
	return plaintext
}
