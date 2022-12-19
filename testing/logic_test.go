package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var nums []int

func TestAddElementsSlice1(t *testing.T) {
	res := sumElementsSlice(nums...)
	require.Equal(t, 12, res)
}

func TestAddElementsSlice2(t *testing.T) {
	res := sumElementsSlice(3, 4, 5)
	require.Equal(t, 12, res)
}

func TestAddElementsArray(t *testing.T) {
	res := sumElementsArray(nums)
	require.Equal(t, 12, res)
}

func TestMain(m *testing.M) {
	fmt.Println("*** SETUP ***")
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)

	m.Run()
	fmt.Println("*** TEARDOWN ***")
}
