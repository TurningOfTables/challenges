// package sortslice provides ascend and descend sort functionality for slices of ints
package sortslice

import "sort"

// SortAsc takes a slice of ints and sorts it in ascending order
func SortAsc(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

// SortDesc takes a slice of ints and sorts it in descending order
func SortDesc(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s
}
