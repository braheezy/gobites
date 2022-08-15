package _4_twitter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var doc1 = "Data is the oil of the digital economy"
var doc2 = "Data is a new oil"

const FloatPrecision = 0.001

func TestMakeCorpus(t *testing.T) {
	corpus := MakeCorpus(doc1, doc2)

	require.ElementsMatch(
		t,
		corpus,
		[]string{"data", "digital", "economy", "is", "new", "of", "oil", "the"},
	)
}

func TestToVector(t *testing.T) {
	corpus := MakeCorpus(doc1, doc2)

	require.ElementsMatch(
		t,
		ToVector(doc1, corpus),
		[]int{1, 1, 1, 1, 0, 1, 1, 2},
	)

	require.ElementsMatch(
		t,
		ToVector(doc2, corpus),
		[]int{1, 0, 0, 1, 1, 0, 1, 0},
	)
}

func TestDotProduct(t *testing.T) {
	v1 := []int{1, 3, 5, 2}
	v2 := []int{4, -2, 1, 8}

	require.Equal(t, 19, DotProduct(v1, v2))

	v1 = []int{1, -3, 5, 5}
	v2 = []int{4, -2, 1}

	require.Equal(t, 0, DotProduct(v1, v2))

	v1 = []int{1, 1, 1, 1, 0, 1, 1, 2}
	v2 = []int{1, 0, 0, 1, 1, 0, 1, 0}

	require.Equal(t, 3, DotProduct(v1, v2))
}

func TestMagnitude(t *testing.T) {
	testVector := []int{2, 4, 1}
	require.InDelta(t, Magnitude(testVector), 4.583, FloatPrecision)

	testVector = []int{2, 0, 2}
	require.InDelta(t, Magnitude(testVector), 2.828, FloatPrecision)

	testVector = []int{0, 0, 0}
	require.InDelta(t, Magnitude(testVector), 0.0, FloatPrecision)

	testVector = []int{1, 1, 1, 1, 0, 1, 1, 2}
	require.InDelta(t, Magnitude(testVector), 3.16227766, FloatPrecision)

	testVector = []int{1, 0, 0, 1, 1, 0, 1, 0}
	require.InDelta(t, Magnitude(testVector), 2.0, FloatPrecision)
}

func TestCosineSimilarity(t *testing.T) {

	require.InDelta(t, CosineSimilarity(doc1, doc2), 0.4743, FloatPrecision)
}
