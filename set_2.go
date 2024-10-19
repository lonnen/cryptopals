package cryptopals

func Set2Challenge7(provided []byte, blockSize int) []byte {
	return padPKCS7([]byte(provided), blockSize)
}
