package main

import (
	"fmt"
)

func f1() {
	fmt.Println("i am goroutine")
}

func mainGoroutines() {
	go f1()
	fmt.Println("1...")
	fmt.Println("2...")
	fmt.Println("3...")
	fmt.Println("4...")
}
