package main

import (
	"fmt"
	"sync"
	"time"
)

func mainSyncWaitGroup() {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}

	names := []string{"one", "two", "three", "four"}
	mp := make(map[string]string)
	for _, n := range names {
		wg.Add(1)
		mtx.Lock()
		go printName(n, mp, &wg, &mtx)
	}

	fmt.Println("wait all guys")
	wg.Wait()
	fmt.Println("we finished")
	fmt.Println(mp)
}

func printName(n string, mp map[string]string, wg *sync.WaitGroup, mtx *sync.Mutex) {
	time.Sleep(3 * time.Second)
	fmt.Println(n, " - good name")
	mp[n] = n + " ggg"
	(*wg).Done()
	(*mtx).Unlock()
}
