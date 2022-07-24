package main

import (
	"os"
	"strings"
	"unicode"
)

const DictionaryFile = "dictionary.txt"

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

func LoadWords() (words []string, err error) {
	// Load dictionary into a list and return list
	file, err := os.ReadFile(DictionaryFile)
	if err != nil {
		panic(err)
	}

	words = strings.Fields(string(file))

	return words, nil
}

func CalcWordValue(word string) (score int) {
	Scores := LetterScores()
	// Calculate the value of the word entered into function using imported constant mapping LetterScores
	for _, char := range word {
		score += Scores(char)
	}
	return score
}

func MaxWordValue(words []string) (bestWord string) {
	// Calculate the word with the max value, receive a list of words as arg
	maxScore := 0
	for _, word := range words {
		if currScore := CalcWordValue(word); currScore > maxScore {
			bestWord = word
			maxScore = currScore
		}
	}
	return bestWord
}

func main() {
	// run tests to verify: go test -v
}
