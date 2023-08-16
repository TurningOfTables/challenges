// Package non_matching_characters provides functionality to determine the all characters which are not present in both provided strings
package non_matching_characters

import "strings"

// NonMatchingCharacters takes two strings, and returns a string slice containing any characters not present in both strings
func NonMatchingCharacters(x, y string) []string {
	var nonMatching []string
	for _, v := range x {
		if !strings.ContainsRune(y, v) {
			nonMatching = append(nonMatching, string(v))
		}
	}

	for _, v := range y {
		if !strings.ContainsRune(x, v) {
			nonMatching = append(nonMatching, string(v))
		}
	}
	return nonMatching
}
