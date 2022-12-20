package main

import (
	"fmt"
	"time"
)

func f5(ch chan int) {
	ch <- 5
}

func f6(ch chan int) {
	ch <- 6
}

func f7(ch chan int) {
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func mainBuffered() {
	chBuf1 := make(chan int, 2)

	go f5(chBuf1)
	go f6(chBuf1)
	time.Sleep(20 * time.Millisecond)

	go f7(chBuf1)

	time.Sleep(100 * time.Millisecond)
}
