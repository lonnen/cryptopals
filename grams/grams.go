package grams

import (
	"embed"
	"log"
	"math"
	"strconv"
)

//go:embed english
var english embed.FS

var files = map[string]string{
	"monograms": "english_monograms.txt",
	"bigrams":   "english_bigrams.txt",
	"trigrams":  "english_trigrams.txt",
	"quadgrams": "english_quadgrams.txt",
}

func grams(count string) {
	filename := files[count]
	f, err := english.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var gramCounts = make(map[string]float64)
	var gramCount = 1 // get this from a map or enum?
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
