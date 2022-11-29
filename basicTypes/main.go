package main

import (
	"fmt"
	"reflect"
)

const a1 = "a"
const a2 = 'a'
const a3 rune = 'a'

func main() {
	//types
	fmt.Printf("%s %T\n", a1, a1)
	fmt.Printf("%c %s\n", a2, reflect.TypeOf(a2))
	fmt.Printf("%c %T\n", a3, a3)
	if reflect.TypeOf(a1).Kind() == reflect.String {
		fmt.Println("a1 is string!")
	}

	//pointers
	a4 := 32
	a5 := a4
	a5 = 11
	fmt.Printf("a4 = %v a5 = %v\n", a4, a5)

	var a6 *int
	a6 = &a4
	(*a6) = 99
	fmt.Printf("a4 = %v *a4 = %v *a6 = %v a6 = %v\n", a4, &a4, *a6, a6)
}
