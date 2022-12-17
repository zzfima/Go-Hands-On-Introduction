package main

import "fmt"

type nums interface {
	int | float64
}

func sum1[T nums](a, b T) T {
	return a + b
}

func mainInterfaceConstrains() {
	fmt.Println(sum1(3, 3))

}
