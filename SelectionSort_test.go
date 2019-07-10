package main

import (
	"math" // used for MaxInt32 value
	"testing"
)

func TestSelectionSort(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"array with distinct values", []int{4, 6, 1, 0, 3}, []int{0, 1, 3, 4, 6}},
		{"array with non-distinct values", []int{4, 6, 1, 0, 3, 3, 4, 7, 10}, []int{0, 1, 3, 3, 4, 4, 6, 7, 10}},
		{"array ordered descending", []int{7, 6, 5, 4, 3, 2, 1, 0, -1, -2}, []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7}},
		{"empty array", []int{}, []int{}},
		{"length one array", []int{12542524545}, []int{12542524545}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := SelectionSort(c.in, false)
			AssertIntArraysEqual(t, got, c.want)
		})
	}

}

func BenchmarkSelectionSort(b *testing.B) {
	benches := []struct {
		name string
		n    int
	}{
		{"n=100 array", 100},
		{"n=1,000 array", 1000},
		{"n=100,000 array", 100000},
	}
	for _, bench := range benches {
		in := MakeRandIntArray(bench.n, math.MaxInt32)
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SelectionSort(in, false)
			}
		})
	}
}
