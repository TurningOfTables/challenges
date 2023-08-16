package count_character

import "strings"

// CountCharacter takes a string (needle) and returns the number of times it appears in the string (haystack)
func CountCharacter(needle, haystack string) int {
	if needle == "" {
		return 0
	} // Avoids strings.Count returning 1 + Unicode if needle is empty string
	return strings.Count(haystack, needle)
}
