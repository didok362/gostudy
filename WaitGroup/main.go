package main

import (
	"math/rand"
	"sync"
	"time"
)

func Farmer(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(500+rand.Intn(500)) * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	randnum := rand.Intn(10) + 1
	wg.Add(randnum)
	for i := 0; i < randnum; i++ {
		go Farmer(&wg)
	}
	wg.Wait()
	println("Gotovo!")
}
