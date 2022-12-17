package main

import "fmt"

type fig struct {
	space int
}

func (f1 fig) Comparable(f2 fig) bool {
	return f1.space > f2.space
}

func mainStlSupport() {
	f1 := fig{11}
	f2 := fig{2}

	fmt.Println(f1.Comparable(f2))
}
