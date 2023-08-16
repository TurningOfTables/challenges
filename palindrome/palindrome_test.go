package palindrome

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var testPalindromes = []string{"kayak", "deified", "rotator", "repaper", "deed", "peep", "wow", "racecar", "tacocat", "wontloversrevoltnow"}
	var testNotPalindromes = []string{"campus", "invest", "shop", "law", "celebrate", "catfish", "jackpot", "sublime", "future", "picnic"}

	for _, palindrome := range testPalindromes {
		res := isPalindrome(palindrome)
		assert.True(t, res, "isPalindrome(%v) - expected true but got %v", palindrome, res)
	}

	for _, notPalindrome := range testNotPalindromes {
		res := isPalindrome(notPalindrome)
		assert.False(t, res, "isPalindrom(%v) - expected false but got %v", notPalindrome, res)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	var mixed = []string{"kayak", "campus", "deified", "invest"}
	for n := 0; n < b.N; n++ {
		isPalindrome(mixed[rand.Intn(len(mixed))])
	}
}
