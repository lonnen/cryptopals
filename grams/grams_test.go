package grams

import (
	"testing"
)

func TestChunk(t *testing.T) {
	expected := [4]string{"ABCD", "BCDE", "CDEF", "DEFG"}
	chunks := chunk("ABCDEFG", 4)

	for i := range chunks {
		if chunks[i] != expected[i] {
			t.Errorf("Computed %q, Expected %q", chunks[i], expected[i])
		}
	}
}

func TestEmptyChunk(t *testing.T) {
	chunks := chunk("AB", 3)

	if len(chunks) != 0 {
		t.Errorf("Chunks with slices returned a non-empty array")
	}
}

func TestMonograms(t *testing.T) {
	expected := 26

	monograms := grams(1)
	monogramCount := len(monograms)

	if monogramCount != expected {
		t.Errorf("Expected %d unique monograms, got %d", expected, monogramCount)
	}
}
