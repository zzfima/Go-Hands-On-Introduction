package main

import (
	"fmt"
	"time"
)

func mainChannelsLoopOnSelect() {
	signallingChan := make(chan struct{})
	timeChannel1 := time.After(1200 * time.Millisecond)
	timeChannel2 := time.After(1250 * time.Millisecond)

	go func() {
		for {
			select {
			case <-signallingChan:
				fmt.Println("done by signalling")
			case <-timeChannel1:
				go func() {
					signallingChan <- struct{}{}
				}()
				fmt.Println("done by timeChannel1")
			case <-timeChannel2:
				fmt.Println("done by timeChannel2")
			default:
				fmt.Println("done by default")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	fmt.Println("waiting..")

	time.Sleep(4 * time.Second)
}
