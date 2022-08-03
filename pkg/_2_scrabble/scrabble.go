package _2_scrabble

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/stat/combin"

	"github.com/braheezy/gobites/pkg/_1_wordvalue"
)

// var dictionaryFile, _ = filepath.Abs("pkg/_1_wordvalue/dictionary.txt")
var Dictionary, _ = _1_wordvalue.LoadWords()

const NumLetters = 7

var Vowels = []rune{'A', 'E', 'I', 'O', 'U'}

type Letter struct {
	name   rune
	amount int
	value  int
}

var distribution = []Letter{
	{name: 'A', amount: 9, value: 1},
	{name: 'B', amount: 2, value: 3},
	{name: 'C', amount: 2, value: 3},
	{name: 'D', amount: 4, value: 2},
	{name: 'E', amount: 12, value: 1},
	{name: 'F', amount: 2, value: 4},
	{name: 'G', amount: 3, value: 2},
	{name: 'H', amount: 2, value: 4},
	{name: 'I', amount: 9, value: 1},
	{name: 'J', amount: 1, value: 8},
	{name: 'K', amount: 1, value: 5},
	{name: 'L', amount: 4, value: 1},
	{name: 'M', amount: 2, value: 3},
	{name: 'N', amount: 6, value: 1},
	{name: 'O', amount: 8, value: 1},
	{name: 'P', amount: 2, value: 3},
	{name: 'Q', amount: 1, value: 10},
	{name: 'R', amount: 6, value: 1},
	{name: 'S', amount: 4, value: 1},
	{name: 'T', amount: 6, value: 1},
	{name: 'U', amount: 4, value: 1},
	{name: 'V', amount: 2, value: 4},
	{name: 'W', amount: 2, value: 4},
	{name: 'X', amount: 1, value: 8},
	{name: 'Y', amount: 2, value: 4},
	{name: 'Z', amount: 1, value: 10},
}

// Ask the Pouch for a number of tiles, get a list of random runes
// Pouch()(5) -> ['A', 'G', 'R', 'W', 'A']
func Pouch() func(int) ([]rune, int) {
	var pouch []rune
	rand.Seed(time.Now().UnixNano())

	for _, letter := range distribution {
		for i := 0; i < letter.amount; i++ {
			pouch = append(pouch, letter.name)
		}
	}
	return func(num int) (letters []rune, pouchCount int) {
		for i := 0; i < num; i++ {
			// Pick random letter by generating a random index to grab.
			index := rand.Intn(len(pouch))
			letters = append(letters, pouch[index])
		}
		return letters, len(pouch)
	}
}

func HasVowels(word string) bool {
	for _, vowel := range Vowels {
		for _, c := range word {
			if vowel == c {
				return true
			}
		}
	}
	return false
}

func IsValid(word string, draw string) (isValid bool) {
	// Word must be at least 2 characters
	if len(word) < 2 {
		fmt.Printf("%s is not enough characters...\n", word)
		return false
	}
	word = strings.ToUpper(word)
	// Word must be in dictionary
	if !slices.Contains(Dictionary, word) {
		fmt.Printf("%s is not a dictionary word...\n", word)
		return false
	}

	// Check if the each letter used is in the drawn letters
	// If so, "remove" that letter from the drawn set
	drawRunes := []rune(strings.Clone(draw))
	for _, c := range word {
		if strings.ContainsRune(string(drawRunes), c) {
			// Remove rune element from array
			j := strings.IndexRune(string(drawRunes), c)
			drawRunes = append(drawRunes[:j], drawRunes[j+1:]...)
		} else {
			fmt.Printf("%s doesn't use drawn letters...\n", word)
			return false
		}
	}
	return true
}

// Given a string of possible letters, return the best word(s) and the best score
func FindOptimalWords(currentLetters string) ([]string, int) {
	/*
		TODO: This function is slow. The test for it takes 0.866s :(

			It's probably the triple-nested for loop. How do we efficiently do functional programming
			in Go?
				- First loop: Generate strings of varying length based on currentLetters
				- Second loop: Iterate over the current set of permuted strings.
				- Third loop: For each permutation, build the possibleWord
			Thoughts:
				- Package `combin` helps but deals with strings indirectly via indices.
				- The length of the possibleWord changes
				- Can't assume words of longest length are best. Shorter words may use higher scoring letters
	*/
	var possibleWords []string

	// Start generating possible words.
	// The word lengths are from 3 to however many tiles the user has.
	for r := 3; r <= len(currentLetters); r++ {
		// Get slice of indexes of r length
		// Ex: [1, 0, 2] for r=3
		possibleWordIndices := combin.Permutations(len(currentLetters), r)
		// Iterate over all the index sets and create words from them.
		for _, possibleWordIndex := range possibleWordIndices {
			var possibleWord strings.Builder
			// Build a word by selecting the tile (by index)
			for _, index := range possibleWordIndex {
				fmt.Fprintf(&possibleWord, "%c", currentLetters[index])
			}
			// If the word we Frankenstein-ed together is an actual dictionary word, keep it
			if slices.Contains(Dictionary, possibleWord.String()) {
				possibleWords = append(possibleWords, possibleWord.String())
			}
		}
	}

	return _1_wordvalue.MaxWordsValue(possibleWords)
}

func PlayScrabble() {
	currentLetters, _ := Pouch()(NumLetters)

	fmt.Printf("Your current letters:\n\t%s\n", string(currentLetters))
	fmt.Print("Pick a word!: ")

	var UserWord string
	fmt.Scanln(&UserWord)
	for !IsValid(UserWord, string(currentLetters)) {
		fmt.Print("Pick a word: ")
		fmt.Scanln(&UserWord)
	}

	fmt.Println("Your opponent is thinking...")

	BestWords, BestScore := FindOptimalWords(string(currentLetters))
	UserScore := _1_wordvalue.CalcWordValue(UserWord)
	Grade := float64(UserScore) / float64(BestScore) * 100

	fmt.Printf("Your word '%s' has a Scrabble value of %v.\n", UserWord, UserScore)
	if len(BestWords) > 1 {
		fmt.Printf("The best words were %v and those all scored %v.\n", BestWords, BestScore)
	} else {
		fmt.Printf("The best word was %v and that scored %v.\n", BestWords, BestScore)
	}

	fmt.Printf("Your Scrabble grade: %.1f%%\n", Grade)
	fmt.Printf("Thanks for playing!\n")
}
