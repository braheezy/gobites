package _1_wordvalue

import (
	"embed"
	"strings"
	"unicode"
)

var scrabbleScores = map[int]string{
	1:  "E A O I N R T L S U",
	2:  "D G",
	3:  "B C M P",
	4:  "F H V W Y",
	5:  "K",
	8:  "J X",
	10: "Q Z",
}

// A "closure" to globally define and provide access to the score map
func LetterScores() func(rune) int {
	var letterScores = make(map[rune]int)

	for score, letters := range scrabbleScores {
		for _, letter := range letters {
			letterScores[letter] = score
		}
	}
	return func(l rune) int {
		return letterScores[unicode.ToUpper(l)]
	}
}

//go:embed dictionary.txt
var f embed.FS
var DictionaryFile, _ = f.ReadFile("dictionary.txt")

func LoadWords() (words []string, err error) {
	// Load dictionary into a list and return list
	words = strings.Fields(string(DictionaryFile))
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}

	return words, nil
}
