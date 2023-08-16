package vowelconsonant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVowelConsonant(t *testing.T) {

	var tests = []struct {
		input    string
		expected Info
	}{
		{
			input:    "abc",
			expected: Info{Vowels: 1, Consonants: 2},
		},
		{
			input:    "carrot",
			expected: Info{Vowels: 2, Consonants: 4},
		},
		{
			input:    "salmon",
			expected: Info{Vowels: 2, Consonants: 4},
		},
		{
			input:    "cry",
			expected: Info{Vowels: 0, Consonants: 3},
		},
		{
			input:    "ai",
			expected: Info{Vowels: 2, Consonants: 0},
		},
		{
			input:    "some spaces",
			expected: Info{Vowels: 4, Consonants: 6},
		},
		{
			input:    "punctuation!!",
			expected: Info{Vowels: 5, Consonants: 6},
		},
	}

	for _, test := range tests {
		res := VowelConsonant(test.input)
		assert.Equal(t, test.expected, res)
	}
}

func BenchmarkVowelConsonant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		VowelConsonant("qwertyuiop")
	}
}
