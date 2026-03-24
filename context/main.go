package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "остановлен")
			return
		default:
			fmt.Println(name, "работает")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(ctx1)
	ctx3, cancel3 := context.WithCancel(ctx2)
	wg.Add(1)
	go worker(ctx1, "Петя", &wg)
	wg.Add(1)
	go worker(ctx2, "Вова", &wg)
	wg.Add(1)
	go worker(ctx3, "Леха", &wg)
	time.Sleep(500 * time.Millisecond)
	cancel3()
	time.Sleep(500 * time.Millisecond)
	cancel2()
	time.Sleep(500 * time.Millisecond)
	cancel1()
	wg.Wait()
}
