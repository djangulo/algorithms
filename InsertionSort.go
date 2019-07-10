package main

import "fmt"

// InsertionSort sorts by insertion
func InsertionSort(A []int, display bool) []int {
	var i, j int
	n := len(A)
	for j = 1; j < n; j++ {
		if display && n <= 10 {
			fmt.Printf("%v =>", A)
		}
		key := A[j]
		i = j - 1
		for i > -1 && A[i] > key {
			A[i+1] = A[i]
			if display && n <= 10 {
				fmt.Printf(" - %v - ", A)
			}
			i = i - 1
		}
		A[i+1] = key
		if display && n <= 10 {
			fmt.Printf(" => %v\n", A)
		}
	}
	return A
}
