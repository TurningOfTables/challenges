// package fizzbuzz provides variations on solutions to the FizzBuzz game
package fizzbuzz

import (
	"fmt"
	"strconv"
)

// Fizzbuzz1 loops n times starting at 1 and returns a string slice of results.
// Uses append and fmt.Sprint
func FizzBuzz1(n int) []string {
	var results []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			results = append(results, "FizzBuzz")
			continue
		}

		if i%3 == 0 {
			results = append(results, "Fizz")
			continue
		}

		if i%5 == 0 {
			results = append(results, "Buzz")
			continue
		}

		results = append(results, fmt.Sprint(i))
	}

	return results
}

// Fizzbuzz2 loops n times starting at 1 and returns a string slice of results.
// Makes the string slice up front and assigns results to index, also uses strconv.Itoa
func FizzBuzz2(n int) []string {
	var results = make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			results[i-1] = "FizzBuzz"
			continue
		}

		if i%3 == 0 {
			results[i-1] = "Fizz"
			continue
		}

		if i%5 == 0 {
			results[i-1] = "Buzz"
			continue
		}

		results[i-1] = strconv.Itoa(i)
	}
	return results
}

// Fizzbuzz2 loops n times starting at 1 and returns a string slice of results.
// Uses strconv.Itoa, uses %15 instead of %3 && %5
func FizzBuzz2a(n int) []string {
	var results = make([]string, n)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			results[i-1] = "FizzBuzz"
			continue
		}

		if i%3 == 0 {
			results[i-1] = "Fizz"
			continue
		}

		if i%5 == 0 {
			results[i-1] = "Buzz"
			continue
		}

		results[i-1] = strconv.Itoa(i)
	}
	return results
}

// Fizzbuzz3 loops n times starting at 1 and returns a string slice of results.
// Makes the string slice up front and starts by checking for non fizz or buzz answers
func FizzBuzz3(n int) []string {
	var results = make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 != 0 && i%5 != 0 {
			results[i-1] = fmt.Sprint(i)
			continue
		}

		if i%3 == 0 && i%5 == 0 {
			results[i-1] = "FizzBuzz"
			continue
		}

		if i%3 == 0 {
			results[i-1] = "Fizz"
			continue
		}

		results[i-1] = "Buzz"

	}

	return results
}

// Fizzbuzz4 loops n times starting at 1 and returns a string slice of results.
// Makes the string slice up front and starts by checking for non fizz or buzz answers
// Sets the index up front instead of i-1 each time
func FizzBuzz4(n int) []string {
	var results = make([]string, n)

	for i := 1; i <= n; i++ {
		index := i - 1
		if i%3 != 0 && i%5 != 0 {
			results[index] = fmt.Sprint(i)
			continue
		}

		if i%3 == 0 && i%5 == 0 {
			results[index] = "FizzBuzz"
			continue
		}

		if i%3 == 0 {
			results[index] = "Fizz"
			continue
		}

		results[index] = "Buzz"

	}

	return results
}

// Fizzbuzz5 loops n times starting at 1 and returns a string slice of results.
// Makes the string slice up front
// Uses string concatenation to produce FizzBuzz where needed
func FizzBuzz5(n int) []string {
	var results = make([]string, n)

	for i := 1; i <= n; i++ {
		index := i - 1
		if i%3 == 0 {
			results[index] = "Fizz"
		}

		if i%5 == 0 {
			results[index] += "Buzz"
		}

		if results[index] == "" {
			results[index] = fmt.Sprint(i)
		}
	}
	return results
}

// Fizzbuzz5 loops n times starting at 1 and returns a string slice of results.
// Makes the string slice up front
// Uses string concatenation to produce FizzBuzz where needed and strconv.Itoa
func FizzBuzz6(n int) []string {
	var results = make([]string, n)

	for i := 1; i <= n; i++ {
		index := i - 1
		if i%3 == 0 {
			results[index] = "Fizz"
		}

		if i%5 == 0 {
			results[index] += "Buzz"
		}

		if results[index] == "" {
			results[index] = strconv.Itoa(i)
		}
	}
	return results
}
