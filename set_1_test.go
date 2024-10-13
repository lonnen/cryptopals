package cryptopals

import (
	_ "embed"
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
		t.Errorf("\nComputed %q, Expected %q\n", computed, expected)
	}
}

func Test2FixedXOR(t *testing.T) {

	const provided string = "1c0111001f010100061a024b53535009181c"
	const provided_xor string = "686974207468652062756c6c277320657965"
	const expected string = "746865206b696420646f6e277420706c6179"

	computed := Set1Challenge2(provided, provided_xor)

	if computed != expected {
		t.Errorf("\nComputed %q, Expected %q\n", computed, expected)
	}
}

func Test3SinglebyteXORcipher(t *testing.T) {

	const provided string = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	const expected string = "Cooking MC's like a pound of bacon"

	computed, key, score := Set1Challenge3(provided)

	if computed != expected {
		t.Errorf("\nComputed %q, %q, %f\nExpected %q\n", computed, key, score, expected)
	}
}

//go:embed data/set1challenge4.txt
var testFourProvided string

func Test4DetectSinglebyteXOR(t *testing.T) {
	const expected string = "Now that the party is jumping\n"

	computed, key, score := Set1Challenge4(testFourProvided)

	if computed != expected {
		t.Errorf("\nComputed %q, %q, %f\nExpected %q\n", computed, key, score, expected)
	}
}

func Test5DetectSinglebyteXOR(t *testing.T) {
	const provided string = "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	const key string = "ICE"
	const expected string = "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	computed := Set1Challenge5(provided, key)

	if computed != expected {
		t.Errorf("\nComputed %s\nExpected %s\n", computed, expected)
	}
}

func Test6AHamming(t *testing.T) {
	const providedHammingA, providedHammingB string = "this is a test", "wokka wokka!!!"
	const expectedHamming int = 37

	computedHamming := Set1Challenge6Hamming(providedHammingA, providedHammingB)
	if computedHamming != expectedHamming {
		t.Errorf("\nHamming: Computed %v\nExpected %v\n", computedHamming, expectedHamming)
	}
}

func Test6Transpose(t *testing.T) {
	provided := [][]byte{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}}
	expected := [][]byte{{0, 1, 2}, {0, 1, 2}, {0, 1, 2}}

	computed := transpose(provided)

	for i := range provided {
		for j := range provided[i] {
			if computed[i][j] != expected[i][j] {
				t.Errorf("\nComputed %x\nExpected %x\n", computed[i], expected[i])
			}
		}
	}
}

func Test6TransposeWider(t *testing.T) {
	provided := [][]byte{{0, 0, 0, 0}, {1, 1, 1, 1}, {2, 2, 2, 2}}
	expected := [][]byte{{0, 1, 2}, {0, 1, 2}, {0, 1, 2}, {0, 1, 2}}

	computed := transpose(provided)

	for i := range expected {
		for j := range expected[i] {
			if computed[i][j] != expected[i][j] {
				t.Errorf("\nComputed %x\nExpected %x\n", computed[i], expected[i])
			}
		}
	}
}

func Test6TransposeTaller(t *testing.T) {
	provided := [][]byte{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}}
	expected := [][]byte{{0, 1, 2, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 2, 3, 4}}

	computed := transpose(provided)

	for i := range expected {
		for j := range expected[i] {
			if computed[i][j] != expected[i][j] {
				t.Errorf("\nComputed %x\nExpected %x\n", computed[i], expected[i])
			}
		}
	}
}

//go:embed data/set1challenge6.txt
var testSixProvided string

func Test6FindKeysize(t *testing.T) {
	const expected int = 29

	computed := Set1Challenge6FindKeysize(testSixProvided)

	found := false
	for _, k := range computed {
		if k.Key == expected {
			found = true
		}
	}

	if !found {
		t.Errorf("\nComputed %v\nExpected %v\n", computed, expected)
	}
}

func Test6FindRepeatingKeyXOR(t *testing.T) {
	expected := "Terminator X: Bring the noise"

	computed := Set1Challenge6(testSixProvided)

	if computed != expected {
		t.Errorf("\nComputed %v\nExpected %v\n", computed, expected)
	}
}
