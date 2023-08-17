package fizzbuzz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	functions := map[string]func(n int) []string{
		"FizzBuzz 1":             FizzBuzz1,
		"FizzBuzz 2 - %3 and %5": FizzBuzz2,
		"FizzBuzz 2a - %15":      FizzBuzz2a,
		"FizzBuzz 3":             FizzBuzz3,
		"FizzBuzz 4":             FizzBuzz4,
		"FizzBuzz 5":             FizzBuzz5,
		"FizzBuzz 6":             FizzBuzz6,
	}
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

	for name, fn := range functions {
		t.Run(name, func(t *testing.T) {
			res := fn(15)
			assert.Equal(t, expected, res)
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	functions := map[string]func(n int) []string{
		"FizzBuzz 1":             FizzBuzz1,
		"FizzBuzz 2 - %3 and %5": FizzBuzz2,
		"FizzBuzz 2a - %15":      FizzBuzz2a,
		"FizzBuzz 3":             FizzBuzz3,
		"FizzBuzz 4":             FizzBuzz4,
		"FizzBuzz 5":             FizzBuzz5,
		"FizzBuzz 6":             FizzBuzz6,
	}
	for name, fn := range functions {
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fn(1000)
			}
		})
	}
}

func ExampleFizzBuzz1() {
	fmt.Println(FizzBuzz1(15))
	// Output: [1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz]
}
