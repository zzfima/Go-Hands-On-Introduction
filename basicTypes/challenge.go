package main

import (
	"fmt"
)

var val string = "my string"

func main() {
	val := 55
	fmt.Printf("Local val: %v %T\n", val, val)
	printGlobal()
	updateGlobal()
	fmt.Printf("Address of Local val: %v %T\n", &val, &val)

	var ptr1 *int
	ptr1 = &val
	(*ptr1) = 44
	fmt.Printf("Local val: %v %T\n", val, val)
}

func printGlobal() {
	fmt.Printf("Global val: %s %T\n", val, val)
}

func updateGlobal() {
	val = "updated string"
}
