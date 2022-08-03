package _1_wordvalue

func CalcWordValue(word string) (score int) {
	Scores := LetterScores()
	// Calculate the value of the word entered into function using imported constant mapping LetterScores
	for _, char := range word {
		score += Scores(char)
	}
	return score
}

// Given a list of words, return list of bestWords and the bestScore
func MaxWordsValue(words []string) (bestWords []string, bestScore int) {

	for _, word := range words {
		if currentScore := CalcWordValue(word); currentScore > bestScore {
			bestWords = []string{word}
			bestScore = currentScore
		} else if currentScore == bestScore {
			bestWords = append(bestWords, word)
		}
	}
	return bestWords, bestScore
}
