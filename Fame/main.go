package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func writealatter(mails *[]string, mu *sync.Mutex) {
	num := rand.IntN(4)
	mu.Lock()
	switch num {
	case 0:
		*mails = append(*mails, "I'm your fan")
	case 1:
		*mails = append(*mails, "I love you")
	case 2:
		*mails = append(*mails, "Hello!")
	case 3:
		*mails = append(*mails, "peniso!")
	}
	mu.Unlock()
}

func main() {
	var mu sync.Mutex
	mails := []string{}
	for i := 0; i < 5; i++ {
		go writealatter(&mails, &mu)
	}
	time.Sleep(1000 * time.Millisecond)
	mu.Lock()
	fmt.Println(mails)
	mu.Unlock()
}
