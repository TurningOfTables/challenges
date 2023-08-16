package count_digits

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigits(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{
			input:    "12345",
			expected: 5,
		},
		{
			input:    "foobar",
			expected: 0,
		},
		{
			input:    "! jk 29 ,dkjfj20 39j cx",
			expected: 6,
		},
		{
			input:    "",
			expected: 0,
		},
	}

	for _, test := range tests {
		res := countDigits(test.input)
		assert.Equal(t, test.expected, res)
	}
}

func BenchmarkCountDigits1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		countDigits(fmt.Sprint(rand.Intn(9999)))
	}
}
