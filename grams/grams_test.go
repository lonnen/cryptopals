package grams

import (
	"testing"
)

func TestChunk(t *testing.T) {
	expected := [4]string{"ABCD", "BCDE", "CDEF", "DEFG"}
	chunks := Chunk("ABCDEFG", 4)

	for i := range chunks {
		if chunks[i] != expected[i] {
			t.Errorf("Computed %q, Expected %q", chunks[i], expected[i])
		}
	}
}

func TestEmptyChunk(t *testing.T) {
	chunks := Chunk("AB", 3)

	if len(chunks) != 0 {
		t.Errorf("Chunks with slices returned a non-empty array")
	}
}

func TestMonograms(t *testing.T) {
	expected := 26

	monograms := Grams(1)
	monogramCount := len(monograms)

	if monogramCount != expected {
		t.Errorf("Expected %d unique monograms, got %d", expected, monogramCount)
	}
}
