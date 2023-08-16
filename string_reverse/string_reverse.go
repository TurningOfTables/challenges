// Package string_reverse provides functionality to reverse a string (surprisingly)
package string_reverse

// StringReverse takes a string, creates a slice of bytes, addes each character in reverse order, then returns a new string containing the result
func StringReverse(input string) string {
	var output []byte
	for i := len(input); i > 0; i-- {
		output = append(output, input[i-1])
	}

	return string(output)
}
