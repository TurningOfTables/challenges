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
			expected: Info{vowels: 1, consonants: 2},
		},
		{
			input:    "carrot",
			expected: Info{vowels: 2, consonants: 4},
		},
		{
			input:    "salmon",
			expected: Info{vowels: 2, consonants: 4},
		},
		{
			input:    "cry",
			expected: Info{vowels: 0, consonants: 3},
		},
		{
			input:    "ai",
			expected: Info{vowels: 2, consonants: 0},
		},
		{
			input:    "some spaces",
			expected: Info{vowels: 4, consonants: 6},
		},
		{
			input:    "punctuation!!",
			expected: Info{vowels: 5, consonants: 6},
		},
	}

	for _, test := range tests {
		res := vowelCosonant(test.input)
		assert.Equal(t, test.expected, res)
	}
}

func BenchmarkVowelConsonant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		vowelCosonant("qwertyuiop")
	}
}
