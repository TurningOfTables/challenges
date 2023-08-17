package sortslice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortAsc(t *testing.T) {
	var slice = []int{1, 4, 5, 2, 3}
	res := SortAsc(slice)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, res)
}

func TestSortDesc(t *testing.T) {
	var slice = []int{1, 4, 5, 2, 3}
	res := SortDesc(slice)

	assert.Equal(t, []int{5, 4, 3, 2, 1}, res)
}

func BenchmarkSortAsc(b *testing.B) {
	var slice = []int{1, 4, 5, 2, 3}

	for n := 0; n < b.N; n++ {
		SortAsc(slice)
	}
}

func BenchmarkSortDesc(b *testing.B) {
	var slice = []int{1, 4, 5, 2, 3}

	for n := 0; n < b.N; n++ {
		SortDesc(slice)
	}
}

func ExampleSortAsc() {
	var slice = []int{1, 4, 5, 2, 3}
	fmt.Println(SortAsc(slice))
	// Output: [1 2 3 4 5]
}

func ExampleSortDesc() {
	var slice = []int{1, 4, 5, 2, 3}
	fmt.Println(SortDesc(slice))
	// Output: [5 4 3 2 1]
}
