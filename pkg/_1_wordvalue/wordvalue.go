package _1_wordvalue

import (
	"os"
	"strings"
)

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
