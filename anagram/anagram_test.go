package anagram

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	var testAnagrams = []struct {
		x string
		y string
	}{
		{"bored", "robed"},
		{"night", "thing"},
		{"dusty", "study"},
		{"a gentleman!", "elegant man!"},
	}

	for _, testAnagram := range testAnagrams {
		res := IsAnagram(testAnagram.x, testAnagram.y)
		assert.True(t, res)
	}

	var testNotAnagrams = []struct {
		x string
		y string
	}{
		{"zebra", "barcode"},
		{"snake", "danger noodle"},
		{"fluffy", "dog"},
		{"soup", "tureen"},
	}

	for _, testNotAnagram := range testNotAnagrams {
		res := IsAnagram(testNotAnagram.x, testNotAnagram.y)
		assert.False(t, res)
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	var testData = []struct {
		x string
		y string
	}{
		{"bored", "robed"},
		{"snake", "nope rope"},
	}
	for n := 0; n < b.N; n++ {
		i := testData[rand.Intn(len(testData))]
		IsAnagram(i.x, i.y)
	}
}
