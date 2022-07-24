package main

import (
	"testing"
)

var TestWords = map[string]int{"bob": 7, "julian": 13, "PyBites": 14, "quit": 13, "barbeque": 21, "benzalphenylhydrazone": 56}

func TestLoadWords(t *testing.T) {
	words, err := LoadWords()
	expected := 235886
	actual := len(words)
	if actual != expected || err != nil {
		t.Fatalf(`Got: %v. Expected: %v.`, actual, expected)
	}
}

func TestCalcWordValue(t *testing.T) {
	for word, expected_score := range TestWords {
		actual_score := CalcWordValue(word)
		if actual_score != expected_score {
			t.Fatalf(`Case: %s. Got: %v. Expected: %v.`, word, actual_score, expected_score)
		}
	}
}

func TestMaxWordValue(t *testing.T) {
	var words []string
	for word, _ := range TestWords {
		words = append(words, word)
	}

	actual := MaxWordValue(words)
	expected := "benzalphenylhydrazone"
	if actual != expected {
		t.Fatalf(`Got: %v. Expected: %v.`, actual, expected)
	}

}
