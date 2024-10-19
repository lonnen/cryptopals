package cryptopals

import (
	"bytes"
	"testing"
)

func Test9PKCS7padding(t *testing.T) {
	provided := []byte("YELLOW SUBMARINE")
	const providedBlockSize = 20
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	computed := Set2Challenge7(provided, providedBlockSize)

	if bytes.Equal(computed, expected) {
		t.Errorf("\nComputed %q, Expected %q\n", computed, expected)
	}
}
