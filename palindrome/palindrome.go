package palindrome

// isPalindrome takes a string and reverses it, then compares it to the original. If both match it returns true, otherwise false
func isPalindrome(input string) bool {
	var reversed []byte

	for i := len(input); i > 0; i-- {
		reversed = append(reversed, input[i-1])
	}

	return input == string(reversed)
}
