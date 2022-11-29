package main

import (
	"fmt"
	"reflect"
)

const a1 = "a"
const a2 = 'a'
const a3 rune = 'a'

func main() {
	fmt.Printf("%s %T\n", a1, a1)
	fmt.Printf("%c %s\n", a2, reflect.TypeOf(a2))
	fmt.Printf("%c %T\n", a3, a3)
	if reflect.TypeOf(a1).Kind() == reflect.String {
		fmt.Println("a1 is string!")
	}
}
