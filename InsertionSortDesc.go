package main

import "fmt"

// InsertionSortDesc sorts by insertion in descending order
func InsertionSortDesc(A []int, display bool) []int {
	var i, j int
	n := len(A)
	for j = 1; j < n; j++ {
		if display {
			fmt.Printf("%v", A)
		}
		key := A[j]
		i = j - 1
		for i > -1 && A[i] < key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
		if display {
			fmt.Printf(" => %v\n", A)
		}
	}
	return A
}
