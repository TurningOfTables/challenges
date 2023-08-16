// Package count_digits provides functionality to determine how many digits are in a string
package count_digits

import (
	"bytes"
)

// CountDigits takes an input string and returns an int of the number of digits within that string
func CountDigits(input string) int {
	digits := []byte("0123456789")
	var count int

	for _, v := range input {
		if bytes.ContainsRune(digits, v) {
			count++
		}
	}
	return count
}
