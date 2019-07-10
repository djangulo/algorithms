package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func AssertIntArraysEqual(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

// MakeRandIntArray makes a random int array of length n
func MakeRandIntArray(length, max int) []int {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(max)
	}
	return array
}

func AssertError(t *testing.T, got, want error) {
	t.Helper()
	if want == nil {
		t.Fatalf("wanted an error but didn't get one")
	}
	if got != want {
		t.Fatalf("got '%s' want '%s'", got, want)
	}
}
