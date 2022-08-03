package _2_scrabble

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPouch(t *testing.T) {
	letters, pouch_len := Pouch()(5)

	fmt.Println(letters)
	fmt.Println(len(letters))

	require.Equal(t, len(letters), 5)
	require.Equal(t, pouch_len, 98)
	require.Regexp(t, "^[A-Z]{5}$", string(letters))
}

func TestHasVowels(t *testing.T) {
	wordWithVowels := "HELLO"
	wordWithoutVowels := "SHY"

	require.True(t, HasVowels(wordWithVowels))
	require.False(t, HasVowels(wordWithoutVowels))
}

func TestIsValid(t *testing.T) {
	draw := "GARYTEV"

	word := "GARYTEV"
	require.False(t, IsValid(word, draw), "Test: Not dict word")

	word = "F"
	require.False(t, IsValid(word, draw), "Test: 1 character")

	word = "RATE"
	require.True(t, IsValid(word, draw), "Test: fully valid word")

	word = "rate"
	require.True(t, IsValid(word, draw), "Test: fully valid word (lowercase)")

	word = "TRAIL"
	require.False(t, IsValid(word, draw), "Test: dict word, but not drawn letters")

	draw = "GWDUOEW"

	word = "wood"
	require.False(t, IsValid(word, draw), "Test: dict word, but repeats letter")

	// Edge cases that found errors in logic
	draw = "UVCYIEO"

	word = "voice"
	require.True(t, IsValid(word, draw), "Test: Logic error #1, using tiles once")

}

func TestFindOptimalWords(t *testing.T) {
	draw := "SENIRY"

	bestWords, score := FindOptimalWords(draw)

	fmt.Println(bestWords)

	require.ElementsMatch(t, bestWords, []string{"RESINY"})
	require.Equal(t, score, 9)
}
