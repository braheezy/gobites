package _1_wordvalue

import "unicode"

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
