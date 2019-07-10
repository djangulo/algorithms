package main

import (
	"math" // used for MaxInt32 value
	"testing"
)

func TestInsertionSort(t *testing.T) {
	t.Run("it sorts the array with distinct values", func(t *testing.T) {
		in := []int{4, 6, 1, 0, 3}
		want := []int{0, 1, 3, 4, 6}

		got := InsertionSort(in)
		AssertIntArraysEqual(t, got, want)
	})
	t.Run("it sorts the array with non-distinct values", func(t *testing.T) {
		in := []int{4, 6, 1, 0, 3, 3, 4, 7, 10}
		want := []int{0, 1, 3, 3, 4, 4, 6, 7, 10}

		got := InsertionSort(in)
		AssertIntArraysEqual(t, got, want)
	})
	t.Run("gracefully handles empty array", func(t *testing.T) {
		in := []int{}
		want := []int{}

		got := InsertionSort(in)
		AssertIntArraysEqual(t, got, want)
	})
	t.Run("gracefully handles length one array", func(t *testing.T) {
		in := []int{12542524545}
		want := []int{12542524545}

		got := InsertionSort(in)
		AssertIntArraysEqual(t, got, want)
	})
}

const MaxInt32 = math.MaxInt32

func BenchmarkInsertionSort(b *testing.B) {
	b.Run("n=100 array", func(b *testing.B) {
		in := MakeRandIntArray(100, MaxInt32)
		for i := 0; i < b.N; i++ {
			InsertionSort(in)
		}
	})
	b.Run("n=1000 array", func(b *testing.B) {
		in := MakeRandIntArray(1000, MaxInt32)
		for i := 0; i < b.N; i++ {
			InsertionSort(in)
		}
	})
	b.Run("n=100,000 array", func(b *testing.B) {
		in := MakeRandIntArray(100000, MaxInt32)
		for i := 0; i < b.N; i++ {
			InsertionSort(in)
		}
	})
}
