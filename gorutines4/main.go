package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func increment(num *int64) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}
func main() {
	var votes int64 = 0

	for i := 0; i < 1000; i++ {
		go increment(&votes)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d\n", votes)
}
