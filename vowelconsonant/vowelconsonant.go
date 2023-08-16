package vowelconsonant

import (
	"strings"
)

type Info struct {
	Vowels     int
	Consonants int
}

// VowelConsonant takes an input string, and returns an Info struct containing vowels (int) and consonants (int)
// Other characters such as spaces and punctuation are ignored
func VowelConsonant(input string) Info {
	var vowels = "aeiou"
	var consonants = "bcdfghjklmnpqrstvwxyz"
	var i Info

	for _, x := range input {
		if strings.ContainsRune(vowels, x) {
			i.Vowels++
		}
		if strings.ContainsRune(consonants, x) {
			i.Consonants++
		}
	}

	return i

}
