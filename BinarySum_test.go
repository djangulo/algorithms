package main

import "testing"

func TestBinarySum(t *testing.T) {
	cases := []struct {
		name string
		A    []int
		B    []int
		want []int
	}{
		{"0 + 0", []int{0}, []int{0}, []int{0, 0}},
		{"0 + 1", []int{0}, []int{1}, []int{0, 1}},
		{"1 + 0", []int{1}, []int{0}, []int{0, 1}},
		{"1 + 1", []int{1}, []int{1}, []int{1, 0}},
		{"10 + 10", []int{1, 0}, []int{1, 0}, []int{1, 0, 0}},
		{"11 + 10", []int{1, 1}, []int{1, 0}, []int{1, 0, 1}},
		{"11 + 11", []int{1, 1}, []int{1, 1}, []int{1, 1, 0}},
		{"1101 + 1011", []int{1, 1, 0, 1}, []int{1, 0, 1, 1}, []int{1, 1, 0, 0, 0}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, _ := BinarySum(c.A, c.B)
			AssertIntArraysEqual(t, got, c.want)
		})
	}
	t.Run("should return error if input lengths are different", func(t *testing.T) {
		A, B := []int{1, 1, 0}, []int{1, 1, 0, 1}
		_, err := BinarySum(A, B)
		AssertError(t, err, DifferentArrayLengthError)
	})
}

func BenchmarkBinarySum(b *testing.B) {
	b.Run("100 digit sum benchmark", func(b *testing.B) {
		A, B := MakeRandIntArray(100, 2), MakeRandIntArray(100, 2)
		for i := 0; i < b.N; i++ {
			BinarySum(A, B)
		}
	})
	b.Run("1000 digit sum benchmark", func(b *testing.B) {
		A, B := MakeRandIntArray(1000, 2), MakeRandIntArray(1000, 2)
		for i := 0; i < b.N; i++ {
			BinarySum(A, B)
		}
	})
	b.Run("100,000 digit sum benchmark", func(b *testing.B) {
		A, B := MakeRandIntArray(100000, 2), MakeRandIntArray(100000, 2)
		for i := 0; i < b.N; i++ {
			BinarySum(A, B)
		}
	})
}
