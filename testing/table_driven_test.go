package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddElements(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		{"one", []int{3, 4, 5}, 12},
		{"two", []int{6, 6}, 12},
		{"three", []int{3, 3, 3, 3}, 12},
	}

	for _, tst := range tests {
		require.Equal(t, tst.want, sumElementsSlice(tst.input...))
	}
}
