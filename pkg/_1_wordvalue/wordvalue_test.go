package _1_wordvalue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var TestWords = map[string]int{"bob": 7, "julian": 13, "quit": 13, "quits": 14, "quite": 14}

func TestLoadWords(t *testing.T) {
	words, _ := LoadWords()

	expected := 178695
	actual := len(words)

	require.Equal(t, actual, expected)
}

func TestCalcWordValue(t *testing.T) {
	for word, expected_score := range TestWords {
		actual_score := CalcWordValue(word)

		require.Equal(t, actual_score, expected_score)
	}
}

func TestMaxWordsValue(t *testing.T) {
	var words []string
	for word := range TestWords {
		words = append(words, word)
	}

	actualMaxWords, actualMaxScore := MaxWordsValue(words)
	expectedMaxWords := []string{"quits", "quite"}
	expectedMaxScore := 14

	require.Equal(t, actualMaxScore, expectedMaxScore)
	require.ElementsMatch(t, actualMaxWords, expectedMaxWords)

}
