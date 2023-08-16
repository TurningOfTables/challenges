package anagram

import "strings"

func isAnagram(x, y string) bool {
	for _, v := range x {
		if !strings.ContainsRune(y, v) {
			return false
		}
	}
	return true
}
