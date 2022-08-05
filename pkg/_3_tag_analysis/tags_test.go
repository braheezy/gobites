package _3_tag_analysis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTags(t *testing.T) {
	tags := GetTags()

	TotalCount := 0
	for tag, count := range tags {
		TotalCount += count
		fmt.Println(tag)
	}

	require.Equal(t, 189, TotalCount, "Number of total tags")
	require.Equal(t, 102, len(tags), "Number of unique tags")

}

func TestGetTopTags(t *testing.T) {
	tags := GetTags()
	topTags := GetTopTags(tags, TopNumber)
	require.Equal(t, TopNumber, len(topTags))
	require.Equal(t, "python", topTags[0])
}

func TestHammingDistance(t *testing.T) {
	require.Equal(t, 3, HammingDistance("karolin", "kathrin"))
	require.Equal(t, 3, HammingDistance("karolin", "kerstin"))
	require.Equal(t, 4, HammingDistance("kathrin", "kerstin"))
	require.Equal(t, 4, HammingDistance("0000", "1111"))
	require.Equal(t, 3, HammingDistance("2173896", "2233796"))
	require.Equal(t, 0, HammingDistance("", ""))
	require.Equal(t, 0, HammingDistance("a", "a"))
	require.Equal(t, 1, HammingDistance("a", "ab"))
}

func TestGetSimilarities(t *testing.T) {
	tags := GetTags()
	similarTags := GetSimilarities(tags, 0.5)

	require.Contains(t, similarTags, "pip")

	similarTags = GetSimilarities(tags, 0.75)
	require.NotContains(t, similarTags, "pip")
	require.Equal(t, similarTags["game"], []string{"games"})
}
