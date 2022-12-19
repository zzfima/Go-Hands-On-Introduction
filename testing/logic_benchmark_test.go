package main

import (
	"testing"
)

// BenchmarkAddElementsArray ...
func BenchmarkAddElementsArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumElementsSlice(1, 2, 3)
	}
}
