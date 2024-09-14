package grams

import (
	"embed"
	"log"
	"math"
	"strconv"
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
	Monograms: "english_monograms.txt",
	Bigrams:   "english_bigrams.txt",
	Trigrams:  "english_trigrams.txt",
	Quadgrams: "english_quadgrams.txt",
}

func grams(size NGrams) {
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
		i += gramCount + 1
		// grab the number
		j = i
		for f[i] != 10 {
			i++
		}
		c := f[j:i]

		if err != nil {
			log.Fatal("Error while parsing file")
		}
		g := string(gram)
		gramCounts[g], _ = strconv.ParseFloat(string(c), 64)
	}

	for g, c := range gramCounts {
		gramCounts[g] = math.Log(c / total)
	}

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
