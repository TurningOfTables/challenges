package matchinttotal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchIntTotal(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 2, 3},
			expected: 4,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			input:    []int{9, 2, 5, 5, 4, 4},
			expected: 18,
		},
		{
			input:    []int{0, 0, 1, 2},
			expected: 0,
		},
		{
			input:    []int{-1, -1, 2, 3},
			expected: -2,
		},
	}

	for _, test := range tests {
		res := matchIntTotal(test.input)
		assert.Equal(t, test.expected, res)
	}
}

func BenchmarkMatchIntTotal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		matchIntTotal([]int{1, 2, 2, 3, 5})
	}
}
