// package middleelement provides functionality to get the middle element of an int slice
package middleelement

// Middleelement takes a []int as input, and returns the int that's in the middle
// If there is no 'middle' such as in a []int with 2 elements, it returns the first
func MiddleElement(input []int) int {
	return input[len(input)/2]
}
