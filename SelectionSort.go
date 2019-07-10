package main

import "fmt"

/*
SelectionSort :
Sorts by first finding the smallest value in the array (initially assuming A[0]
is the smallest), and swapping it for A[0], then repeating until the array is
sorted.
*/
func SelectionSort(A []int, display bool) []int {
	var i, j, min, tmp int
	n := len(A)
	for j = 0; j < n-1; j++ {
		if display == true {
			fmt.Printf("%v", A)
		}
		min = j
		for i = j + 1; i < n; i++ {
			if A[i] <= A[min] {
				min = i
			}
		}
		tmp = A[min]
		A[min] = A[j]
		A[j] = tmp
		if display == true {
			fmt.Printf(" => %v\n", A)
		}
	}
	return A
}
