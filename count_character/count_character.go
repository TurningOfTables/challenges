package count_character

import "strings"

// countCharacter takes a string (needle) and returns the number of times it appears in the string (haystack)
func countCharacter(needle, haystack string) int {
	if needle == "" {
		return 0
	} // Avoids strings.Count returning 1 + Unicode if needle is empty string
	return strings.Count(haystack, needle)
}

func countCharacter1(needle, haystack string) int {
	var count int

	for _, v := range haystack {
		if strings.ContainsRune(needle, v) {
			count++
		}
	}

	return count
}
