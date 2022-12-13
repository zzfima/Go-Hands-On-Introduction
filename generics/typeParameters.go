package main

import (
	"fmt"
)

func sumIntegers(a, b int) int {
	return a + b
}

func sumFloats(a, b float64) float64 {
	return a + b
}

func sumGeneric[T int | float64](a, b T) T {
	return a + b
}

// Tilde supports custom types based on those (int, float) types
func minusGeneric[T ~int | ~float64](a, b T) T {
	return a - b
}

type myInt int

type list[T any] struct {
	next *list[T]
	val  T
}

func mainTypeParameters() {
	fmt.Println(sumFloats(3.0, 4.0))
	fmt.Println(sumIntegers(3, 4))
	fmt.Println(sumGeneric[int](3, 4))
	fmt.Println(sumGeneric(3.0, 4.0))

	m1 := myInt(3)
	m2 := myInt(1)
	//fmt.Println(sumGeneric(m1, m2)) not compiled - no tilde!
	fmt.Println(minusGeneric(m1, m2))

}
