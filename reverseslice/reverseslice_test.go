package reverseslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctions(t *testing.T) {
	var functions = map[string]func([]int) []int{
		"Using slices.Reverse()":       reverseSlice,
		"Using append":                 reverseSlice1,
		"Using set index directly":     reverseSlice2,
		"Using slices.Reverse() (raw)": reverseSlice3,
	}

	for name, fn := range functions {
		testSlice := []int{1, 2, 3, 4}
		t.Run(name, func(t *testing.T) {
			res := fn(testSlice)
			assert.Equal(t, []int{4, 3, 2, 1}, res)
		})
	}
}

func BenchmarkFunctions(b *testing.B) {
	var functions = map[string]func([]int) []int{
		"Using slices.Reverse()":       reverseSlice,
		"Using append":                 reverseSlice1,
		"Using set index directly":     reverseSlice2,
		"Using slices.Reverse() (raw)": reverseSlice3,
	}

	for name, fn := range functions {
		testSlice := []int{1, 2, 3, 4}
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fn(testSlice)
			}
		})
	}
}
