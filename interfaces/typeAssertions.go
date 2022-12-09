package main

import "fmt"

func mainAssertions() {
	var i1 interface{} = 4.56
	var i2 any = "7.89 money"

	fmt.Printf("%v %T\n", i1, i1)
	fmt.Println(i1.(float64)) //do assertion to float64 and print value

	fmt.Printf("%v %T\n", i2, i2)
	fmt.Println(i2.(string)) //do assertion to string print value

	//check assertion
	if _, converted := i1.(string); converted == false {
		fmt.Println("i1 is not a string")
	}
}
