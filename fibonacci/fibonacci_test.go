package fibonacci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	res := Fibonacci(10)
	assert.Equal(t, expected, res)
}

func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(100)
	}
}

func ExampleFibonacci() {
	fmt.Println(Fibonacci(10))
	// Output: [0 1 1 2 3 5 8 13 21 34]
}
