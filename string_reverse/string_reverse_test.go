package string_reverse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringReverse(t *testing.T) {
	input := "How quickly daft jumping zebras vex!"

	output := StringReverse(input)
	expectedOutput := "!xev sarbez gnipmuj tfad ylkciuq woH"

	assert.Equal(t, expectedOutput, output)
}

func BenchmarkStringReverse(b *testing.B) {
	input := "How quickly daft jumping zebras vex!"
	for n := 0; n < b.N; n++ {
		StringReverse(input)
	}
}

func ExampleStringReverse() {
	fmt.Println(StringReverse("How quickly daft jumping zebras vex!"))
	// Output: !xev sarbez gnipmuj tfad ylkciuq woH
}
