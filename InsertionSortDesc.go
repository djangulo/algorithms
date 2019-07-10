/*
Input:  A sequence of n numbers [a1, a2, ..., an]
Output: A permutation (reordering) [a1', a2', ..., an'] of the input sequence such that
*/

package main

// InsertionSort sorts by insertion
func InsertionSortDesc(A []int) []int {
	var i, j int
	n := len(A)
	for j = 1; j < n; j++ {
		key := A[j]
		i = j - 1
		for i > -1 && A[i] < key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	return A
}
