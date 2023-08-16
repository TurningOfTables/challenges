package count_digits

import (
	"bytes"
)

func countDigits(input string) int {
	digits := []byte("0123456789")
	var count int

	for _, v := range input {
		if bytes.ContainsRune(digits, v) {
			count++
		}
	}
	return count
}
