package non_matching_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonMatchingCharacters(t *testing.T) {

	var tests = []struct {
		x        string
		y        string
		expected []string
	}{
		{"abc", "ab", []string{"c"}},
		{"abc", "abd", []string{"c", "d"}},
		{"abd", "abc", []string{"d", "c"}},
		{"AbC", "aBc", []string{"A", "b", "C", "a", "B", "c"}},
		{"abc", "abc", []string(nil)},
	}

	for _, test := range tests {
		res := nonMatchingCharacters(test.x, test.y)
		assert.Equal(t, test.expected, res)
	}
}

func BenchmarkNonMatchingCharacters(b *testing.B) {
	for n := 0; n < b.N; n++ {
		nonMatchingCharacters("abc", "abd")
	}
}
