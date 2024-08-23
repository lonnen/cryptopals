package cryptopals

import (
	"testing"
)

// https://www.cryptopals.com/sets/1
// This is the qualifying set. We picked the exercises in it to ramp developers up gradually into coding cryptography, but also to verify that we were working with people who were ready to write code.

// This set is relatively easy. With one exception, most of these exercises should take only a couple minutes. But don't beat yourself up if it takes longer than that. It took Alex two weeks to get through the set!

// If you've written any crypto code in the past, you're going to feel like skipping a lot of this. Don't skip them. At least two of them (we won't say which) are important stepping stones to later attacks.

// 1. Convert hex to base64

func Test1ConvertHexToBase64(t *testing.T) {

	const provided string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	computed := Set1Challenge1(provided)

	if computed != expected {
		t.Errorf("Computed %q, Expected %q", computed, expected)
	}
}

func Test2FixedXOR(t *testing.T) {

	const provided string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	computed := Set1Challenge2(provided)

	if computed != expected {
		t.Errorf("Computed %q, Expected %q", computed, expected)
	}
}

func Test3SinglebyteXORcipher(t *testing.T) {

	const provided string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const expected string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	computed := Set1Challenge3(provided)

	if computed != expected {
		t.Errorf("Computed %q, Expected %q", computed, expected)
	}
}
