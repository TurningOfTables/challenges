package reverseslice

import "golang.org/x/exp/slices"

func reverseSlice(s []int) []int {
	slices.Reverse(s)
	return s
}

func reverseSlice1(s []int) []int {
	var revSlice []int

	for x := len(s); x > 0; x-- {
		revSlice = append(revSlice, s[x-1])
	}
	return revSlice
}

func reverseSlice2(s []int) []int {
	var revSlice = make([]int, len(s))

	for x := len(s); x > 0; x-- {
		revSlice[len(s)-x] = s[x-1]
	}
	return revSlice
}

func reverseSlice3(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
