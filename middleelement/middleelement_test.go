package middleelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleElement(t *testing.T) {
	example := []int{1, 2, 3, 4, 5}
	res := MiddleElement(example)
	assert.Equal(t, 3, res)
}

func BenchmarkMiddleElement(b *testing.B) {
	example := []int{1, 2, 3, 4, 5}

	for n := 0; n < b.N; n++ {
		MiddleElement(example)
	}
}

func ExampleMiddleElement() {
	fmt.Println(MiddleElement([]int{1, 2, 3, 4, 5}))
	// Output: 3
}
