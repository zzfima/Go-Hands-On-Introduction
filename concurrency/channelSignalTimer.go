package main

import (
	"fmt"
	"time"
)

func mainSignalTimer() {
	// signalling - only for signal. Have no payload.
	// Empty struct does not store anything in memory
	signallingChan := make(chan struct{})
	s1 := struct{}{}
	go f8(signallingChan)
	fmt.Println("Start wait...")
	time.Sleep(1 * time.Second)
	signallingChan <- s1

	timeChan := time.After(1 * time.Second)
	go f9(timeChan)

	fmt.Println("wait for all...")
	time.Sleep(2 * time.Second)
}

func f9(timeChan <-chan time.Time) {
	<-timeChan
	fmt.Println("Done by time")
}

func f8(signallingChan chan struct{}) {
	<-signallingChan
	fmt.Println("Done by signal")
}
