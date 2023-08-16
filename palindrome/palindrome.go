// Package palindrome provides functionality to determine if a string is a palindrome
package palindrome

// IsPalindrome takes a string and reverses it, then compares it to the original. If both match it returns true, otherwise false
func IsPalindrome(input string) bool {
	var reversed []byte

	for i := len(input); i > 0; i-- {
		reversed = append(reversed, input[i-1])
	}

	return input == string(reversed)
}
