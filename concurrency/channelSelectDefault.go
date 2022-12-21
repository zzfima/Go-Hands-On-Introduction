package main

import "fmt"

func mainSelectDefault() {
	signallingChan := make(chan struct{})

	//without default we will stuck on first case, because it no goroutine and nobody send signal
	select {
	case <-signallingChan:
		fmt.Println("done by signalling")
	default:
		fmt.Println("done by default mechanism")
	}
}
