package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddElementsSlice1(t *testing.T) {
	nums := []int{3, 4, 5}
	res := sumElementsSlice(nums...)
	require.Equal(t, 12, res)
}

func TestAddElementsSlice2(t *testing.T) {
	res := sumElementsSlice(3, 4, 5)
	require.Equal(t, 12, res)
}

func TestAddElementsArray(t *testing.T) {
	nums := []int{3, 4, 5}
	res := sumElementsArray(nums)
	require.Equal(t, 12, res)
}
