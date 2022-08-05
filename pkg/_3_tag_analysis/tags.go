package _3_tag_analysis

import (
	"embed"
	"fmt"
	"sort"

	"github.com/mmcdole/gofeed"
)

//go:embed rss.xml
var f embed.FS
var RSSFile, _ = f.ReadFile("rss.xml")

const TopNumber = 10
const SimilarCutoff = 0.75

type Tags map[string]int

// Parse RSS tags out of file
func GetTags() (tags Tags) {
	fp := gofeed.NewParser()
	rssFeed, _ := fp.ParseString(string(RSSFile))

	tags = make(map[string]int)

	for _, item := range rssFeed.Items {
		for _, category := range item.Categories {
			tags[category] += 1
		}
	}
	return tags
}

// Return the N most common tags
func GetTopTags(tags Tags, n int) []string {
	keys := make([]string, 0, len(tags))

	for key := range tags {
		keys = append(keys, key)
	}

	// SliceStable sorts the 'keys' by a function we define
	// We use the value in the 'tags' map to figure it out. Neat!
	sort.SliceStable(keys, func(i, j int) bool {
		return tags[keys[i]] > tags[keys[j]]
	})
	return keys[:n]
}

// Return how similar s1 is to s2
func Similarity(s1 string, s2 string) float64 {
	// Protect divide by zero
	if len(s1) == 0 {
		return 0
	}
	// Computer similarity based on HammingDistance
	distance := HammingDistance(s1, s2)
	return 1 - (float64(distance) / float64(len(s1)))
}

// Return number of positions at which the corresponding symbols are different
func HammingDistance(s1 string, s2 string) (distance int) {
	// Hamming distance assumes equal strings, so any extra characters are counted as a "ding"
	// https://en.wikipedia.org/wiki/Hamming_distance

	// Check if both strings are empty
	if len(s1) == 0 && len(s2) == 0 {
		return 0
	}

	// Knowing which word is shorter will help
	shorterWord := s1
	longerWord := s2
	if len(s2) < len(s1) {
		shorterWord = s2
		longerWord = s1
	}
	// Initialize distance to difference in string lengths
	distance = len(longerWord) - len(shorterWord)
	for i, c := range shorterWord {
		if rune(longerWord[i]) != c {
			distance += 1
		}
	}
	return distance
}

// Return tags with a similarity score >= cutoff. {baseWord: [similarWords...]}
func GetSimilarities(tags Tags, cutoff float64) map[string][]string {
	// Get the list of words, the keys of the Tags map
	words := make([]string, 0, len(tags))

	for word := range tags {
		words = append(words, word)
	}
	sort.Strings(words)

	similarities := make(map[string][]string)
	// Compare each word to every word after it
	// Otherwise, we get illuminating results like:
	// 		pip is similar to: [zip]
	// 		zip is similar to: [pip]
	for i := 0; i < len(words)-1; i++ {
		baseWord := words[i]
		for _, nextWord := range words[i+1:] {
			if Similarity(baseWord, nextWord) >= cutoff {
				similarities[baseWord] = append(similarities[baseWord], nextWord)
			}
		}
	}
	return similarities
}

func Run() {
	tags := GetTags()
	topTags := GetTopTags(tags, TopNumber)
	fmt.Printf("Top %v tags:\n", TopNumber)
	for _, topTag := range topTags {
		fmt.Printf("%15s: %v\n", topTag, tags[topTag])
	}

	similarTags := GetSimilarities(tags, SimilarCutoff)
	for baseTag, similarTags := range similarTags {
		fmt.Printf("%s is similar to: %v\n", baseTag, similarTags)
	}
}
