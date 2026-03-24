package main

import (
	"fmt"
	"math/rand"
)

func generate(ch chan int) {
	ch <- rand.Intn(100)
}

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go generate(ch)
	}
	var n int
	for i := 0; i < 5; i++ {
		n = <-ch
		fmt.Printf("%d\n", n)
	}
}
