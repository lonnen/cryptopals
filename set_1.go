package cryptopals

import (
	"encoding/hex"
	"log"
)

func Set1Challenge1(i string) string {
	b64, _ := hexToBase64(i)
	return b64
}

func Set1Challenge2(p string, q string) string {
	decoded_p, _ := hex.DecodeString(p)
	decoded_q, _ := hex.DecodeString(q)
	decoded, e := xor([]byte(decoded_p), []byte(decoded_q))
	if e != nil {
		log.Printf("Error: %q", e)
	}
	return string(decoded)
}

func Set1Challenge3(hexEncoded string) (string, string) {
	return "string", "string"
}
