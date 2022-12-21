package main

import (
	"fmt"
	"time"
)

// send only channel: ch chan<- int
// Into channel:
// chan<-
func f2(ch chan<- int) {
	time.Sleep(500 * time.Millisecond)
	time.Sleep(120 * time.Millisecond)
	ch <- 5
}

// receive only channel: ch <-chan int
// From channel:
// <-chan
func f3(ch <-chan int) {
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
	channelForReadWrite := make(chan int)
	//channelForRead := make(<-chan int)
	//channelForWrite := make(chan<- int)
	go f2(channelForReadWrite)
	go f3(channelForReadWrite)
	go f4()

	time.Sleep(1000 * time.Millisecond)
}
