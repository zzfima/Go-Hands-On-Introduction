package main

import (
	"fmt"
	"time"
)

func f2(ch chan int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("i am 2")
	time.Sleep(120 * time.Millisecond)
	ch <- 5
}

func f3(ch chan int) {
	fmt.Println("i am 3", <-ch)
}

func f4() {
	cnt := 0
	for {
		cnt++
		fmt.Println(cnt)
		time.Sleep(100 * time.Millisecond)
	}
}

func mainChannels() {
	ch := make(chan int)

	go f2(ch)
	go f3(ch)
	go f4()

	time.Sleep(1000 * time.Millisecond)
}
