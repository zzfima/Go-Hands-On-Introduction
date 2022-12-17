package main

import "fmt"

// Part 1: print function refactoring

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

func printGeneric[T int | string | bool](toPrint T) {
	fmt.Println(toPrint)
}

// Part 2 sum function refactoring

type constrainedValues interface {
	int | float64 | uint
}

func sumAny[T constrainedValues](vals ...T) T {
	var result T
	for _, n := range vals {
		result += n
	}
	return result
}

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

func mainGenerics() {
	// call non-generic print functions
	printString("Hello")
	printInt(42)
	printBool(true)
	printGeneric("Hello")
	printGeneric(42)
	printGeneric(true)

	// call generic printAny function for each value above

	// call sum function
	fmt.Println("result", sum([]interface{}{1, 2, 3}))
	fmt.Println("result", sumAny(1, 2, 3))
	fmt.Println("result", sumAny(1.0, 2.0, 3.01))
	// call generics sumAny function
}
