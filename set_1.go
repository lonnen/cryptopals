package cryptopals

func Set1Challenge1(i string) string {
	b64, _ := hexToBase64(i)
	return b64
}

func Set1Challenge2(p string, q string) string {
	z, _ := xorCipherStrings(p, q)
	return z
}

func Set1Challenge3(hexEncoded string) (string, byte) {
	return findSingleByteXOR(hexEncoded)
}
