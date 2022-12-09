package main

import (
	"fmt"
)

func mainTypeSwitches() {
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI(1.12))
}

func whatAmI(val interface{}) string {
	switch val.(type) {
	case string:
		return "i am string"
	case int:
		return "i am int"
	default:
		return fmt.Sprintf("i am not int not a string. I am %T", val)
	}
}
