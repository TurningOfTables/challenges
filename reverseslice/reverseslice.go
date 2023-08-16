package reverseslice

import "golang.org/x/exp/slices"

// ReverseSlice takes a slice of ints and returns a slice of ints in reverse order
func ReverseSlice(s []int) []int {
	slices.Reverse(s)
	return s
}

// ReverseSlice1 takes a slice of ints and returns a slice of ints in reverse order
func ReverseSlice1(s []int) []int {
	var revSlice []int

	for x := len(s); x > 0; x-- {
		revSlice = append(revSlice, s[x-1])
	}
	return revSlice
}

// ReverseSlice2 takes a slice of ints and returns a slice of ints in reverse order
func ReverseSlice2(s []int) []int {
	var revSlice = make([]int, len(s))

	for x := len(s); x > 0; x-- {
		revSlice[len(s)-x] = s[x-1]
	}
	return revSlice
}

// ReverseSlice4 takes a slice of ints and returns a slice of ints in reverse order
func ReverseSlice3(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
