package grams

import (
	"embed"
)

//go:embed english
var english embed.FS

var files = map[string]string{
	"monograms": "english_monograms.txt",
	"bigrams":   "english_bigrams.txt",
	"trigrams":  "english_trigrams.txt",
	"quadgrams": "english_quadgrams.txt",
}

func chunk(text string, size int) []string {
	text_len := len(text)
	if text_len < size {
		return []string{}
	}

	chunks := make([]string, 0, text_len-size+1)

	for i, j := 0, size; j < text_len; i, j = i+1, j+1 {
		chunks = append(chunks, text[i:j])
	}

	return chunks
}
