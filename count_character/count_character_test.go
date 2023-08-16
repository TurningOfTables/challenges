package count_character

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCharacter(t *testing.T) {
	var tests = []struct {
		needle   string
		haystack string
		expected int
	}{
		{
			needle:   "a",
			haystack: "cat",
			expected: 1,
		},
		{
			needle:   "a",
			haystack: "abacus",
			expected: 2,
		},
		{
			needle:   "c",
			haystack: "dog",
			expected: 0,
		},
		{
			needle:   "f",
			haystack: "foofoofoo",
			expected: 3,
		},
		{
			needle:   "",
			haystack: "bar",
			expected: 0,
		},
	}

	for _, test := range tests {
		res := CountCharacter(test.needle, test.haystack)
		assert.Equal(t, test.expected, res)
	}
}

func BenchmarkCountCharacter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountCharacter("a", "83gnsdawondfivamnoicsvaosicv")
	}
}

func ExampleCountCharacter() {
	fmt.Println(CountCharacter("a", "abacus"))
	// Output: 2
}
