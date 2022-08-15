package _4_twitter

import (
	"math"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

/*
A native implementation of cosine similarity. Compare 2 `documents` and get a score back.

This is smart about comparing docs of different lengths but still isn't all that great b/c it doesn't consider semantics e.g. 'Hello' and 'Hi' are considered very different from each other.
*/

// One of the first things needed is a vector representation of the documents. We're using count vectorization.
type CountVector []int

// Create the corpus, the total set of words to process
func MakeCorpus(doc1, doc2 string) (corpus []string) {
	// Get unique words by using map
	corpusMap := make(map[string]bool)
	for _, word := range strings.Fields(doc1) {
		corpusMap[word] = true
	}
	for _, word := range strings.Fields(doc2) {
		corpusMap[word] = true
	}

	// Now get keys from map but only if they meet condition aka light tokenization
	for word := range corpusMap {
		if len(word) > 1 {
			corpus = append(corpus, strings.ToLower(word))
		}
	}
	// Enforce some order in the universe for consistency.
	sort.Strings(corpus)

	return corpus
}

// Return vector representation of document.
func ToVector(doc string, corpus []string) []int {
	/*
		Make a 0-filled array based on corpus words.
		Iterate over document words and increment array based on location in corpus.
	*/
	vector := make([]int, len(corpus))
	for _, word := range strings.Fields(doc) {
		if i := slices.Index[string](corpus, strings.ToLower(word)); i != -1 {
			vector[i] += 1
		}
	}
	return vector
}

// Multiply corresponding values in vectors together, sum total
func DotProduct(v1, v2 []int) int {
	if len(v1) != len(v2) {
		return 0
	}

	sum := 0
	for i, value := range v1 {
		sum += value * v2[i]
	}

	return sum
}

// Sqrt of sum of squares
func Magnitude(vector []int) float64 {
	if len(vector) == 0 {
		return 0.0
	}

	total := 0.0
	for _, value := range vector {
		total += float64(value * value)
	}

	return math.Sqrt(total)
}

// CosSim is defined as DotProduct of documents of product of Magnitude of documents
func CosineSimilarity(doc1, doc2 string) float64 {
	// Make a new corpus based on docs
	corpus := MakeCorpus(doc1, doc2)

	// Vectorize the docs now that corpus is known
	docVector1, docVector2 := ToVector(doc1, corpus), ToVector(doc2, corpus)

	// Get Dot Product of documents for numerator
	dotProd := DotProduct(docVector1, docVector2)

	// Get product of Magnitude of documents
	magProd := Magnitude(docVector1) * Magnitude(docVector2)

	// Compute and return cosine similarity, after some checks
	if magProd == 0 {
		// Don't divide by zero. One of the docs was weird
		return 0
	}
	return float64(dotProd) / magProd
}
