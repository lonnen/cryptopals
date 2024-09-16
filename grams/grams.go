package grams

import (
	"embed"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed english
var english embed.FS

type NGrams int

const (
	Monograms NGrams = 1
	Bigrams          = 2
	Trigrams         = 3
	Quadgrams        = 4
)

var files = map[NGrams]string{
	Monograms: "english/english_monograms.txt",
	Bigrams:   "english/english_bigrams.txt",
	Trigrams:  "english/english_trigrams.txt",
	Quadgrams: "english/english_quadgrams.txt",
}

func Grams(size NGrams) map[string]float64 {
	filename := files[size]
	f, err := english.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var gramCounts = make(map[string]float64)
	var gramCount = int(size)
	var total = 0.0
	// parse file into string: int
	i := 0
	j := 0
	for i < len(f) {
		// count-length collection of letters
		gram := f[i : i+gramCount]
		// advance the pointer past the letters and space
		i = i + gramCount + 1
		// grab the number
		j = i
		for i < len(f) {
			if f[i] == 10 {
				break
			}
			i++
		}
		c := f[j : i-1] // exclude the newline from the number
		i++             // move to the start of the next line

		if err != nil {
			log.Fatal("Error while parsing file")
		}
		g := string(gram)
		gramCounts[g], _ = strconv.ParseFloat(string(c), 64)
		total += gramCounts[g]
	}

	for g, c := range gramCounts {
		gramCounts[g] = c / total
	}
	return gramCounts
}

func Chunk(text string, size int) []string {
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

func Score(buffer []byte, frequency map[string]float64) float64 {
	score := 0.0
	for _, b := range buffer {
		score += math.Max(math.Log(frequency[strings.ToUpper(string(b))]), -100.0)
	}
	return score
}
