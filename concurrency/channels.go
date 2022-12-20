package main

import (
	"fmt"
	"time"
)

func f2(ch chan int) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("i am 2")
	ch <- 5
}

func f3(ch chan int) {
	fmt.Println("i am 3", <-ch)
}

func mainChannels() {
	ch := make(chan int)

	go f2(ch)
	go f3(ch)

	time.Sleep(100 * time.Millisecond)
}
