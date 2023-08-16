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
		res := CountDigits(test.input)
		assert.Equal(t, test.expected, res)
	}
}

func BenchmarkCountDigits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountDigits(fmt.Sprint(rand.Intn(9999)))
	}
}

func ExampleCountDigits() {
	fmt.Println(CountDigits("ab23c5"))
	// Output: 3
}
