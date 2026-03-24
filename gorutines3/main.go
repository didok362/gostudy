package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(num *int, Mutex *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		Mutex.Lock()
		*num++
		Mutex.Unlock()
	}
}

func main() {
	votes := 0
	var Mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go increment(&votes, &Mutex)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d\n", votes)
}
