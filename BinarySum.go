package main

import (
	"errors"
)

var DifferentArrayLengthError = errors.New("Arrays A and B have different lengths.")

func BinarySum(A, B []int) ([]int, error) {
	n := len(A)
	if n != len(B) {
		err := DifferentArrayLengthError
		return nil, err
	}
	C := make([]int, n+1)
	var temp, carry int
	carry = 0
	for j := n - 1; j > -1; j-- {
		temp = A[j] + B[j] + carry
		C[j+1] = temp % 2
		carry = temp / 2
	}
	C[0] = carry
	return C, nil
}
