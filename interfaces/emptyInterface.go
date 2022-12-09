package main

import "fmt"

type emptyInterface interface{}

func mainEmptyInterface() {
	s1 := "mama hello"
	do(s1)

	f1 := 5.678
	do(f1)

	fmt.Println(doFormat(f1))
}

func do(ei emptyInterface) {
	fmt.Println(ei)
}

func doFormat(val interface{}) string {
	return fmt.Sprintf("%v, %T", val, val)
}
