package anagram

import "strings"

// IsAnagram takes two strings, and returns true if all characters are present in both strings (they are anagrams)
func IsAnagram(x, y string) bool {
	for _, v := range x {
		if !strings.ContainsRune(y, v) {
			return false
		}
	}
	return true
}
