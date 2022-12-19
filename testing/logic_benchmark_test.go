package main

import (
	"testing"
)

// BenchmarkAddElementsSlice ...
func BenchmarkAddElementsSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumElementsSlice(1, 2, 3)
	}
}

func BenchmarkAddElementsSlice2(b *testing.B) {
	nums := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		sumElementsSlice(nums...)
	}
}

func BenchmarkAddElementsArray(b *testing.B) {
	nums := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		sumElementsArray(nums)
	}
}
