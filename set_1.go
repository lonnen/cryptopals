package cryptopals

func Set1Challenge1(i string) string {
	b64, _ := hexToBase64(i)
	return b64
}

func Set1Challenge2(p string, q string) string {
	decoded_p, _ := hexToBase64(p)
	decoded_q, _ := hexToBase64(q)
	decoded, _ := xor([]byte(decoded_p), []byte(decoded_q))
	return string(decoded)
}

func Set1Challenge3(hexEncoded string) (string, string) {
	return "string", "string"
}
