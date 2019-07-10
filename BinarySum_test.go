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
	benches := []struct {
		name string
		n    int
	}{
		{"100 digit sum", 100},
		{"1,000 digit sum", 1000},
		{"100,000 digit sum", 100000},
	}
	for _, bench := range benches {
		A, B := MakeRandIntArray(bench.n, 2), MakeRandIntArray(bench.n, 2)
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BinarySum(A, B)
			}
		})
	}
}
